// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/olezhek28/link-shortener/internal/repository/links (interfaces: Repository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddLink mocks base method.
func (m *MockRepository) AddLink(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLink", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddLink indicates an expected call of AddLink.
func (mr *MockRepositoryMockRecorder) AddLink(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLink", reflect.TypeOf((*MockRepository)(nil).AddLink), arg0, arg1, arg2)
}

// GetLongLink mocks base method.
func (m *MockRepository) GetLongLink(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLongLink", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLongLink indicates an expected call of GetLongLink.
func (mr *MockRepositoryMockRecorder) GetLongLink(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLongLink", reflect.TypeOf((*MockRepository)(nil).GetLongLink), arg0, arg1)
}
