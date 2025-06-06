// Code generated by MockGen. DO NOT EDIT.
// Source: account_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	domain "gate/internal/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAccountRepository is a mock of AccountRepository interface.
type MockAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepositoryMockRecorder
}

// MockAccountRepositoryMockRecorder is the mock recorder for MockAccountRepository.
type MockAccountRepositoryMockRecorder struct {
	mock *MockAccountRepository
}

// NewMockAccountRepository creates a new mock instance.
func NewMockAccountRepository(ctrl *gomock.Controller) *MockAccountRepository {
	mock := &MockAccountRepository{ctrl: ctrl}
	mock.recorder = &MockAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepository) EXPECT() *MockAccountRepositoryMockRecorder {
	return m.recorder
}

// GetAccountInfoByAccount mocks base method.
func (m *MockAccountRepository) GetAccountInfoByAccount(account string) (domain.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountInfoByAccount", account)
	ret0, _ := ret[0].(domain.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountInfoByAccount indicates an expected call of GetAccountInfoByAccount.
func (mr *MockAccountRepositoryMockRecorder) GetAccountInfoByAccount(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountInfoByAccount", reflect.TypeOf((*MockAccountRepository)(nil).GetAccountInfoByAccount), account)
}

// InsertAccount mocks base method.
func (m *MockAccountRepository) InsertAccount(account domain.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertAccount", account)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertAccount indicates an expected call of InsertAccount.
func (mr *MockAccountRepositoryMockRecorder) InsertAccount(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertAccount", reflect.TypeOf((*MockAccountRepository)(nil).InsertAccount), account)
}
