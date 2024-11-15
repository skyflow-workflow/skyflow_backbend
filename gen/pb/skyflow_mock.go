// Code generated by MockGen. DO NOT EDIT.
// Source: skyflow.trpc.go

// Package pb is a generated GoMock package.
package pb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	client "trpc.group/trpc-go/trpc-go/client"
)

// MockCommonService is a mock of CommonService interface.
type MockCommonService struct {
	ctrl     *gomock.Controller
	recorder *MockCommonServiceMockRecorder
}

// MockCommonServiceMockRecorder is the mock recorder for MockCommonService.
type MockCommonServiceMockRecorder struct {
	mock *MockCommonService
}

// NewMockCommonService creates a new mock instance.
func NewMockCommonService(ctrl *gomock.Controller) *MockCommonService {
	mock := &MockCommonService{ctrl: ctrl}
	mock.recorder = &MockCommonServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommonService) EXPECT() *MockCommonServiceMockRecorder {
	return m.recorder
}

// HTTP mocks base method.
func (m *MockCommonService) HTTP(ctx context.Context, req *emptypb.Empty) (*HTTPResponseMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTP", ctx, req)
	ret0, _ := ret[0].(*HTTPResponseMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HTTP indicates an expected call of HTTP.
func (mr *MockCommonServiceMockRecorder) HTTP(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTP", reflect.TypeOf((*MockCommonService)(nil).HTTP), ctx, req)
}

// Paging mocks base method.
func (m *MockCommonService) Paging(ctx context.Context, req *PageRequest) (*PageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Paging", ctx, req)
	ret0, _ := ret[0].(*PageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Paging indicates an expected call of Paging.
func (mr *MockCommonServiceMockRecorder) Paging(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Paging", reflect.TypeOf((*MockCommonService)(nil).Paging), ctx, req)
}

// MockSkyflowService is a mock of SkyflowService interface.
type MockSkyflowService struct {
	ctrl     *gomock.Controller
	recorder *MockSkyflowServiceMockRecorder
}

// MockSkyflowServiceMockRecorder is the mock recorder for MockSkyflowService.
type MockSkyflowServiceMockRecorder struct {
	mock *MockSkyflowService
}

// NewMockSkyflowService creates a new mock instance.
func NewMockSkyflowService(ctrl *gomock.Controller) *MockSkyflowService {
	mock := &MockSkyflowService{ctrl: ctrl}
	mock.recorder = &MockSkyflowServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSkyflowService) EXPECT() *MockSkyflowServiceMockRecorder {
	return m.recorder
}

// CreateActivity mocks base method.
func (m *MockSkyflowService) CreateActivity(ctx context.Context, req *CreateActivityRequest) (*CreateActivityResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateActivity", ctx, req)
	ret0, _ := ret[0].(*CreateActivityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateActivity indicates an expected call of CreateActivity.
func (mr *MockSkyflowServiceMockRecorder) CreateActivity(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateActivity", reflect.TypeOf((*MockSkyflowService)(nil).CreateActivity), ctx, req)
}

// CreateNamespace mocks base method.
func (m *MockSkyflowService) CreateNamespace(ctx context.Context, req *CreateNamespaceRequest) (*CreateNamespaceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNamespace", ctx, req)
	ret0, _ := ret[0].(*CreateNamespaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNamespace indicates an expected call of CreateNamespace.
func (mr *MockSkyflowServiceMockRecorder) CreateNamespace(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNamespace", reflect.TypeOf((*MockSkyflowService)(nil).CreateNamespace), ctx, req)
}

// DescribeActivity mocks base method.
func (m *MockSkyflowService) DescribeActivity(ctx context.Context, req *DescribeActivityRequest) (*DescribeActivityResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeActivity", ctx, req)
	ret0, _ := ret[0].(*DescribeActivityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeActivity indicates an expected call of DescribeActivity.
func (mr *MockSkyflowServiceMockRecorder) DescribeActivity(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeActivity", reflect.TypeOf((*MockSkyflowService)(nil).DescribeActivity), ctx, req)
}

// ListActivities mocks base method.
func (m *MockSkyflowService) ListActivities(ctx context.Context, req *ListActivitiesRequest) (*ListActivitiesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListActivities", ctx, req)
	ret0, _ := ret[0].(*ListActivitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListActivities indicates an expected call of ListActivities.
func (mr *MockSkyflowServiceMockRecorder) ListActivities(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListActivities", reflect.TypeOf((*MockSkyflowService)(nil).ListActivities), ctx, req)
}

// ListNamespaces mocks base method.
func (m *MockSkyflowService) ListNamespaces(ctx context.Context, req *ListNamespacesRequest) (*ListNamespacesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNamespaces", ctx, req)
	ret0, _ := ret[0].(*ListNamespacesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNamespaces indicates an expected call of ListNamespaces.
func (mr *MockSkyflowServiceMockRecorder) ListNamespaces(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamespaces", reflect.TypeOf((*MockSkyflowService)(nil).ListNamespaces), ctx, req)
}

// MockCommonClientProxy is a mock of CommonClientProxy interface.
type MockCommonClientProxy struct {
	ctrl     *gomock.Controller
	recorder *MockCommonClientProxyMockRecorder
}

// MockCommonClientProxyMockRecorder is the mock recorder for MockCommonClientProxy.
type MockCommonClientProxyMockRecorder struct {
	mock *MockCommonClientProxy
}

// NewMockCommonClientProxy creates a new mock instance.
func NewMockCommonClientProxy(ctrl *gomock.Controller) *MockCommonClientProxy {
	mock := &MockCommonClientProxy{ctrl: ctrl}
	mock.recorder = &MockCommonClientProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommonClientProxy) EXPECT() *MockCommonClientProxyMockRecorder {
	return m.recorder
}

// HTTP mocks base method.
func (m *MockCommonClientProxy) HTTP(ctx context.Context, req *emptypb.Empty, opts ...client.Option) (*HTTPResponseMessage, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HTTP", varargs...)
	ret0, _ := ret[0].(*HTTPResponseMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HTTP indicates an expected call of HTTP.
func (mr *MockCommonClientProxyMockRecorder) HTTP(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTP", reflect.TypeOf((*MockCommonClientProxy)(nil).HTTP), varargs...)
}

// Paging mocks base method.
func (m *MockCommonClientProxy) Paging(ctx context.Context, req *PageRequest, opts ...client.Option) (*PageResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Paging", varargs...)
	ret0, _ := ret[0].(*PageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Paging indicates an expected call of Paging.
func (mr *MockCommonClientProxyMockRecorder) Paging(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Paging", reflect.TypeOf((*MockCommonClientProxy)(nil).Paging), varargs...)
}

// MockSkyflowClientProxy is a mock of SkyflowClientProxy interface.
type MockSkyflowClientProxy struct {
	ctrl     *gomock.Controller
	recorder *MockSkyflowClientProxyMockRecorder
}

// MockSkyflowClientProxyMockRecorder is the mock recorder for MockSkyflowClientProxy.
type MockSkyflowClientProxyMockRecorder struct {
	mock *MockSkyflowClientProxy
}

// NewMockSkyflowClientProxy creates a new mock instance.
func NewMockSkyflowClientProxy(ctrl *gomock.Controller) *MockSkyflowClientProxy {
	mock := &MockSkyflowClientProxy{ctrl: ctrl}
	mock.recorder = &MockSkyflowClientProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSkyflowClientProxy) EXPECT() *MockSkyflowClientProxyMockRecorder {
	return m.recorder
}

// CreateActivity mocks base method.
func (m *MockSkyflowClientProxy) CreateActivity(ctx context.Context, req *CreateActivityRequest, opts ...client.Option) (*CreateActivityResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateActivity", varargs...)
	ret0, _ := ret[0].(*CreateActivityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateActivity indicates an expected call of CreateActivity.
func (mr *MockSkyflowClientProxyMockRecorder) CreateActivity(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateActivity", reflect.TypeOf((*MockSkyflowClientProxy)(nil).CreateActivity), varargs...)
}

// CreateNamespace mocks base method.
func (m *MockSkyflowClientProxy) CreateNamespace(ctx context.Context, req *CreateNamespaceRequest, opts ...client.Option) (*CreateNamespaceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateNamespace", varargs...)
	ret0, _ := ret[0].(*CreateNamespaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNamespace indicates an expected call of CreateNamespace.
func (mr *MockSkyflowClientProxyMockRecorder) CreateNamespace(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNamespace", reflect.TypeOf((*MockSkyflowClientProxy)(nil).CreateNamespace), varargs...)
}

// DescribeActivity mocks base method.
func (m *MockSkyflowClientProxy) DescribeActivity(ctx context.Context, req *DescribeActivityRequest, opts ...client.Option) (*DescribeActivityResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeActivity", varargs...)
	ret0, _ := ret[0].(*DescribeActivityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeActivity indicates an expected call of DescribeActivity.
func (mr *MockSkyflowClientProxyMockRecorder) DescribeActivity(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeActivity", reflect.TypeOf((*MockSkyflowClientProxy)(nil).DescribeActivity), varargs...)
}

// ListActivities mocks base method.
func (m *MockSkyflowClientProxy) ListActivities(ctx context.Context, req *ListActivitiesRequest, opts ...client.Option) (*ListActivitiesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListActivities", varargs...)
	ret0, _ := ret[0].(*ListActivitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListActivities indicates an expected call of ListActivities.
func (mr *MockSkyflowClientProxyMockRecorder) ListActivities(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListActivities", reflect.TypeOf((*MockSkyflowClientProxy)(nil).ListActivities), varargs...)
}

// ListNamespaces mocks base method.
func (m *MockSkyflowClientProxy) ListNamespaces(ctx context.Context, req *ListNamespacesRequest, opts ...client.Option) (*ListNamespacesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListNamespaces", varargs...)
	ret0, _ := ret[0].(*ListNamespacesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNamespaces indicates an expected call of ListNamespaces.
func (mr *MockSkyflowClientProxyMockRecorder) ListNamespaces(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamespaces", reflect.TypeOf((*MockSkyflowClientProxy)(nil).ListNamespaces), varargs...)
}
