// Code generated by MockGen. DO NOT EDIT.
// Source: internal/ports/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/decadevs/lunch-api/internal/core/models"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreateStaff mocks base method.
func (m *MockUserRepository) CreateStaff(user *models.KitchenStaff) (*models.KitchenStaff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStaff", user)
	ret0, _ := ret[0].(*models.KitchenStaff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStaff indicates an expected call of CreateStaff.
func (mr *MockUserRepositoryMockRecorder) CreateStaff(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStaff", reflect.TypeOf((*MockUserRepository)(nil).CreateStaff), user)
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(user *models.FoodBeneficiary) (*models.FoodBeneficiary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(*models.FoodBeneficiary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), user)
}

// FindStaffByEmail mocks base method.
func (m *MockUserRepository) FindStaffByEmail(email string) (*models.KitchenStaff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindStaffByEmail", email)
	ret0, _ := ret[0].(*models.KitchenStaff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindStaffByEmail indicates an expected call of FindStaffByEmail.
func (mr *MockUserRepositoryMockRecorder) FindStaffByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindStaffByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindStaffByEmail), email)
}

// FindStaffByFullName mocks base method.
func (m *MockUserRepository) FindStaffByFullName(fullname string) (*models.KitchenStaff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindStaffByFullName", fullname)
	ret0, _ := ret[0].(*models.KitchenStaff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindStaffByFullName indicates an expected call of FindStaffByFullName.
func (mr *MockUserRepositoryMockRecorder) FindStaffByFullName(fullname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindStaffByFullName", reflect.TypeOf((*MockUserRepository)(nil).FindStaffByFullName), fullname)
}

// FindStaffByLocation mocks base method.
func (m *MockUserRepository) FindStaffByLocation(location string) (*models.KitchenStaff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindStaffByLocation", location)
	ret0, _ := ret[0].(*models.KitchenStaff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindStaffByLocation indicates an expected call of FindStaffByLocation.
func (mr *MockUserRepositoryMockRecorder) FindStaffByLocation(location interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindStaffByLocation", reflect.TypeOf((*MockUserRepository)(nil).FindStaffByLocation), location)
}

// FindUserByEmail mocks base method.
func (m *MockUserRepository) FindUserByEmail(email string) (*models.FoodBeneficiary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", email)
	ret0, _ := ret[0].(*models.FoodBeneficiary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserRepositoryMockRecorder) FindUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindUserByEmail), email)
}

// FindUserByFullName mocks base method.
func (m *MockUserRepository) FindUserByFullName(fullname string) (*models.FoodBeneficiary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByFullName", fullname)
	ret0, _ := ret[0].(*models.FoodBeneficiary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByFullName indicates an expected call of FindUserByFullName.
func (mr *MockUserRepositoryMockRecorder) FindUserByFullName(fullname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByFullName", reflect.TypeOf((*MockUserRepository)(nil).FindUserByFullName), fullname)
}

// FindUserByLocation mocks base method.
func (m *MockUserRepository) FindUserByLocation(location string) (*models.FoodBeneficiary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByLocation", location)
	ret0, _ := ret[0].(*models.FoodBeneficiary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByLocation indicates an expected call of FindUserByLocation.
func (mr *MockUserRepositoryMockRecorder) FindUserByLocation(location interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByLocation", reflect.TypeOf((*MockUserRepository)(nil).FindUserByLocation), location)
}

// GetByID mocks base method.
func (m *MockUserRepository) GetByID(id string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUserRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUserRepository)(nil).GetByID), id)
}