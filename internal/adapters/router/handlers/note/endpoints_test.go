package note

import (
	"bytes"
	e "errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/require"

	"web/internal/config"
	"web/internal/domain/enteties/dto"
	"web/internal/domain/enteties/model"
	"web/internal/domain/errors"
	"web/internal/domain/services"
	"web/internal/domain/services/mocks"
	"web/internal/utils"
	l "web/pkg/logger"
)

func TestHandler_CreateNote(t *testing.T) {
	// Init mock func obj
	type mockBehavior func(s *mock_services.MockNoteService, note *model.Note, userID string)

	testTable := []struct {
		headerName         string
		headerValue        string
		inputJson          string
		inputNote          *model.Note
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
		testName           string
	}{
		{
			headerName: "user_id",
			inputJson: `{
				"title": "test_title",
				"info": "test_info"
			}`,
			// service request
			headerValue: "1",
			inputNote: &model.Note{
				Title: "test_title",
				Info:  "test_info",
			},
			mockBehavior: func(s *mock_services.MockNoteService, note *model.Note, userID string) {
				// service response
				outputNote := &model.Note{
					ID:    "1",
					Title: "test_title",
					Info:  "test_info",
				}
				s.EXPECT().CreateNote(note, userID).Return(outputNote, nil)
			},
			expectedStatusCode: http.StatusCreated,
			expectedResponse: `{"Created note 'test_title' with id":"1"}
`,
			testName: "test-1-Handler:OK",
		},
		{
			inputJson: `{
				"title": "test_title
			}`,
			inputNote:          &model.Note{},
			mockBehavior:       func(s *mock_services.MockNoteService, note *model.Note, userID string) {},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `invalid character '\n' in string literal
`,
			testName: "test-2-Handler:Bad JSON",
		},
		{
			inputJson:          `{}`,
			inputNote:          &model.Note{},
			mockBehavior:       func(s *mock_services.MockNoteService, note *model.Note, userID string) {},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `Key: 'Note.Title' Error:Field validation for 'Title' failed on the 'required' tag
`,
			testName: "test-3-Handler:Validation Err",
		},
		{
			headerName: "user_id",
			inputJson: `{
				"title": "test_title",
				"info": "test_info"
			}`,
			// service request
			headerValue: "1",
			inputNote: &model.Note{
				Title: "test_title",
				Info:  "test_info",
			},
			mockBehavior: func(s *mock_services.MockNoteService, note *model.Note, userID string) {
				// service response
				s.EXPECT().CreateNote(note, userID).Return(nil, e.New(errors.ErrDBDuplicate))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"error":"note 'test_title' is already exists"}
`,
			testName: "test-4-Service:Note is already exists Err",
		},
		{
			headerName: "user_id",
			inputJson: `{
				"title": "test_title",
				"info": "test_info"
			}`,
			// service request
			headerValue: "1",
			inputNote: &model.Note{
				Title: "test_title",
				Info:  "test_info",
			},
			mockBehavior: func(s *mock_services.MockNoteService, note *model.Note, userID string) {
				// service response
				s.EXPECT().CreateNote(note, userID).Return(nil, e.New("some db Err"))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"desc":"some db Err","error":"db response error"}
`,
			testName: "test-5-Service:Db resp Err",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			// Init mock controller
			c := gomock.NewController(t)
			defer c.Finish()
			// Init mock service
			noteSrv := mock_services.NewMockNoteService(c)
			testCase.mockBehavior(noteSrv, testCase.inputNote, testCase.headerValue)
			// Init testing logger with "fatal" level (5)
			logger := l.NewLogger(&config.Config{Logger: config.Logger{LogLevel: 5}})
			loggingMiddleware := l.NewLoggerMiddleware(logger)
			// Init service
			service := &services.Services{Note: noteSrv}
			handler := NewHandler(service, loggingMiddleware)
			// Test server
			router := httprouter.New()
			router.POST(utils.NotesURL, handler.logMiddleware(handler.CreateNote))
			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, utils.NotesURL, bytes.NewBufferString(testCase.inputJson))
			req.Header.Set(testCase.headerName, testCase.headerValue)
			// Make Request
			router.ServeHTTP(w, req)
			// Assert
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			require.Equal(t, testCase.expectedResponse, w.Body.String())
		})
	}
}

func TestHandler_GetNoteByID(t *testing.T) {
	// Init mock func obj
	type mockBehavior func(s *mock_services.MockNoteService, noteID, userID string)

	testTable := []struct {
		headerName         string
		headerValue        string
		inputNote          string
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
		testName           string
	}{
		{
			headerName: "user_id",
			// service request
			headerValue: "1",
			inputNote:   "1",
			mockBehavior: func(s *mock_services.MockNoteService, noteID, userID string) {
				// service response
				outputNote := &dto.NoteResp{
					Title: "test_title",
					Info:  "test_info",
				}
				s.EXPECT().GetNoteByID(noteID, userID).Return(outputNote, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: `{"title":"test_title","info":"test_info"}
`,
			testName: "test-1-Handler:OK",
		},
		{
			headerName: "user_id",
			// service request
			headerValue: "1",
			inputNote:   "1",
			mockBehavior: func(s *mock_services.MockNoteService, noteID, userID string) {
				// service response
				s.EXPECT().GetNoteByID(noteID, userID).Return(nil, e.New(errors.ErrDBNotExists))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"error":"No note with id '1'"}
`,
			testName: "test-2-Service:Note not found",
		},
		{
			headerName: "user_id",
			// service request
			headerValue: "1",
			inputNote:   "1",
			mockBehavior: func(s *mock_services.MockNoteService, noteID, userID string) {
				// service response
				s.EXPECT().GetNoteByID(noteID, userID).Return(nil, e.New("some db Err"))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"desc":"some db Err","error":"db response error"}
`,
			testName: "test-3-Service:Db resp Err",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			// Init mock controller
			c := gomock.NewController(t)
			defer c.Finish()
			// Init mock service
			noteSrv := mock_services.NewMockNoteService(c)
			testCase.mockBehavior(noteSrv, testCase.inputNote, testCase.headerValue)
			// Init testing logger with "fatal" level (5)
			logger := l.NewLogger(&config.Config{Logger: config.Logger{LogLevel: 5}})
			loggingMiddleware := l.NewLoggerMiddleware(logger)
			// Init service
			service := &services.Services{Note: noteSrv}
			handler := NewHandler(service, loggingMiddleware)
			// Test server
			router := httprouter.New()
			router.GET(utils.NoteURL, handler.logMiddleware(handler.GetNoteByID))
			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/notes/%s", testCase.inputNote), nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)
			// Make Request
			router.ServeHTTP(w, req)
			// Assert
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			require.Equal(t, testCase.expectedResponse, w.Body.String())
		})
	}
}

func TestHandler_GetAllNotesByUser(t *testing.T) {
	// Init mock func obj
	type mockBehavior func(s *mock_services.MockNoteService, userID string)

	testTable := []struct {
		headerName         string
		headerValue        string
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
		testName           string
	}{
		{
			headerName: "user_id",
			// service request
			headerValue: "1",
			mockBehavior: func(s *mock_services.MockNoteService, userID string) {
				// service response
				outputTag := []dto.NotesResp{
					{
						ID:    "1",
						Title: "test_title1",
						Info:  "test_info1",
					},
				}
				s.EXPECT().GetAllNotesByUser(userID).Return(outputTag, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: `[{"id":"1","title":"test_title1","info":"test_info1"}]
`,
			testName: "test-1-Handler:OK",
		},
		{
			headerName: "user_id",
			// service request
			headerValue: "1",
			mockBehavior: func(s *mock_services.MockNoteService, userID string) {
				// service response
				s.EXPECT().GetAllNotesByUser(userID).Return(nil, nil)
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"error":"no notes"}
`,
			testName: "test-2-Service:Tags not found",
		},
		{
			headerName: "user_id",
			// service request
			headerValue: "1",
			mockBehavior: func(s *mock_services.MockNoteService, userID string) {
				// service response
				s.EXPECT().GetAllNotesByUser(userID).Return(nil, e.New("some db Err"))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"desc":"some db Err","error":"db response error"}
`,
			testName: "test-3-Service:Db resp Err",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			// Init mock controller
			c := gomock.NewController(t)
			defer c.Finish()
			// Init mock service
			noteSrv := mock_services.NewMockNoteService(c)
			testCase.mockBehavior(noteSrv, testCase.headerValue)
			// Init testing logger with "fatal" level (5)
			logger := l.NewLogger(&config.Config{Logger: config.Logger{LogLevel: 5}})
			loggingMiddleware := l.NewLoggerMiddleware(logger)
			// Init service
			service := &services.Services{Note: noteSrv}
			handler := NewHandler(service, loggingMiddleware)
			// Test server
			router := httprouter.New()
			router.GET(utils.NotesURL, handler.logMiddleware(handler.GetAllNotesByUser))
			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, utils.NotesURL, nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)
			// Make Request
			router.ServeHTTP(w, req)
			// Assert
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			require.Equal(t, testCase.expectedResponse, w.Body.String())
		})
	}
}

func TestHandler_UpdateNote(t *testing.T) {
	// Init mock func obj
	type mockBehavior func(s *mock_services.MockNoteService, note *dto.NoteUpdate, noteID, userID string)
	noteTitle := "test_title"
	noteInfo := "test_info"

	testTable := []struct {
		inputJson          string
		headerName         string
		headerValue        string
		noteID             string
		inputNote          *dto.NoteUpdate
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
		testName           string
	}{
		{
			inputJson: `{
				"title": "test_title",
				"info": "test_info"
			}`,
			headerName: "user_id",
			// service request
			headerValue: "1",
			noteID:      "1",
			inputNote: &dto.NoteUpdate{
				Title: &noteTitle,
				Info:  &noteInfo,
			},
			mockBehavior: func(s *mock_services.MockNoteService, note *dto.NoteUpdate, noteID, userID string) {
				// service response
				s.EXPECT().GetNoteByID(noteID, userID).Return(nil, nil)
				s.EXPECT().UpdateNote(note, noteID).Return(nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: `{"Updated note with id":"1"}
`,
			testName: "test-1-Handler:OK",
		},
		{
			inputJson: `{
				"title": "test_title",
				"info": "test_info
			}`,
			headerName: "user_id",
			// service request
			headerValue:        "1",
			noteID:             "1",
			inputNote:          &dto.NoteUpdate{},
			mockBehavior:       func(s *mock_services.MockNoteService, note *dto.NoteUpdate, noteID, userID string) {},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `invalid character '\n' in string literal
`,
			testName: "test-2-Handler:Bad JSON",
		},
		{
			inputJson:  `{}`,
			headerName: "user_id",
			// service request
			headerValue: "1",
			noteID:      "1",
			inputNote:   &dto.NoteUpdate{},
			mockBehavior: func(s *mock_services.MockNoteService, note *dto.NoteUpdate, noteID, userID string) {
				// service response
				s.EXPECT().GetNoteByID(noteID, userID).Return(nil, e.New(errors.ErrDBNotExists))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"error":"No note with id '1'"}
`,
			testName: "test-3-Service:Note not found",
		},
		{
			inputJson: `{
				"title": "test_title",
				"info": "test_info"
			}`,
			headerName: "user_id",
			// service request
			headerValue: "1",
			noteID:      "1",
			inputNote: &dto.NoteUpdate{
				Title: &noteTitle,
				Info:  &noteInfo,
			},
			mockBehavior: func(s *mock_services.MockNoteService, note *dto.NoteUpdate, noteID, userID string) {
				// service response
				s.EXPECT().GetNoteByID(noteID, userID).Return(nil, nil)
				s.EXPECT().UpdateNote(note, noteID).Return(e.New("some db Err"))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"desc":"some db Err","error":"db response error"}
`,
			testName: "test-4-Service:Db resp Err",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			// Init mock controller
			c := gomock.NewController(t)
			defer c.Finish()
			// Init mock service
			noteSrv := mock_services.NewMockNoteService(c)
			testCase.mockBehavior(noteSrv, testCase.inputNote, testCase.noteID, testCase.headerValue)
			// Init testing logger with "fatal" level (5)
			logger := l.NewLogger(&config.Config{Logger: config.Logger{LogLevel: 5}})
			loggingMiddleware := l.NewLoggerMiddleware(logger)
			// Init service
			service := &services.Services{Note: noteSrv}
			handler := NewHandler(service, loggingMiddleware)
			// Test server
			router := httprouter.New()
			router.PUT(utils.NoteURL, handler.logMiddleware(handler.UpdateNote))
			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/notes/%s", testCase.noteID), bytes.NewBufferString(testCase.inputJson))
			req.Header.Set(testCase.headerName, testCase.headerValue)
			// Make Request
			router.ServeHTTP(w, req)
			// Assert
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			require.Equal(t, testCase.expectedResponse, w.Body.String())
		})
	}
}

func TestHandler_DeleteNote(t *testing.T) {
	// Init mock func obj
	type mockBehavior func(s *mock_services.MockNoteService, noteID, userID string)

	testTable := []struct {
		headerName         string
		headerValue        string
		inputNote          string
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
		testName           string
	}{
		{
			headerName: "user_id",
			// service request
			headerValue: "1",
			inputNote:   "1",
			mockBehavior: func(s *mock_services.MockNoteService, noteID, userID string) {
				// service response
				s.EXPECT().GetNoteByID(noteID, userID).Return(nil, nil)
				s.EXPECT().DeleteNote(noteID, userID).Return(1, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: `{"Deleted note with id":1}
`,
			testName: "test-1-Handler:OK",
		},
		{
			headerName: "user_id",
			// service request
			headerValue: "1",
			inputNote:   "1",
			mockBehavior: func(s *mock_services.MockNoteService, noteID, userID string) {
				// service response
				s.EXPECT().GetNoteByID(noteID, userID).Return(nil, e.New(errors.ErrDBNotExists))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"error":"No note with id '1'"}
`,
			testName: "test-2-Service:Note not found",
		},
		{
			headerName: "user_id",
			// service request
			headerValue: "1",
			inputNote:   "1",
			mockBehavior: func(s *mock_services.MockNoteService, noteID, userID string) {
				// service response
				s.EXPECT().GetNoteByID(noteID, userID).Return(nil, nil)
				s.EXPECT().DeleteNote(noteID, userID).Return(0, e.New("some db Err"))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"desc":"some db Err","error":"db response error"}
`,
			testName: "test-3-Service:Db resp Err",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			// Init mock controller
			c := gomock.NewController(t)
			defer c.Finish()
			// Init mock service
			noteSrv := mock_services.NewMockNoteService(c)
			testCase.mockBehavior(noteSrv, testCase.inputNote, testCase.headerValue)
			// Init testing logger with "fatal" level (5)
			logger := l.NewLogger(&config.Config{Logger: config.Logger{LogLevel: 5}})
			loggingMiddleware := l.NewLoggerMiddleware(logger)
			// Init service
			service := &services.Services{Note: noteSrv}
			handler := NewHandler(service, loggingMiddleware)
			// Test server
			router := httprouter.New()
			router.DELETE(utils.NoteURL, handler.logMiddleware(handler.DeleteNote))
			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/notes/%s", testCase.inputNote), nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)
			// Make Request
			router.ServeHTTP(w, req)
			// Assert
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			require.Equal(t, testCase.expectedResponse, w.Body.String())
		})
	}
}
