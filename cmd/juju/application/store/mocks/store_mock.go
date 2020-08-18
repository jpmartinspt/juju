// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/application/store (interfaces: CharmAdder,MacaroonGetter)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	charm "github.com/juju/charm/v7"
	params "github.com/juju/charmrepo/v5/csclient/params"
	macaroon "gopkg.in/macaroon.v2"
	reflect "reflect"
)

// MockCharmAdder is a mock of CharmAdder interface
type MockCharmAdder struct {
	ctrl     *gomock.Controller
	recorder *MockCharmAdderMockRecorder
}

// MockCharmAdderMockRecorder is the mock recorder for MockCharmAdder
type MockCharmAdderMockRecorder struct {
	mock *MockCharmAdder
}

// NewMockCharmAdder creates a new mock instance
func NewMockCharmAdder(ctrl *gomock.Controller) *MockCharmAdder {
	mock := &MockCharmAdder{ctrl: ctrl}
	mock.recorder = &MockCharmAdderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCharmAdder) EXPECT() *MockCharmAdderMockRecorder {
	return m.recorder
}

// AddCharm mocks base method
func (m *MockCharmAdder) AddCharm(arg0 *charm.URL, arg1 params.Channel, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCharm", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCharm indicates an expected call of AddCharm
func (mr *MockCharmAdderMockRecorder) AddCharm(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCharm", reflect.TypeOf((*MockCharmAdder)(nil).AddCharm), arg0, arg1, arg2)
}

// AddCharmWithAuthorization mocks base method
func (m *MockCharmAdder) AddCharmWithAuthorization(arg0 *charm.URL, arg1 params.Channel, arg2 *macaroon.Macaroon, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCharmWithAuthorization", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCharmWithAuthorization indicates an expected call of AddCharmWithAuthorization
func (mr *MockCharmAdderMockRecorder) AddCharmWithAuthorization(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCharmWithAuthorization", reflect.TypeOf((*MockCharmAdder)(nil).AddCharmWithAuthorization), arg0, arg1, arg2, arg3)
}

// AddLocalCharm mocks base method
func (m *MockCharmAdder) AddLocalCharm(arg0 *charm.URL, arg1 charm.Charm, arg2 bool) (*charm.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLocalCharm", arg0, arg1, arg2)
	ret0, _ := ret[0].(*charm.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddLocalCharm indicates an expected call of AddLocalCharm
func (mr *MockCharmAdderMockRecorder) AddLocalCharm(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLocalCharm", reflect.TypeOf((*MockCharmAdder)(nil).AddLocalCharm), arg0, arg1, arg2)
}

// MockMacaroonGetter is a mock of MacaroonGetter interface
type MockMacaroonGetter struct {
	ctrl     *gomock.Controller
	recorder *MockMacaroonGetterMockRecorder
}

// MockMacaroonGetterMockRecorder is the mock recorder for MockMacaroonGetter
type MockMacaroonGetterMockRecorder struct {
	mock *MockMacaroonGetter
}

// NewMockMacaroonGetter creates a new mock instance
func NewMockMacaroonGetter(ctrl *gomock.Controller) *MockMacaroonGetter {
	mock := &MockMacaroonGetter{ctrl: ctrl}
	mock.recorder = &MockMacaroonGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMacaroonGetter) EXPECT() *MockMacaroonGetterMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockMacaroonGetter) Get(arg0 string, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockMacaroonGetterMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMacaroonGetter)(nil).Get), arg0, arg1)
}
