// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/blocknative/dreamboat/pkg/api (interfaces: Service)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	structs "github.com/blocknative/dreamboat/pkg/structs"
	types "github.com/flashbots/go-boost-utils/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetBlockReceived mocks base method
func (m *MockService) GetBlockReceived(arg0 context.Context, arg1 structs.TraceQuery) ([]structs.BidTraceWithTimestamp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockReceived", arg0, arg1)
	ret0, _ := ret[0].([]structs.BidTraceWithTimestamp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockReceived indicates an expected call of GetBlockReceived
func (mr *MockServiceMockRecorder) GetBlockReceived(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockReceived", reflect.TypeOf((*MockService)(nil).GetBlockReceived), arg0, arg1)
}

// GetHeader mocks base method
func (m *MockService) GetHeader(arg0 context.Context, arg1 structs.HeaderRequest) (*types.GetHeaderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeader", arg0, arg1)
	ret0, _ := ret[0].(*types.GetHeaderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeader indicates an expected call of GetHeader
func (mr *MockServiceMockRecorder) GetHeader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeader", reflect.TypeOf((*MockService)(nil).GetHeader), arg0, arg1)
}

// GetPayload mocks base method
func (m *MockService) GetPayload(arg0 context.Context, arg1 *types.SignedBlindedBeaconBlock) (*types.GetPayloadResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPayload", arg0, arg1)
	ret0, _ := ret[0].(*types.GetPayloadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPayload indicates an expected call of GetPayload
func (mr *MockServiceMockRecorder) GetPayload(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayload", reflect.TypeOf((*MockService)(nil).GetPayload), arg0, arg1)
}

// GetPayloadDelivered mocks base method
func (m *MockService) GetPayloadDelivered(arg0 context.Context, arg1 structs.TraceQuery) ([]structs.BidTraceExtended, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPayloadDelivered", arg0, arg1)
	ret0, _ := ret[0].([]structs.BidTraceExtended)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPayloadDelivered indicates an expected call of GetPayloadDelivered
func (mr *MockServiceMockRecorder) GetPayloadDelivered(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayloadDelivered", reflect.TypeOf((*MockService)(nil).GetPayloadDelivered), arg0, arg1)
}

// GetValidators mocks base method
func (m *MockService) GetValidators() structs.BuilderGetValidatorsResponseEntrySlice {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidators")
	ret0, _ := ret[0].(structs.BuilderGetValidatorsResponseEntrySlice)
	return ret0
}

// GetValidators indicates an expected call of GetValidators
func (mr *MockServiceMockRecorder) GetValidators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidators", reflect.TypeOf((*MockService)(nil).GetValidators))
}

// RegisterValidator mocks base method
func (m *MockService) RegisterValidator(arg0 context.Context, arg1 []structs.SignedValidatorRegistration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterValidator", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterValidator indicates an expected call of RegisterValidator
func (mr *MockServiceMockRecorder) RegisterValidator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterValidator", reflect.TypeOf((*MockService)(nil).RegisterValidator), arg0, arg1)
}

// Registration mocks base method
func (m *MockService) Registration(arg0 context.Context, arg1 types.PublicKey) (types.SignedValidatorRegistration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Registration", arg0, arg1)
	ret0, _ := ret[0].(types.SignedValidatorRegistration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Registration indicates an expected call of Registration
func (mr *MockServiceMockRecorder) Registration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Registration", reflect.TypeOf((*MockService)(nil).Registration), arg0, arg1)
}

// SubmitBlock mocks base method
func (m *MockService) SubmitBlock(arg0 context.Context, arg1 *types.BuilderSubmitBlockRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitBlock", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitBlock indicates an expected call of SubmitBlock
func (mr *MockServiceMockRecorder) SubmitBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitBlock", reflect.TypeOf((*MockService)(nil).SubmitBlock), arg0, arg1)
}