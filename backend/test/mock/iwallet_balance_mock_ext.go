// Code generated by MockGen. DO NOT EDIT.
// Source: sleuth/repository/interface (interfaces: IWalletBalanceExt)
//
// Generated by this command:
//
//	mockgen -destination=test/mock/iwallet_balance_mock_ext.go -package=mock sleuth/repository/interface IWalletBalanceExt
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	repository "sleuth/model/repository"
	errs "sleuth/utils/errs"

	gomock "go.uber.org/mock/gomock"
)

// MockIWalletBalanceExt is a mock of IWalletBalanceExt interface.
type MockIWalletBalanceExt struct {
	ctrl     *gomock.Controller
	recorder *MockIWalletBalanceExtMockRecorder
}

// MockIWalletBalanceExtMockRecorder is the mock recorder for MockIWalletBalanceExt.
type MockIWalletBalanceExtMockRecorder struct {
	mock *MockIWalletBalanceExt
}

// NewMockIWalletBalanceExt creates a new mock instance.
func NewMockIWalletBalanceExt(ctrl *gomock.Controller) *MockIWalletBalanceExt {
	mock := &MockIWalletBalanceExt{ctrl: ctrl}
	mock.recorder = &MockIWalletBalanceExtMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIWalletBalanceExt) EXPECT() *MockIWalletBalanceExtMockRecorder {
	return m.recorder
}

// GetWalletBalance mocks base method.
func (m *MockIWalletBalanceExt) GetWalletBalance(arg0 context.Context, arg1 string) (*repository.Response, *errs.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletBalance", arg0, arg1)
	ret0, _ := ret[0].(*repository.Response)
	ret1, _ := ret[1].(*errs.ErrorResponse)
	return ret0, ret1
}

// GetWalletBalance indicates an expected call of GetWalletBalance.
func (mr *MockIWalletBalanceExtMockRecorder) GetWalletBalance(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletBalance", reflect.TypeOf((*MockIWalletBalanceExt)(nil).GetWalletBalance), arg0, arg1)
}