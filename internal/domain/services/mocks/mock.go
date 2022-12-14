// Code generated by MockGen. DO NOT EDIT.
// Source: services.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	reflect "reflect"
	dto "web/internal/domain/entities/dto"
	model "web/internal/domain/entities/model"

	gomock "github.com/golang/mock/gomock"
)

// MockUserAuthService is a mock of UserAuthService interface.
type MockUserAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockUserAuthServiceMockRecorder
}

// MockUserAuthServiceMockRecorder is the mock recorder for MockUserAuthService.
type MockUserAuthServiceMockRecorder struct {
	mock *MockUserAuthService
}

// NewMockUserAuthService creates a new mock instance.
func NewMockUserAuthService(ctrl *gomock.Controller) *MockUserAuthService {
	mock := &MockUserAuthService{ctrl: ctrl}
	mock.recorder = &MockUserAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAuthService) EXPECT() *MockUserAuthServiceMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockUserAuthService) GenerateToken(userName, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", userName, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockUserAuthServiceMockRecorder) GenerateToken(userName, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockUserAuthService)(nil).GenerateToken), userName, password)
}

// ParseToken mocks base method.
func (m *MockUserAuthService) ParseToken(token string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockUserAuthServiceMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockUserAuthService)(nil).ParseToken), token)
}

// RegisterUser mocks base method.
func (m *MockUserAuthService) RegisterUser(user *model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", user)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockUserAuthServiceMockRecorder) RegisterUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockUserAuthService)(nil).RegisterUser), user)
}

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// DeleteUser mocks base method.
func (m *MockUserService) DeleteUser(id string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserServiceMockRecorder) DeleteUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserService)(nil).DeleteUser), id)
}

// GetAllUsers mocks base method.
func (m *MockUserService) GetAllUsers() ([]dto.UserResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].([]dto.UserResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockUserServiceMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserService)(nil).GetAllUsers))
}

// GetUserByID mocks base method.
func (m *MockUserService) GetUserByID(id string) (*dto.UserResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", id)
	ret0, _ := ret[0].(*dto.UserResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserServiceMockRecorder) GetUserByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserService)(nil).GetUserByID), id)
}

// UpdateUser mocks base method.
func (m *MockUserService) UpdateUser(newUser *dto.UserUpdate, userId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", newUser, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserServiceMockRecorder) UpdateUser(newUser, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserService)(nil).UpdateUser), newUser, userId)
}

// MockNoteService is a mock of NoteService interface.
type MockNoteService struct {
	ctrl     *gomock.Controller
	recorder *MockNoteServiceMockRecorder
}

// MockNoteServiceMockRecorder is the mock recorder for MockNoteService.
type MockNoteServiceMockRecorder struct {
	mock *MockNoteService
}

// NewMockNoteService creates a new mock instance.
func NewMockNoteService(ctrl *gomock.Controller) *MockNoteService {
	mock := &MockNoteService{ctrl: ctrl}
	mock.recorder = &MockNoteServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNoteService) EXPECT() *MockNoteServiceMockRecorder {
	return m.recorder
}

// CreateNote mocks base method.
func (m *MockNoteService) CreateNote(note *model.Note, userID string) (*model.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNote", note, userID)
	ret0, _ := ret[0].(*model.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNote indicates an expected call of CreateNote.
func (mr *MockNoteServiceMockRecorder) CreateNote(note, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNote", reflect.TypeOf((*MockNoteService)(nil).CreateNote), note, userID)
}

// DeleteNote mocks base method.
func (m *MockNoteService) DeleteNote(noteID, userID string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNote", noteID, userID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNote indicates an expected call of DeleteNote.
func (mr *MockNoteServiceMockRecorder) DeleteNote(noteID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNote", reflect.TypeOf((*MockNoteService)(nil).DeleteNote), noteID, userID)
}

// GetAllNotesByUser mocks base method.
func (m *MockNoteService) GetAllNotesByUser(userID string) ([]dto.NotesResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllNotesByUser", userID)
	ret0, _ := ret[0].([]dto.NotesResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllNotesByUser indicates an expected call of GetAllNotesByUser.
func (mr *MockNoteServiceMockRecorder) GetAllNotesByUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllNotesByUser", reflect.TypeOf((*MockNoteService)(nil).GetAllNotesByUser), userID)
}

// GetAllNotesWithTags mocks base method.
func (m *MockNoteService) GetAllNotesWithTags(userID string, notes []dto.NotesResp) ([]dto.NoteWithTagsResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllNotesWithTags", userID, notes)
	ret0, _ := ret[0].([]dto.NoteWithTagsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllNotesWithTags indicates an expected call of GetAllNotesWithTags.
func (mr *MockNoteServiceMockRecorder) GetAllNotesWithTags(userID, notes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllNotesWithTags", reflect.TypeOf((*MockNoteService)(nil).GetAllNotesWithTags), userID, notes)
}

// GetNoteByID mocks base method.
func (m *MockNoteService) GetNoteByID(noteID, userID string) (*dto.NoteResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNoteByID", noteID, userID)
	ret0, _ := ret[0].(*dto.NoteResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNoteByID indicates an expected call of GetNoteByID.
func (mr *MockNoteServiceMockRecorder) GetNoteByID(noteID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNoteByID", reflect.TypeOf((*MockNoteService)(nil).GetNoteByID), noteID, userID)
}

// GetNoteWithAllTags mocks base method.
func (m *MockNoteService) GetNoteWithAllTags(userID, NoteID string, note *dto.NoteResp) (dto.NoteWithTagsResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNoteWithAllTags", userID, NoteID, note)
	ret0, _ := ret[0].(dto.NoteWithTagsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNoteWithAllTags indicates an expected call of GetNoteWithAllTags.
func (mr *MockNoteServiceMockRecorder) GetNoteWithAllTags(userID, NoteID, note interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNoteWithAllTags", reflect.TypeOf((*MockNoteService)(nil).GetNoteWithAllTags), userID, NoteID, note)
}

// RemoveTags mocks base method.
func (m *MockNoteService) RemoveTags(noteID string, tags map[string]string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTags", noteID, tags)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveTags indicates an expected call of RemoveTags.
func (mr *MockNoteServiceMockRecorder) RemoveTags(noteID, tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTags", reflect.TypeOf((*MockNoteService)(nil).RemoveTags), noteID, tags)
}

// SetTags mocks base method.
func (m *MockNoteService) SetTags(noteID string, tags map[string]string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTags", noteID, tags)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetTags indicates an expected call of SetTags.
func (mr *MockNoteServiceMockRecorder) SetTags(noteID, tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTags", reflect.TypeOf((*MockNoteService)(nil).SetTags), noteID, tags)
}

// UpdateNote mocks base method.
func (m *MockNoteService) UpdateNote(newNote *dto.NoteUpdate, noteID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNote", newNote, noteID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNote indicates an expected call of UpdateNote.
func (mr *MockNoteServiceMockRecorder) UpdateNote(newNote, noteID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNote", reflect.TypeOf((*MockNoteService)(nil).UpdateNote), newNote, noteID)
}

// MockTagService is a mock of TagService interface.
type MockTagService struct {
	ctrl     *gomock.Controller
	recorder *MockTagServiceMockRecorder
}

// MockTagServiceMockRecorder is the mock recorder for MockTagService.
type MockTagServiceMockRecorder struct {
	mock *MockTagService
}

// NewMockTagService creates a new mock instance.
func NewMockTagService(ctrl *gomock.Controller) *MockTagService {
	mock := &MockTagService{ctrl: ctrl}
	mock.recorder = &MockTagServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTagService) EXPECT() *MockTagServiceMockRecorder {
	return m.recorder
}

// CreateTag mocks base method.
func (m *MockTagService) CreateTag(tag *model.Tag, userID string) (*model.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTag", tag, userID)
	ret0, _ := ret[0].(*model.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTag indicates an expected call of CreateTag.
func (mr *MockTagServiceMockRecorder) CreateTag(tag, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTag", reflect.TypeOf((*MockTagService)(nil).CreateTag), tag, userID)
}

// DeleteTag mocks base method.
func (m *MockTagService) DeleteTag(tagID, userID string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTag", tagID, userID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTag indicates an expected call of DeleteTag.
func (mr *MockTagServiceMockRecorder) DeleteTag(tagID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTag", reflect.TypeOf((*MockTagService)(nil).DeleteTag), tagID, userID)
}

// GetAllTagsByUser mocks base method.
func (m *MockTagService) GetAllTagsByUser(userID string) ([]dto.TagsResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTagsByUser", userID)
	ret0, _ := ret[0].([]dto.TagsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTagsByUser indicates an expected call of GetAllTagsByUser.
func (mr *MockTagServiceMockRecorder) GetAllTagsByUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTagsByUser", reflect.TypeOf((*MockTagService)(nil).GetAllTagsByUser), userID)
}

// GetTagByID mocks base method.
func (m *MockTagService) GetTagByID(tagID, userID string) (*dto.TagResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagByID", tagID, userID)
	ret0, _ := ret[0].(*dto.TagResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagByID indicates an expected call of GetTagByID.
func (mr *MockTagServiceMockRecorder) GetTagByID(tagID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagByID", reflect.TypeOf((*MockTagService)(nil).GetTagByID), tagID, userID)
}

// UpdateTag mocks base method.
func (m *MockTagService) UpdateTag(tag *dto.TagUpdate, tagID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTag", tag, tagID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTag indicates an expected call of UpdateTag.
func (mr *MockTagServiceMockRecorder) UpdateTag(tag, tagID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTag", reflect.TypeOf((*MockTagService)(nil).UpdateTag), tag, tagID)
}
