// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go
//
// Generated by this command:
//
//	mockgen -source=repository.go -destination=repository_mock.go -package=event
//

// Package event is a generated GoMock package.
package event

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
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

// Create mocks base method.
func (m *MockRepository) Create(e Event) (*Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(*Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(e any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), e)
}

// Delete mocks base method.
func (m *MockRepository) Delete(id string) (*Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(*Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockRepository) Get(id string) (*Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoryMockRecorder) Get(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), id)
}

// GetAll mocks base method.
func (m *MockRepository) GetAll() ([]Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepository)(nil).GetAll))
}

// GetAttendance mocks base method.
func (m *MockRepository) GetAttendance(eventID string) (*EventAttendance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttendance", eventID)
	ret0, _ := ret[0].(*EventAttendance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttendance indicates an expected call of GetAttendance.
func (mr *MockRepositoryMockRecorder) GetAttendance(eventID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttendance", reflect.TypeOf((*MockRepository)(nil).GetAttendance), eventID)
}

// GetEventsBetween mocks base method.
func (m *MockRepository) GetEventsBetween(hoursStart, hoursEnd uint) ([]Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventsBetween", hoursStart, hoursEnd)
	ret0, _ := ret[0].([]Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventsBetween indicates an expected call of GetEventsBetween.
func (mr *MockRepositoryMockRecorder) GetEventsBetween(hoursStart, hoursEnd any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventsBetween", reflect.TypeOf((*MockRepository)(nil).GetEventsBetween), hoursStart, hoursEnd)
}

// GetTimeOptionData mocks base method.
func (m *MockRepository) GetTimeOptionData(eventId, invitationId string) ([]*TimeOption, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimeOptionData", eventId, invitationId)
	ret0, _ := ret[0].([]*TimeOption)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTimeOptionData indicates an expected call of GetTimeOptionData.
func (mr *MockRepositoryMockRecorder) GetTimeOptionData(eventId, invitationId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeOptionData", reflect.TypeOf((*MockRepository)(nil).GetTimeOptionData), eventId, invitationId)
}

// GetUnmarkedStaleEvents mocks base method.
func (m *MockRepository) GetUnmarkedStaleEvents() ([]Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnmarkedStaleEvents")
	ret0, _ := ret[0].([]Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnmarkedStaleEvents indicates an expected call of GetUnmarkedStaleEvents.
func (mr *MockRepositoryMockRecorder) GetUnmarkedStaleEvents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnmarkedStaleEvents", reflect.TypeOf((*MockRepository)(nil).GetUnmarkedStaleEvents))
}

// Update mocks base method.
func (m *MockRepository) Update(id string, details Event) (*Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, details)
	ret0, _ := ret[0].(*Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(id, details any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), id, details)
}
