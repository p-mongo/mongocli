// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/database_users.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	reflect "reflect"
)

// MockDatabaseUserLister is a mock of DatabaseUserLister interface
type MockDatabaseUserLister struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseUserListerMockRecorder
}

// MockDatabaseUserListerMockRecorder is the mock recorder for MockDatabaseUserLister
type MockDatabaseUserListerMockRecorder struct {
	mock *MockDatabaseUserLister
}

// NewMockDatabaseUserLister creates a new mock instance
func NewMockDatabaseUserLister(ctrl *gomock.Controller) *MockDatabaseUserLister {
	mock := &MockDatabaseUserLister{ctrl: ctrl}
	mock.recorder = &MockDatabaseUserListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabaseUserLister) EXPECT() *MockDatabaseUserListerMockRecorder {
	return m.recorder
}

// DatabaseUsers mocks base method
func (m *MockDatabaseUserLister) DatabaseUsers(groupID string, opts *mongodbatlas.ListOptions) ([]mongodbatlas.DatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseUsers", groupID, opts)
	ret0, _ := ret[0].([]mongodbatlas.DatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatabaseUsers indicates an expected call of DatabaseUsers
func (mr *MockDatabaseUserListerMockRecorder) DatabaseUsers(groupID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseUsers", reflect.TypeOf((*MockDatabaseUserLister)(nil).DatabaseUsers), groupID, opts)
}

// MockDatabaseUserCreator is a mock of DatabaseUserCreator interface
type MockDatabaseUserCreator struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseUserCreatorMockRecorder
}

// MockDatabaseUserCreatorMockRecorder is the mock recorder for MockDatabaseUserCreator
type MockDatabaseUserCreatorMockRecorder struct {
	mock *MockDatabaseUserCreator
}

// NewMockDatabaseUserCreator creates a new mock instance
func NewMockDatabaseUserCreator(ctrl *gomock.Controller) *MockDatabaseUserCreator {
	mock := &MockDatabaseUserCreator{ctrl: ctrl}
	mock.recorder = &MockDatabaseUserCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabaseUserCreator) EXPECT() *MockDatabaseUserCreatorMockRecorder {
	return m.recorder
}

// CreateDatabaseUser mocks base method
func (m *MockDatabaseUserCreator) CreateDatabaseUser(arg0 *mongodbatlas.DatabaseUser) (*mongodbatlas.DatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDatabaseUser", arg0)
	ret0, _ := ret[0].(*mongodbatlas.DatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDatabaseUser indicates an expected call of CreateDatabaseUser
func (mr *MockDatabaseUserCreatorMockRecorder) CreateDatabaseUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDatabaseUser", reflect.TypeOf((*MockDatabaseUserCreator)(nil).CreateDatabaseUser), arg0)
}

// MockDatabaseUserDeleter is a mock of DatabaseUserDeleter interface
type MockDatabaseUserDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseUserDeleterMockRecorder
}

// MockDatabaseUserDeleterMockRecorder is the mock recorder for MockDatabaseUserDeleter
type MockDatabaseUserDeleterMockRecorder struct {
	mock *MockDatabaseUserDeleter
}

// NewMockDatabaseUserDeleter creates a new mock instance
func NewMockDatabaseUserDeleter(ctrl *gomock.Controller) *MockDatabaseUserDeleter {
	mock := &MockDatabaseUserDeleter{ctrl: ctrl}
	mock.recorder = &MockDatabaseUserDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabaseUserDeleter) EXPECT() *MockDatabaseUserDeleterMockRecorder {
	return m.recorder
}

// DeleteDatabaseUser mocks base method
func (m *MockDatabaseUserDeleter) DeleteDatabaseUser(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDatabaseUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDatabaseUser indicates an expected call of DeleteDatabaseUser
func (mr *MockDatabaseUserDeleterMockRecorder) DeleteDatabaseUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDatabaseUser", reflect.TypeOf((*MockDatabaseUserDeleter)(nil).DeleteDatabaseUser), arg0, arg1, arg2)
}

// MockDatabaseUserUpdater is a mock of DatabaseUserUpdater interface
type MockDatabaseUserUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseUserUpdaterMockRecorder
}

// MockDatabaseUserUpdaterMockRecorder is the mock recorder for MockDatabaseUserUpdater
type MockDatabaseUserUpdaterMockRecorder struct {
	mock *MockDatabaseUserUpdater
}

// NewMockDatabaseUserUpdater creates a new mock instance
func NewMockDatabaseUserUpdater(ctrl *gomock.Controller) *MockDatabaseUserUpdater {
	mock := &MockDatabaseUserUpdater{ctrl: ctrl}
	mock.recorder = &MockDatabaseUserUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabaseUserUpdater) EXPECT() *MockDatabaseUserUpdaterMockRecorder {
	return m.recorder
}

// UpdateDatabaseUser mocks base method
func (m *MockDatabaseUserUpdater) UpdateDatabaseUser(arg0 *mongodbatlas.DatabaseUser) (*mongodbatlas.DatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDatabaseUser", arg0)
	ret0, _ := ret[0].(*mongodbatlas.DatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDatabaseUser indicates an expected call of UpdateDatabaseUser
func (mr *MockDatabaseUserUpdaterMockRecorder) UpdateDatabaseUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDatabaseUser", reflect.TypeOf((*MockDatabaseUserUpdater)(nil).UpdateDatabaseUser), arg0)
}

// MockDatabaseUserDescriber is a mock of DatabaseUserDescriber interface
type MockDatabaseUserDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseUserDescriberMockRecorder
}

// MockDatabaseUserDescriberMockRecorder is the mock recorder for MockDatabaseUserDescriber
type MockDatabaseUserDescriberMockRecorder struct {
	mock *MockDatabaseUserDescriber
}

// NewMockDatabaseUserDescriber creates a new mock instance
func NewMockDatabaseUserDescriber(ctrl *gomock.Controller) *MockDatabaseUserDescriber {
	mock := &MockDatabaseUserDescriber{ctrl: ctrl}
	mock.recorder = &MockDatabaseUserDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabaseUserDescriber) EXPECT() *MockDatabaseUserDescriberMockRecorder {
	return m.recorder
}

// DatabaseUser mocks base method
func (m *MockDatabaseUserDescriber) DatabaseUser(arg0, arg1, arg2 string) (*mongodbatlas.DatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.DatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatabaseUser indicates an expected call of DatabaseUser
func (mr *MockDatabaseUserDescriberMockRecorder) DatabaseUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseUser", reflect.TypeOf((*MockDatabaseUserDescriber)(nil).DatabaseUser), arg0, arg1, arg2)
}

// MockDatabaseUserStore is a mock of DatabaseUserStore interface
type MockDatabaseUserStore struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseUserStoreMockRecorder
}

// MockDatabaseUserStoreMockRecorder is the mock recorder for MockDatabaseUserStore
type MockDatabaseUserStoreMockRecorder struct {
	mock *MockDatabaseUserStore
}

// NewMockDatabaseUserStore creates a new mock instance
func NewMockDatabaseUserStore(ctrl *gomock.Controller) *MockDatabaseUserStore {
	mock := &MockDatabaseUserStore{ctrl: ctrl}
	mock.recorder = &MockDatabaseUserStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabaseUserStore) EXPECT() *MockDatabaseUserStoreMockRecorder {
	return m.recorder
}

// CreateDatabaseUser mocks base method
func (m *MockDatabaseUserStore) CreateDatabaseUser(arg0 *mongodbatlas.DatabaseUser) (*mongodbatlas.DatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDatabaseUser", arg0)
	ret0, _ := ret[0].(*mongodbatlas.DatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDatabaseUser indicates an expected call of CreateDatabaseUser
func (mr *MockDatabaseUserStoreMockRecorder) CreateDatabaseUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDatabaseUser", reflect.TypeOf((*MockDatabaseUserStore)(nil).CreateDatabaseUser), arg0)
}

// DeleteDatabaseUser mocks base method
func (m *MockDatabaseUserStore) DeleteDatabaseUser(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDatabaseUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDatabaseUser indicates an expected call of DeleteDatabaseUser
func (mr *MockDatabaseUserStoreMockRecorder) DeleteDatabaseUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDatabaseUser", reflect.TypeOf((*MockDatabaseUserStore)(nil).DeleteDatabaseUser), arg0, arg1, arg2)
}

// UpdateDatabaseUser mocks base method
func (m *MockDatabaseUserStore) UpdateDatabaseUser(arg0 *mongodbatlas.DatabaseUser) (*mongodbatlas.DatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDatabaseUser", arg0)
	ret0, _ := ret[0].(*mongodbatlas.DatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDatabaseUser indicates an expected call of UpdateDatabaseUser
func (mr *MockDatabaseUserStoreMockRecorder) UpdateDatabaseUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDatabaseUser", reflect.TypeOf((*MockDatabaseUserStore)(nil).UpdateDatabaseUser), arg0)
}

// DatabaseUser mocks base method
func (m *MockDatabaseUserStore) DatabaseUser(arg0, arg1, arg2 string) (*mongodbatlas.DatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.DatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatabaseUser indicates an expected call of DatabaseUser
func (mr *MockDatabaseUserStoreMockRecorder) DatabaseUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseUser", reflect.TypeOf((*MockDatabaseUserStore)(nil).DatabaseUser), arg0, arg1, arg2)
}
