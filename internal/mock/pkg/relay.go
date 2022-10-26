// Code generated by MockGen. DO NOT EDIT.
// Source: relay.go

// Package mock_relay is a generated GoMock package.
package mock_relay

import (
	context "context"
	relay "github.com/blocknative/dreamboat/pkg"
	types "github.com/flashbots/go-boost-utils/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockState is a mock of State interface
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// Datastore mocks base method
func (m *MockState) Datastore() relay.Datastore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Datastore")
	ret0, _ := ret[0].(relay.Datastore)
	return ret0
}

// Datastore indicates an expected call of Datastore
func (mr *MockStateMockRecorder) Datastore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Datastore", reflect.TypeOf((*MockState)(nil).Datastore))
}

// Beacon mocks base method
func (m *MockState) Beacon() relay.BeaconState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Beacon")
	ret0, _ := ret[0].(relay.BeaconState)
	return ret0
}

// Beacon indicates an expected call of Beacon
func (mr *MockStateMockRecorder) Beacon() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Beacon", reflect.TypeOf((*MockState)(nil).Beacon))
}

// MockBeaconState is a mock of BeaconState interface
type MockBeaconState struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconStateMockRecorder
}

// MockBeaconStateMockRecorder is the mock recorder for MockBeaconState
type MockBeaconStateMockRecorder struct {
	mock *MockBeaconState
}

// NewMockBeaconState creates a new mock instance
func NewMockBeaconState(ctrl *gomock.Controller) *MockBeaconState {
	mock := &MockBeaconState{ctrl: ctrl}
	mock.recorder = &MockBeaconStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBeaconState) EXPECT() *MockBeaconStateMockRecorder {
	return m.recorder
}

// KnownValidatorByIndex mocks base method
func (m *MockBeaconState) KnownValidatorByIndex(arg0 uint64) (types.PubkeyHex, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KnownValidatorByIndex", arg0)
	ret0, _ := ret[0].(types.PubkeyHex)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KnownValidatorByIndex indicates an expected call of KnownValidatorByIndex
func (mr *MockBeaconStateMockRecorder) KnownValidatorByIndex(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KnownValidatorByIndex", reflect.TypeOf((*MockBeaconState)(nil).KnownValidatorByIndex), arg0)
}

// IsKnownValidator mocks base method
func (m *MockBeaconState) IsKnownValidator(arg0 types.PubkeyHex) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsKnownValidator", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsKnownValidator indicates an expected call of IsKnownValidator
func (mr *MockBeaconStateMockRecorder) IsKnownValidator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsKnownValidator", reflect.TypeOf((*MockBeaconState)(nil).IsKnownValidator), arg0)
}

// HeadSlot mocks base method
func (m *MockBeaconState) HeadSlot() relay.Slot {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeadSlot")
	ret0, _ := ret[0].(relay.Slot)
	return ret0
}

// HeadSlot indicates an expected call of HeadSlot
func (mr *MockBeaconStateMockRecorder) HeadSlot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadSlot", reflect.TypeOf((*MockBeaconState)(nil).HeadSlot))
}

// ValidatorsMap mocks base method
func (m *MockBeaconState) ValidatorsMap() relay.BuilderGetValidatorsResponseEntrySlice {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatorsMap")
	ret0, _ := ret[0].(relay.BuilderGetValidatorsResponseEntrySlice)
	return ret0
}

// ValidatorsMap indicates an expected call of ValidatorsMap
func (mr *MockBeaconStateMockRecorder) ValidatorsMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorsMap", reflect.TypeOf((*MockBeaconState)(nil).ValidatorsMap))
}

// MockRelay is a mock of Relay interface
type MockRelay struct {
	ctrl     *gomock.Controller
	recorder *MockRelayMockRecorder
}

// MockRelayMockRecorder is the mock recorder for MockRelay
type MockRelayMockRecorder struct {
	mock *MockRelay
}

// NewMockRelay creates a new mock instance
func NewMockRelay(ctrl *gomock.Controller) *MockRelay {
	mock := &MockRelay{ctrl: ctrl}
	mock.recorder = &MockRelayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRelay) EXPECT() *MockRelayMockRecorder {
	return m.recorder
}

// RegisterValidator mocks base method
func (m *MockRelay) RegisterValidator(arg0 context.Context, arg1 []types.SignedValidatorRegistration, arg2 relay.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterValidator", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterValidator indicates an expected call of RegisterValidator
func (mr *MockRelayMockRecorder) RegisterValidator(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterValidator", reflect.TypeOf((*MockRelay)(nil).RegisterValidator), arg0, arg1, arg2)
}

// GetHeader mocks base method
func (m *MockRelay) GetHeader(arg0 context.Context, arg1 relay.HeaderRequest, arg2 relay.State) (*types.GetHeaderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeader", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.GetHeaderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeader indicates an expected call of GetHeader
func (mr *MockRelayMockRecorder) GetHeader(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeader", reflect.TypeOf((*MockRelay)(nil).GetHeader), arg0, arg1, arg2)
}

// GetPayload mocks base method
func (m *MockRelay) GetPayload(arg0 context.Context, arg1 *types.SignedBlindedBeaconBlock, arg2 relay.State) (*types.GetPayloadResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPayload", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.GetPayloadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPayload indicates an expected call of GetPayload
func (mr *MockRelayMockRecorder) GetPayload(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayload", reflect.TypeOf((*MockRelay)(nil).GetPayload), arg0, arg1, arg2)
}

// SubmitBlock mocks base method
func (m *MockRelay) SubmitBlock(arg0 context.Context, arg1 *types.BuilderSubmitBlockRequest, arg2 relay.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitBlock indicates an expected call of SubmitBlock
func (mr *MockRelayMockRecorder) SubmitBlock(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitBlock", reflect.TypeOf((*MockRelay)(nil).SubmitBlock), arg0, arg1, arg2)
}

// GetValidators mocks base method
func (m *MockRelay) GetValidators(arg0 relay.State) relay.BuilderGetValidatorsResponseEntrySlice {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidators", arg0)
	ret0, _ := ret[0].(relay.BuilderGetValidatorsResponseEntrySlice)
	return ret0
}

// GetValidators indicates an expected call of GetValidators
func (mr *MockRelayMockRecorder) GetValidators(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidators", reflect.TypeOf((*MockRelay)(nil).GetValidators), arg0)
}
