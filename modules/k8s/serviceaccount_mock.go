// Code generated by MockGen. DO NOT EDIT.
// Source: serviceaccount.go
//
// Generated by this command:
//
//	mockgen -source=serviceaccount.go -package=k8s -destination=serviceaccount_mock.go
//

// Package k8s is a generated GoMock package.
package k8s

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockServiceAccountHelper is a mock of ServiceAccountHelper interface.
type MockServiceAccountHelper struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAccountHelperMockRecorder
	isgomock struct{}
}

// MockServiceAccountHelperMockRecorder is the mock recorder for MockServiceAccountHelper.
type MockServiceAccountHelperMockRecorder struct {
	mock *MockServiceAccountHelper
}

// NewMockServiceAccountHelper creates a new mock instance.
func NewMockServiceAccountHelper(ctrl *gomock.Controller) *MockServiceAccountHelper {
	mock := &MockServiceAccountHelper{ctrl: ctrl}
	mock.recorder = &MockServiceAccountHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceAccountHelper) EXPECT() *MockServiceAccountHelperMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockServiceAccountHelper) Create(ctx context.Context, namespace string, serviceAccount *v1.ServiceAccount) (*v1.ServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, namespace, serviceAccount)
	ret0, _ := ret[0].(*v1.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockServiceAccountHelperMockRecorder) Create(ctx, namespace, serviceAccount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockServiceAccountHelper)(nil).Create), ctx, namespace, serviceAccount)
}

// Delete mocks base method.
func (m *MockServiceAccountHelper) Delete(ctx context.Context, namespace, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, namespace, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceAccountHelperMockRecorder) Delete(ctx, namespace, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceAccountHelper)(nil).Delete), ctx, namespace, name)
}

// Get mocks base method.
func (m *MockServiceAccountHelper) Get(ctx context.Context, namespace, name string) (*v1.ServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, namespace, name)
	ret0, _ := ret[0].(*v1.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockServiceAccountHelperMockRecorder) Get(ctx, namespace, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockServiceAccountHelper)(nil).Get), ctx, namespace, name)
}

// List mocks base method.
func (m *MockServiceAccountHelper) List(ctx context.Context, namespace string) (*v1.ServiceAccountList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, namespace)
	ret0, _ := ret[0].(*v1.ServiceAccountList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockServiceAccountHelperMockRecorder) List(ctx, namespace any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServiceAccountHelper)(nil).List), ctx, namespace)
}

// Update mocks base method.
func (m *MockServiceAccountHelper) Update(ctx context.Context, namespace string, serviceAccount *v1.ServiceAccount) (*v1.ServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, namespace, serviceAccount)
	ret0, _ := ret[0].(*v1.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceAccountHelperMockRecorder) Update(ctx, namespace, serviceAccount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServiceAccountHelper)(nil).Update), ctx, namespace, serviceAccount)
}
