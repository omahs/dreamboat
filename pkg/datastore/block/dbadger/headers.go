package dbadger

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"regexp"
	"sort"
	"strconv"
	"time"

	"github.com/blocknative/dreamboat/pkg/datastore/block/headerscontroller"
	"github.com/blocknative/dreamboat/pkg/structs"
	"github.com/dgraph-io/badger/v2"
	"github.com/flashbots/go-boost-utils/types"
	ds "github.com/ipfs/go-datastore"
)

const (
	HeaderPrefix        = "header-"
	HeaderContentPrefix = "hc/"
)

var (
	HeaderPrefixBytes = []byte("header-")
)

func HeaderKeyContent(slot uint64, blockHash string) ds.Key {
	return ds.NewKey(fmt.Sprintf("hc/%d/%s", slot, blockHash))
}

func HeaderMaxNewKey(slot uint64) ds.Key {
	return ds.NewKey(fmt.Sprintf("hm/%d", slot))
}

func HeaderKey(slot uint64) ds.Key {
	return ds.NewKey(fmt.Sprintf("%s%d", HeaderPrefix, slot))
}

type StoredIndex struct {
	Index                []IndexMeta
	MaxProfit            IndexMeta
	SubmissionsByPubKeys map[[48]byte]IndexMeta
}

func NewStoreIndex() StoredIndex {
	return StoredIndex{
		SubmissionsByPubKeys: make(map[[48]byte]IndexMeta),
	}
}

type IndexMeta struct {
	Hash          [32]byte
	Value         *big.Int
	BuilderPubkey [48]byte
}

type SlotInfo struct {
	Slot  uint64
	Added time.Time
}

func (s *Datastore) GetMaxProfitHeader(ctx context.Context, slot uint64) (structs.HeaderAndTrace, error) {
	p, ok := s.hc.GetMaxProfit(uint64(slot))
	if ok {
		return p, nil
	}

	p, err := s.getMaxHeader(ctx, slot)
	return p, err
}

var ErrNotFound = errors.New("not found")

func (s *Datastore) getMaxHeader(ctx context.Context, slot uint64) (h structs.HeaderAndTrace, err error) {
	txn := s.DBInter.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get(HeaderMaxNewKey(slot).Bytes())
	if err != nil {
		return h, err
	}

	v, err := item.ValueCopy(nil)
	if err != nil {
		return h, err
	}
	item, err = txn.Get(HeaderKeyContent(slot, string(v)).Bytes())
	if err != nil {
		return h, err
	}

	h = structs.HeaderAndTrace{}
	err = item.Value(func(val []byte) error {
		return json.Unmarshal(val, &h)
	})

	return h, err
}

func (s *Datastore) PutHeader(ctx context.Context, hd structs.HeaderData, ttl time.Duration) (err error) {
	if err := storeHeader(s.DBInter, hd, ttl); err != nil {
		return err
	}

	newlyCreated, err := s.hc.Add(uint64(hd.Slot), hd.HeaderAndTrace)
	if err != nil {
		return err
	}

	if !newlyCreated {
		return // success
	}

	// check and load keys if exists
	return s.loadKeysAndCleanup(ctx, uint64(hd.Slot))
}

func (s *Datastore) loadKeysAndCleanup(ctx context.Context, slot uint64) error {
	s.l.Lock()
	defer s.l.Unlock()

	// need to load keys to memory
	data, err := s.DB.Get(ctx, HeaderKey(slot))
	if err != nil {
		if errors.Is(err, ds.ErrNotFound) {
			return nil // success - key is empty, nothing to load
		}
		return err
	}
	var h []structs.HeaderAndTrace
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}

	return s.hc.PrependMultiple(slot, h)
}

func storeHeader(s DBInter, h structs.HeaderData, ttl time.Duration) error {
	txn := s.NewTransaction(true)
	defer txn.Discard()

	// we don't need to lock here, as the value would be always different from different block
	if err := txn.SetEntry(badger.NewEntry(HeaderKeyContent(uint64(h.Slot), h.Header.BlockHash.String()).Bytes(), h.Marshaled).WithTTL(ttl)); err != nil {
		return err
	}
	slot := make([]byte, 8)
	binary.LittleEndian.PutUint64(slot, uint64(h.Slot))

	if err := txn.SetEntry(badger.NewEntry(HeaderHashKey(h.Header.BlockHash).Bytes(), slot).WithTTL(ttl)); err != nil {
		return err
	}

	// not needed every time
	if err := txn.SetEntry(badger.NewEntry(HeaderNumKey(h.Header.BlockNumber).Bytes(), slot).WithTTL(ttl)); err != nil {
		return err
	}

	return txn.Commit()
}

// SaveHeaders is meant to persist the all the keys under one key
// As optimization in future this function can operate only on database, so instead from memory it may just reorganize keys
func (s *Datastore) SaveHeaders(ctx context.Context, slots []uint64, ttl time.Duration) error {
	for _, slot := range slots {
		if err := s.saveHeader(ctx, slot, ttl); err != nil {
			return err
		}
	}
	return nil
}

func (s *Datastore) saveHeader(ctx context.Context, slot uint64, ttl time.Duration) error {
	el, maxP, rev, err := s.hc.GetSingleSlot(slot)
	if err != nil {
		return err
	}

	if err = putHeaders(ctx, s, slot, el, ttl); err != nil {
		return err
	}

	if err := s.DB.PutWithTTL(ctx, HeaderMaxNewKey(slot), []byte(types.Hash(maxP).String()), ttl); err != nil {
		return err
	}

	if s.hc.RemoveSlot(slot, rev) {
		return nil // success
	}

	// revert the saveHeaders operation as revision changed
	_ = s.DBInter.Update(func(txn *badger.Txn) error {
		return txn.Delete(HeaderKey(slot).Bytes())
	})

	return errors.New("revision changed")
}

func putHeaders(ctx context.Context, s *Datastore, slot uint64, cont []structs.HeaderAndTrace, ttl time.Duration) error {
	buff := bytes.NewBuffer(nil)
	buff.WriteString("[")

	enc := json.NewEncoder(buff)
	for i, c := range cont {
		if i > 0 {
			buff.WriteString(",")
		}
		if err := enc.Encode(c); err != nil {
			return err
		}
	}
	buff.WriteString("]")
	if err := s.DB.PutWithTTL(ctx, HeaderKey(slot), buff.Bytes(), ttl); err != nil {
		return err
	}

	buff.Truncate(0) // immediately remove
	return nil
}

type LoadItem struct {
	Time    uint64
	Content []byte
}

// FixOrphanHeaders is reading all the orphan headers from
func (s *Datastore) FixOrphanHeaders(ctx context.Context, ttl time.Duration) error {
	slotDoesNotExist := make(map[uint64][]LoadItem)

	slotExists := make(map[uint64]struct{})

	// Get all headers, rebuild
	err := s.DBInter.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false

		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		prefix := []byte("/" + HeaderContentPrefix)
		re := regexp.MustCompile(`\/hc\/([^\/]+)\/([^\/]+)`)

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			subM := re.FindSubmatch(item.Key())
			if len(subM) != 3 {
				continue
			}
			slot, err := strconv.ParseUint(string(subM[1]), 10, 64)
			if err != nil {
				continue
			}

			if _, ok := slotExists[slot]; ok { // we know it already exists
				continue
			}

			if content, ok := slotDoesNotExist[slot]; ok { // we know it doesn't exists
				li := LoadItem{Time: item.Version()}
				li.Content, err = item.ValueCopy(nil)
				if err != nil {
					return err
				}
				slotDoesNotExist[slot] = append(content, li)
				continue
			}

			_, err = txn.Get(HeaderKey(slot).Bytes())
			if err == nil {
				slotExists[slot] = struct{}{}
				continue
			}
			if !errors.Is(err, badger.ErrKeyNotFound) {
				return err
			}

			li := LoadItem{Time: item.Version()}
			li.Content, err = item.ValueCopy(nil)
			if err != nil {
				return err
			}
			slotDoesNotExist[slot] = append([]LoadItem{}, li)
		}
		return nil
	})
	if err != nil {
		return err
	}

	buff := new(bytes.Buffer)
	for slot, v := range slotDoesNotExist {
		if v != nil || len(v) != 0 {
			tempHC := headerscontroller.NewHeaderController(100, time.Hour) // params doesn't matter here

			buff.Reset()
			sort.Slice(v, func(i, j int) bool {
				return v[i].Time > v[j].Time
			})

			buff.WriteString("[")
			for i, payload := range v {
				if i > 0 {
					buff.WriteString(",")
				}
				io.Copy(buff, bytes.NewReader(payload.Content))
				hnt := structs.HeaderAndTrace{}
				if err := json.Unmarshal(payload.Content, &hnt); err != nil {
					return err
				}

				if _, err := tempHC.Add(slot, hnt); err != nil {
					return err
				}
			}
			buff.WriteString("]")
			if err := s.DB.PutWithTTL(ctx, HeaderKey(slot), buff.Bytes(), ttl); err != nil {
				return err
			}

			maxProfit, ok := tempHC.GetMaxProfit(slot)
			if !ok {
				return errors.New("max profit from records not calculated")
			}
			if err := s.DB.PutWithTTL(ctx, HeaderMaxNewKey(slot), []byte(maxProfit.Trace.BlockHash.String()), ttl); err != nil {
				return err
			}
		}
	}
	return err
}

func (s *Datastore) MemoryCleanup(ctx context.Context, slotPurgeDuration time.Duration, ttl time.Duration) error {
	for {
		time.Sleep(slotPurgeDuration)
		slots, ok := s.hc.CheckForRemoval()
		if !ok {
			continue
		}
		s.SaveHeaders(ctx, slots, ttl)
	}
}
