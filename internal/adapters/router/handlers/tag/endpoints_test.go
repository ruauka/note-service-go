package tag

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

func TestHandler_CreateTag(t *testing.T) {
	// Init mock func obj
	type mockBehavior func(s *mock_services.MockTagService, tag *model.Tag, userID string)

	testTable := []struct {
		headerName         string
		headerValue        string
		inputJson          string
		inputTag           *model.Tag
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
		testName           string
	}{
		{
			headerName: "user_id",
			inputJson: `{
				"tagname": "test_name"
			}`,
			// service request
			headerValue: "1",
			inputTag: &model.Tag{
				TagName: "test_name",
			},
			mockBehavior: func(s *mock_services.MockTagService, tag *model.Tag, userID string) {
				// service response
				outputTag := &model.Tag{
					ID:      "1",
					TagName: "test_name",
				}
				s.EXPECT().CreateTag(tag, userID).Return(outputTag, nil)
			},
			expectedStatusCode: http.StatusCreated,
			expectedResponse: `{"Created tag 'test_name' with id":"1"}
`,
			testName: "test-1-Handler:OK",
		},
		{
			inputJson: `{
				"tagname": "test_name
			}`,
			inputTag:           &model.Tag{},
			mockBehavior:       func(s *mock_services.MockTagService, tag *model.Tag, userID string) {},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `invalid character '\n' in string literal
`,
			testName: "test-2-Handler:Bad JSON",
		},
		{
			inputJson:          `{}`,
			inputTag:           &model.Tag{},
			mockBehavior:       func(s *mock_services.MockTagService, tag *model.Tag, userID string) {},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `Key: 'Tag.TagName' Error:Field validation for 'TagName' failed on the 'required' tag
`,
			testName: "test-3-Handler:Validation Err",
		},
		{
			headerName: "user_id",
			inputJson: `{
				"tagname": "test_name"
			}`,
			// service request
			headerValue: "1",
			inputTag: &model.Tag{
				TagName: "test_name",
			},
			mockBehavior: func(s *mock_services.MockTagService, tag *model.Tag, userID string) {
				// service response
				s.EXPECT().CreateTag(tag, userID).Return(nil, e.New(errors.ErrDBDuplicate))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"error":"tag 'test_name' is already exists"}
`,
			testName: "test-4-Service:Tag is already exists Err",
		},
		{
			headerName: "user_id",
			inputJson: `{
				"tagname": "test_name"
			}`,
			// service request
			headerValue: "1",
			inputTag: &model.Tag{
				TagName: "test_name",
			},
			mockBehavior: func(s *mock_services.MockTagService, tag *model.Tag, userID string) {
				// service response
				s.EXPECT().CreateTag(tag, userID).Return(nil, e.New("some db Err"))
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
			tagSrv := mock_services.NewMockTagService(c)
			testCase.mockBehavior(tagSrv, testCase.inputTag, testCase.headerValue)
			// Init testing logger with "fatal" level (5)
			logger := l.NewLogger(&config.Config{Logger: config.Logger{LogLevel: 5}})
			loggingMiddleware := l.NewLoggerMiddleware(logger)
			// Init service
			service := &services.Services{Tag: tagSrv}
			handler := NewHandler(service, loggingMiddleware)
			// Test server
			router := httprouter.New()
			router.POST(utils.TagsURL, handler.logMiddleware(handler.CreateTag))
			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, utils.TagsURL, bytes.NewBufferString(testCase.inputJson))
			req.Header.Set(testCase.headerName, testCase.headerValue)
			// Make Request
			router.ServeHTTP(w, req)
			// Assert
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			require.Equal(t, testCase.expectedResponse, w.Body.String())
		})
	}
}

func TestHandler_GetTagByID(t *testing.T) {
	// Init mock func obj
	type mockBehavior func(s *mock_services.MockTagService, tagID, userID string)

	testTable := []struct {
		headerName         string
		headerValue        string
		inputTag           string
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
		testName           string
	}{
		{
			headerName: "user_id",
			// service request
			headerValue: "1",
			inputTag:    "1",
			mockBehavior: func(s *mock_services.MockTagService, tagID, userID string) {
				// service response
				outputTag := &dto.TagResp{
					TagName: "test_name",
				}
				s.EXPECT().GetTagByID(tagID, userID).Return(outputTag, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: `{"tagname":"test_name"}
`,
			testName: "test-1-Handler:OK",
		},
		{
			headerName: "user_id",
			// service request
			headerValue: "1",
			inputTag:    "1",
			mockBehavior: func(s *mock_services.MockTagService, tagID, userID string) {
				// service response
				s.EXPECT().GetTagByID(tagID, userID).Return(nil, e.New(errors.ErrDBNotExists))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"error":"No tag with id '1'"}
`,
			testName: "test-2-Service:Tag not found",
		},
		{
			headerName: "user_id",
			// service request
			headerValue: "1",
			inputTag:    "1",
			mockBehavior: func(s *mock_services.MockTagService, tagID, userID string) {
				// service response
				s.EXPECT().GetTagByID(tagID, userID).Return(nil, e.New("some db Err"))
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
			tagSrv := mock_services.NewMockTagService(c)
			testCase.mockBehavior(tagSrv, testCase.inputTag, testCase.headerValue)
			// Init testing logger with "fatal" level (5)
			logger := l.NewLogger(&config.Config{Logger: config.Logger{LogLevel: 5}})
			loggingMiddleware := l.NewLoggerMiddleware(logger)
			// Init service
			service := &services.Services{Tag: tagSrv}
			handler := NewHandler(service, loggingMiddleware)
			// Test server
			router := httprouter.New()
			router.GET(utils.TagURL, handler.logMiddleware(handler.GetTagByID))
			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/tags/%s", testCase.inputTag), nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)
			// Make Request
			router.ServeHTTP(w, req)
			// Assert
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			require.Equal(t, testCase.expectedResponse, w.Body.String())
		})
	}
}

func TestHandler_GetAllTagsByUser(t *testing.T) {
	// Init mock func obj
	type mockBehavior func(s *mock_services.MockTagService, userID string)

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
			mockBehavior: func(s *mock_services.MockTagService, userID string) {
				// service response
				outputTag := []dto.TagsResp{
					{
						ID:      "1",
						TagName: "test_name1",
					},
					{
						ID:      "2",
						TagName: "test_name2",
					},
				}
				s.EXPECT().GetAllTagsByUser(userID).Return(outputTag, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: `[{"id":"1","tagname":"test_name1"},{"id":"2","tagname":"test_name2"}]
`,
			testName: "test-1-Handler:OK",
		},
		{
			headerName: "user_id",
			// service request
			headerValue: "1",
			mockBehavior: func(s *mock_services.MockTagService, userID string) {
				// service response
				s.EXPECT().GetAllTagsByUser(userID).Return(nil, nil)
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"error":"no tags"}
`,
			testName: "test-2-Service:Tags not found",
		},
		{
			headerName: "user_id",
			// service request
			headerValue: "1",
			mockBehavior: func(s *mock_services.MockTagService, userID string) {
				// service response
				s.EXPECT().GetAllTagsByUser(userID).Return(nil, e.New("some db Err"))
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
			tagSrv := mock_services.NewMockTagService(c)
			testCase.mockBehavior(tagSrv, testCase.headerValue)
			// Init testing logger with "fatal" level (5)
			logger := l.NewLogger(&config.Config{Logger: config.Logger{LogLevel: 5}})
			loggingMiddleware := l.NewLoggerMiddleware(logger)
			// Init service
			service := &services.Services{Tag: tagSrv}
			handler := NewHandler(service, loggingMiddleware)
			// Test server
			router := httprouter.New()
			router.GET(utils.TagsURL, handler.logMiddleware(handler.GetAllTagsByUser))
			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, utils.TagsURL, nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)
			// Make Request
			router.ServeHTTP(w, req)
			// Assert
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			require.Equal(t, testCase.expectedResponse, w.Body.String())
		})
	}
}

func TestHandler_UpdateTag(t *testing.T) {
	// Init mock func obj
	type mockBehavior func(s *mock_services.MockTagService, tag *dto.TagUpdate, tagID, userID string)
	tagName := "test_name"

	testTable := []struct {
		inputJson          string
		headerName         string
		headerValue        string
		tagID              string
		inputTag           *dto.TagUpdate
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
		testName           string
	}{
		{
			inputJson: `{
				"tagname": "test_name"
			}`,
			headerName: "user_id",
			// service request
			headerValue: "1",
			tagID:       "1",
			inputTag: &dto.TagUpdate{
				TagName: &tagName,
			},
			mockBehavior: func(s *mock_services.MockTagService, tag *dto.TagUpdate, tagID, userID string) {
				// service response
				s.EXPECT().GetTagByID(tagID, userID).Return(nil, nil)
				s.EXPECT().UpdateTag(tag, tagID).Return(nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: `{"Updated tag with id":"1"}
`,
			testName: "test-1-Handler:OK",
		},
		{
			inputJson: `{
				"tagname": "test_name
			}`,
			headerName: "user_id",
			// service request
			headerValue:        "1",
			tagID:              "1",
			inputTag:           &dto.TagUpdate{},
			mockBehavior:       func(s *mock_services.MockTagService, tag *dto.TagUpdate, tagID, userID string) {},
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
			tagID:       "1",
			inputTag:    &dto.TagUpdate{},
			mockBehavior: func(s *mock_services.MockTagService, tag *dto.TagUpdate, tagID, userID string) {
				// service response
				s.EXPECT().GetTagByID(tagID, userID).Return(nil, e.New(errors.ErrDBNotExists))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"error":"No tag with id '1'"}
`,
			testName: "test-3-Service:Tag not found",
		},
		{
			inputJson: `{
				"tagname": "test_name"
			}`,
			headerName: "user_id",
			// service request
			headerValue: "1",
			tagID:       "1",
			inputTag: &dto.TagUpdate{
				TagName: &tagName,
			},
			mockBehavior: func(s *mock_services.MockTagService, tag *dto.TagUpdate, tagID, userID string) {
				// service response
				s.EXPECT().GetTagByID(tagID, userID).Return(nil, nil)
				s.EXPECT().UpdateTag(tag, tagID).Return(e.New("some db Err"))
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
			tagSrv := mock_services.NewMockTagService(c)
			testCase.mockBehavior(tagSrv, testCase.inputTag, testCase.tagID, testCase.headerValue)
			// Init testing logger with "fatal" level (5)
			logger := l.NewLogger(&config.Config{Logger: config.Logger{LogLevel: 5}})
			loggingMiddleware := l.NewLoggerMiddleware(logger)
			// Init service
			service := &services.Services{Tag: tagSrv}
			handler := NewHandler(service, loggingMiddleware)
			// Test server
			router := httprouter.New()
			router.PUT(utils.TagURL, handler.logMiddleware(handler.UpdateTag))
			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/tags/%s", testCase.tagID), bytes.NewBufferString(testCase.inputJson))
			req.Header.Set(testCase.headerName, testCase.headerValue)
			// Make Request
			router.ServeHTTP(w, req)
			// Assert
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			require.Equal(t, testCase.expectedResponse, w.Body.String())
		})
	}
}

func TestHandler_DeleteTag(t *testing.T) {
	// Init mock func obj
	type mockBehavior func(s *mock_services.MockTagService, tagID, userID string)

	testTable := []struct {
		headerName         string
		headerValue        string
		inputTag           string
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
		testName           string
	}{
		{
			headerName: "user_id",
			// service request
			headerValue: "1",
			inputTag:    "1",
			mockBehavior: func(s *mock_services.MockTagService, tagID, userID string) {
				// service response
				s.EXPECT().GetTagByID(tagID, userID).Return(nil, nil)
				s.EXPECT().DeleteTag(tagID, userID).Return(1, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: `{"Deleted tag with id":1}
`,
			testName: "test-1-Handler:OK",
		},
		{
			headerName: "user_id",
			// service request
			headerValue: "1",
			inputTag:    "1",
			mockBehavior: func(s *mock_services.MockTagService, tagID, userID string) {
				// service response
				s.EXPECT().GetTagByID(tagID, userID).Return(nil, e.New(errors.ErrDBNotExists))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"error":"No tag with id '1'"}
`,
			testName: "test-2-Service:Tag not found",
		},
		{
			headerName: "user_id",
			// service request
			headerValue: "1",
			inputTag:    "1",
			mockBehavior: func(s *mock_services.MockTagService, tagID, userID string) {
				// service response
				s.EXPECT().GetTagByID(tagID, userID).Return(nil, nil)
				s.EXPECT().DeleteTag(tagID, userID).Return(0, e.New("some db Err"))
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
			tagSrv := mock_services.NewMockTagService(c)
			testCase.mockBehavior(tagSrv, testCase.inputTag, testCase.headerValue)
			// Init testing logger with "fatal" level (5)
			logger := l.NewLogger(&config.Config{Logger: config.Logger{LogLevel: 5}})
			loggingMiddleware := l.NewLoggerMiddleware(logger)
			// Init service
			service := &services.Services{Tag: tagSrv}
			handler := NewHandler(service, loggingMiddleware)
			// Test server
			router := httprouter.New()
			router.DELETE(utils.TagURL, handler.logMiddleware(handler.DeleteTag))
			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/tags/%s", testCase.inputTag), nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)
			// Make Request
			router.ServeHTTP(w, req)
			// Assert
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			require.Equal(t, testCase.expectedResponse, w.Body.String())
		})
	}
}
