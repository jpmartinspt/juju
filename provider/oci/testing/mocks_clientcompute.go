// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/provider/oci/common (interfaces: OCIComputeClient)

// Package testing is a generated GoMock package.
package testing

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	core "github.com/oracle/oci-go-sdk/v47/core"
)

// MockOCIComputeClient is a mock of OCIComputeClient interface.
type MockOCIComputeClient struct {
	ctrl     *gomock.Controller
	recorder *MockOCIComputeClientMockRecorder
}

// MockOCIComputeClientMockRecorder is the mock recorder for MockOCIComputeClient.
type MockOCIComputeClientMockRecorder struct {
	mock *MockOCIComputeClient
}

// NewMockOCIComputeClient creates a new mock instance.
func NewMockOCIComputeClient(ctrl *gomock.Controller) *MockOCIComputeClient {
	mock := &MockOCIComputeClient{ctrl: ctrl}
	mock.recorder = &MockOCIComputeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOCIComputeClient) EXPECT() *MockOCIComputeClientMockRecorder {
	return m.recorder
}

// AttachVolume mocks base method.
func (m *MockOCIComputeClient) AttachVolume(arg0 context.Context, arg1 core.AttachVolumeRequest) (core.AttachVolumeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachVolume", arg0, arg1)
	ret0, _ := ret[0].(core.AttachVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachVolume indicates an expected call of AttachVolume.
func (mr *MockOCIComputeClientMockRecorder) AttachVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachVolume", reflect.TypeOf((*MockOCIComputeClient)(nil).AttachVolume), arg0, arg1)
}

// DetachVolume mocks base method.
func (m *MockOCIComputeClient) DetachVolume(arg0 context.Context, arg1 core.DetachVolumeRequest) (core.DetachVolumeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachVolume", arg0, arg1)
	ret0, _ := ret[0].(core.DetachVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetachVolume indicates an expected call of DetachVolume.
func (mr *MockOCIComputeClientMockRecorder) DetachVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachVolume", reflect.TypeOf((*MockOCIComputeClient)(nil).DetachVolume), arg0, arg1)
}

// GetInstance mocks base method.
func (m *MockOCIComputeClient) GetInstance(arg0 context.Context, arg1 core.GetInstanceRequest) (core.GetInstanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstance", arg0, arg1)
	ret0, _ := ret[0].(core.GetInstanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstance indicates an expected call of GetInstance.
func (mr *MockOCIComputeClientMockRecorder) GetInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstance", reflect.TypeOf((*MockOCIComputeClient)(nil).GetInstance), arg0, arg1)
}

// GetVolumeAttachment mocks base method.
func (m *MockOCIComputeClient) GetVolumeAttachment(arg0 context.Context, arg1 core.GetVolumeAttachmentRequest) (core.GetVolumeAttachmentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumeAttachment", arg0, arg1)
	ret0, _ := ret[0].(core.GetVolumeAttachmentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumeAttachment indicates an expected call of GetVolumeAttachment.
func (mr *MockOCIComputeClientMockRecorder) GetVolumeAttachment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumeAttachment", reflect.TypeOf((*MockOCIComputeClient)(nil).GetVolumeAttachment), arg0, arg1)
}

// LaunchInstance mocks base method.
func (m *MockOCIComputeClient) LaunchInstance(arg0 context.Context, arg1 core.LaunchInstanceRequest) (core.LaunchInstanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LaunchInstance", arg0, arg1)
	ret0, _ := ret[0].(core.LaunchInstanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LaunchInstance indicates an expected call of LaunchInstance.
func (mr *MockOCIComputeClientMockRecorder) LaunchInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LaunchInstance", reflect.TypeOf((*MockOCIComputeClient)(nil).LaunchInstance), arg0, arg1)
}

// ListImages mocks base method.
func (m *MockOCIComputeClient) ListImages(arg0 context.Context, arg1 core.ListImagesRequest) (core.ListImagesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImages", arg0, arg1)
	ret0, _ := ret[0].(core.ListImagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImages indicates an expected call of ListImages.
func (mr *MockOCIComputeClientMockRecorder) ListImages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImages", reflect.TypeOf((*MockOCIComputeClient)(nil).ListImages), arg0, arg1)
}

// ListInstances mocks base method.
func (m *MockOCIComputeClient) ListInstances(arg0 context.Context, arg1 core.ListInstancesRequest) (core.ListInstancesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInstances", arg0, arg1)
	ret0, _ := ret[0].(core.ListInstancesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInstances indicates an expected call of ListInstances.
func (mr *MockOCIComputeClientMockRecorder) ListInstances(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInstances", reflect.TypeOf((*MockOCIComputeClient)(nil).ListInstances), arg0, arg1)
}

// ListShapes mocks base method.
func (m *MockOCIComputeClient) ListShapes(arg0 context.Context, arg1 core.ListShapesRequest) (core.ListShapesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListShapes", arg0, arg1)
	ret0, _ := ret[0].(core.ListShapesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListShapes indicates an expected call of ListShapes.
func (mr *MockOCIComputeClientMockRecorder) ListShapes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListShapes", reflect.TypeOf((*MockOCIComputeClient)(nil).ListShapes), arg0, arg1)
}

// ListVnicAttachments mocks base method.
func (m *MockOCIComputeClient) ListVnicAttachments(arg0 context.Context, arg1 core.ListVnicAttachmentsRequest) (core.ListVnicAttachmentsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVnicAttachments", arg0, arg1)
	ret0, _ := ret[0].(core.ListVnicAttachmentsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVnicAttachments indicates an expected call of ListVnicAttachments.
func (mr *MockOCIComputeClientMockRecorder) ListVnicAttachments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVnicAttachments", reflect.TypeOf((*MockOCIComputeClient)(nil).ListVnicAttachments), arg0, arg1)
}

// ListVolumeAttachments mocks base method.
func (m *MockOCIComputeClient) ListVolumeAttachments(arg0 context.Context, arg1 core.ListVolumeAttachmentsRequest) (core.ListVolumeAttachmentsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVolumeAttachments", arg0, arg1)
	ret0, _ := ret[0].(core.ListVolumeAttachmentsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVolumeAttachments indicates an expected call of ListVolumeAttachments.
func (mr *MockOCIComputeClientMockRecorder) ListVolumeAttachments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVolumeAttachments", reflect.TypeOf((*MockOCIComputeClient)(nil).ListVolumeAttachments), arg0, arg1)
}

// TerminateInstance mocks base method.
func (m *MockOCIComputeClient) TerminateInstance(arg0 context.Context, arg1 core.TerminateInstanceRequest) (core.TerminateInstanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TerminateInstance", arg0, arg1)
	ret0, _ := ret[0].(core.TerminateInstanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TerminateInstance indicates an expected call of TerminateInstance.
func (mr *MockOCIComputeClientMockRecorder) TerminateInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminateInstance", reflect.TypeOf((*MockOCIComputeClient)(nil).TerminateInstance), arg0, arg1)
}
