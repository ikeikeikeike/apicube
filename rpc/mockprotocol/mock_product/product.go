// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ikeikeikeike/apicube/rpc/protocol/apicube/product (interfaces: ProductServiceServer,ProductServiceClient)

// Package mock_product is a generated GoMock package.
package mock_product

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	product "github.com/ikeikeikeike/apicube/rpc/protocol/apicube/product"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockProductServiceServer is a mock of ProductServiceServer interface
type MockProductServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceServerMockRecorder
}

// MockProductServiceServerMockRecorder is the mock recorder for MockProductServiceServer
type MockProductServiceServerMockRecorder struct {
	mock *MockProductServiceServer
}

// NewMockProductServiceServer creates a new mock instance
func NewMockProductServiceServer(ctrl *gomock.Controller) *MockProductServiceServer {
	mock := &MockProductServiceServer{ctrl: ctrl}
	mock.recorder = &MockProductServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProductServiceServer) EXPECT() *MockProductServiceServerMockRecorder {
	return m.recorder
}

// SimilarProducts mocks base method
func (m *MockProductServiceServer) SimilarProducts(arg0 context.Context, arg1 *product.SimilarProductsRequest) (*product.SimilarProductsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SimilarProducts", arg0, arg1)
	ret0, _ := ret[0].(*product.SimilarProductsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SimilarProducts indicates an expected call of SimilarProducts
func (mr *MockProductServiceServerMockRecorder) SimilarProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SimilarProducts", reflect.TypeOf((*MockProductServiceServer)(nil).SimilarProducts), arg0, arg1)
}

// MockProductServiceClient is a mock of ProductServiceClient interface
type MockProductServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceClientMockRecorder
}

// MockProductServiceClientMockRecorder is the mock recorder for MockProductServiceClient
type MockProductServiceClientMockRecorder struct {
	mock *MockProductServiceClient
}

// NewMockProductServiceClient creates a new mock instance
func NewMockProductServiceClient(ctrl *gomock.Controller) *MockProductServiceClient {
	mock := &MockProductServiceClient{ctrl: ctrl}
	mock.recorder = &MockProductServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProductServiceClient) EXPECT() *MockProductServiceClientMockRecorder {
	return m.recorder
}

// SimilarProducts mocks base method
func (m *MockProductServiceClient) SimilarProducts(arg0 context.Context, arg1 *product.SimilarProductsRequest, arg2 ...grpc.CallOption) (*product.SimilarProductsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SimilarProducts", varargs...)
	ret0, _ := ret[0].(*product.SimilarProductsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SimilarProducts indicates an expected call of SimilarProducts
func (mr *MockProductServiceClientMockRecorder) SimilarProducts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SimilarProducts", reflect.TypeOf((*MockProductServiceClient)(nil).SimilarProducts), varargs...)
}
