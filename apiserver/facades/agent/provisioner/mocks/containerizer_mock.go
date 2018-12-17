// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/network/containerizer (interfaces: LinkLayerDevice,Unit,Application,Charm)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	containerizer "github.com/juju/juju/network/containerizer"
	state "github.com/juju/juju/state"
	charm_v6 "gopkg.in/juju/charm.v6"
)

// MockLinkLayerDevice is a mock of LinkLayerDevice interface
type MockLinkLayerDevice struct {
	ctrl     *gomock.Controller
	recorder *MockLinkLayerDeviceMockRecorder
}

// MockLinkLayerDeviceMockRecorder is the mock recorder for MockLinkLayerDevice
type MockLinkLayerDeviceMockRecorder struct {
	mock *MockLinkLayerDevice
}

// NewMockLinkLayerDevice creates a new mock instance
func NewMockLinkLayerDevice(ctrl *gomock.Controller) *MockLinkLayerDevice {
	mock := &MockLinkLayerDevice{ctrl: ctrl}
	mock.recorder = &MockLinkLayerDeviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLinkLayerDevice) EXPECT() *MockLinkLayerDeviceMockRecorder {
	return m.recorder
}

// Addresses mocks base method
func (m *MockLinkLayerDevice) Addresses() ([]*state.Address, error) {
	ret := m.ctrl.Call(m, "Addresses")
	ret0, _ := ret[0].([]*state.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Addresses indicates an expected call of Addresses
func (mr *MockLinkLayerDeviceMockRecorder) Addresses() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addresses", reflect.TypeOf((*MockLinkLayerDevice)(nil).Addresses))
}

// EthernetDeviceForBridge mocks base method
func (m *MockLinkLayerDevice) EthernetDeviceForBridge(arg0 string) (state.LinkLayerDeviceArgs, error) {
	ret := m.ctrl.Call(m, "EthernetDeviceForBridge", arg0)
	ret0, _ := ret[0].(state.LinkLayerDeviceArgs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EthernetDeviceForBridge indicates an expected call of EthernetDeviceForBridge
func (mr *MockLinkLayerDeviceMockRecorder) EthernetDeviceForBridge(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EthernetDeviceForBridge", reflect.TypeOf((*MockLinkLayerDevice)(nil).EthernetDeviceForBridge), arg0)
}

// IsAutoStart mocks base method
func (m *MockLinkLayerDevice) IsAutoStart() bool {
	ret := m.ctrl.Call(m, "IsAutoStart")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAutoStart indicates an expected call of IsAutoStart
func (mr *MockLinkLayerDeviceMockRecorder) IsAutoStart() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAutoStart", reflect.TypeOf((*MockLinkLayerDevice)(nil).IsAutoStart))
}

// IsUp mocks base method
func (m *MockLinkLayerDevice) IsUp() bool {
	ret := m.ctrl.Call(m, "IsUp")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUp indicates an expected call of IsUp
func (mr *MockLinkLayerDeviceMockRecorder) IsUp() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUp", reflect.TypeOf((*MockLinkLayerDevice)(nil).IsUp))
}

// MACAddress mocks base method
func (m *MockLinkLayerDevice) MACAddress() string {
	ret := m.ctrl.Call(m, "MACAddress")
	ret0, _ := ret[0].(string)
	return ret0
}

// MACAddress indicates an expected call of MACAddress
func (mr *MockLinkLayerDeviceMockRecorder) MACAddress() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MACAddress", reflect.TypeOf((*MockLinkLayerDevice)(nil).MACAddress))
}

// MTU mocks base method
func (m *MockLinkLayerDevice) MTU() uint {
	ret := m.ctrl.Call(m, "MTU")
	ret0, _ := ret[0].(uint)
	return ret0
}

// MTU indicates an expected call of MTU
func (mr *MockLinkLayerDeviceMockRecorder) MTU() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MTU", reflect.TypeOf((*MockLinkLayerDevice)(nil).MTU))
}

// Name mocks base method
func (m *MockLinkLayerDevice) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockLinkLayerDeviceMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockLinkLayerDevice)(nil).Name))
}

// ParentDevice mocks base method
func (m *MockLinkLayerDevice) ParentDevice() (containerizer.LinkLayerDevice, error) {
	ret := m.ctrl.Call(m, "ParentDevice")
	ret0, _ := ret[0].(containerizer.LinkLayerDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParentDevice indicates an expected call of ParentDevice
func (mr *MockLinkLayerDeviceMockRecorder) ParentDevice() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParentDevice", reflect.TypeOf((*MockLinkLayerDevice)(nil).ParentDevice))
}

// ParentName mocks base method
func (m *MockLinkLayerDevice) ParentName() string {
	ret := m.ctrl.Call(m, "ParentName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ParentName indicates an expected call of ParentName
func (mr *MockLinkLayerDeviceMockRecorder) ParentName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParentName", reflect.TypeOf((*MockLinkLayerDevice)(nil).ParentName))
}

// Type mocks base method
func (m *MockLinkLayerDevice) Type() state.LinkLayerDeviceType {
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(state.LinkLayerDeviceType)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockLinkLayerDeviceMockRecorder) Type() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockLinkLayerDevice)(nil).Type))
}

// MockUnit is a mock of Unit interface
type MockUnit struct {
	ctrl     *gomock.Controller
	recorder *MockUnitMockRecorder
}

// MockUnitMockRecorder is the mock recorder for MockUnit
type MockUnitMockRecorder struct {
	mock *MockUnit
}

// NewMockUnit creates a new mock instance
func NewMockUnit(ctrl *gomock.Controller) *MockUnit {
	mock := &MockUnit{ctrl: ctrl}
	mock.recorder = &MockUnitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnit) EXPECT() *MockUnitMockRecorder {
	return m.recorder
}

// Application mocks base method
func (m *MockUnit) Application() (containerizer.Application, error) {
	ret := m.ctrl.Call(m, "Application")
	ret0, _ := ret[0].(containerizer.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application
func (mr *MockUnitMockRecorder) Application() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockUnit)(nil).Application))
}

// Name mocks base method
func (m *MockUnit) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockUnitMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockUnit)(nil).Name))
}

// MockApplication is a mock of Application interface
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// Charm mocks base method
func (m *MockApplication) Charm() (containerizer.Charm, bool, error) {
	ret := m.ctrl.Call(m, "Charm")
	ret0, _ := ret[0].(containerizer.Charm)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Charm indicates an expected call of Charm
func (mr *MockApplicationMockRecorder) Charm() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Charm", reflect.TypeOf((*MockApplication)(nil).Charm))
}

// Name mocks base method
func (m *MockApplication) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockApplicationMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockApplication)(nil).Name))
}

// MockCharm is a mock of Charm interface
type MockCharm struct {
	ctrl     *gomock.Controller
	recorder *MockCharmMockRecorder
}

// MockCharmMockRecorder is the mock recorder for MockCharm
type MockCharmMockRecorder struct {
	mock *MockCharm
}

// NewMockCharm creates a new mock instance
func NewMockCharm(ctrl *gomock.Controller) *MockCharm {
	mock := &MockCharm{ctrl: ctrl}
	mock.recorder = &MockCharmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCharm) EXPECT() *MockCharmMockRecorder {
	return m.recorder
}

// LXDProfile mocks base method
func (m *MockCharm) LXDProfile() *charm_v6.LXDProfile {
	ret := m.ctrl.Call(m, "LXDProfile")
	ret0, _ := ret[0].(*charm_v6.LXDProfile)
	return ret0
}

// LXDProfile indicates an expected call of LXDProfile
func (mr *MockCharmMockRecorder) LXDProfile() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LXDProfile", reflect.TypeOf((*MockCharm)(nil).LXDProfile))
}

// Revision mocks base method
func (m *MockCharm) Revision() int {
	ret := m.ctrl.Call(m, "Revision")
	ret0, _ := ret[0].(int)
	return ret0
}

// Revision indicates an expected call of Revision
func (mr *MockCharmMockRecorder) Revision() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revision", reflect.TypeOf((*MockCharm)(nil).Revision))
}
