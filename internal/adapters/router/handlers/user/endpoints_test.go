package user

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

func TestHandler_RegisterUser(t *testing.T) {
	// Init mock func obj
	type mockBehavior func(s *mock_services.MockUserAuthService, user *model.User)

	testTable := []struct {
		inputJson          string
		inputUser          *model.User
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
		testName           string
	}{
		{
			inputJson: `{
				"username": "test_name",
				"password": "test_password"
			}`,
			// service request
			inputUser: &model.User{
				Username: "test_name",
				Password: "test_password",
			},
			mockBehavior: func(s *mock_services.MockUserAuthService, user *model.User) {
				// service response
				outputUser := &model.User{
					ID:       "1",
					Username: "test_name",
					Password: "test_password",
				}
				s.EXPECT().RegisterUser(user).Return(outputUser, nil)
			},
			expectedStatusCode: http.StatusCreated,
			expectedResponse: `{"Created new user 'test_name' with id":"1"}
`,
			testName: "test-1-Handler:OK",
		},
		{
			inputJson: `{
				"username": "test_name
			}`,
			inputUser:          &model.User{},
			mockBehavior:       func(s *mock_services.MockUserAuthService, user *model.User) {},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `invalid character '\n' in string literal
`,
			testName: "test-2-Handler:Bad JSON",
		},
		{
			inputJson: `{
				"username": "test_name"
			}`,
			inputUser:          &model.User{},
			mockBehavior:       func(s *mock_services.MockUserAuthService, user *model.User) {},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `Key: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag
`,
			testName: "test-3-Handler:Validation Err",
		},
		{
			inputJson: `{
				"username": "test_name",
				"password": "test_password"
			}`,
			// service request
			inputUser: &model.User{
				Username: "test_name",
				Password: "test_password",
			},
			mockBehavior: func(s *mock_services.MockUserAuthService, user *model.User) {
				// service response
				outputUser := &model.User{}
				s.EXPECT().RegisterUser(user).Return(outputUser, e.New(errors.ErrDBDuplicate))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"error":"user 'test_name' is already exists"}
`,
			testName: "test-4-Service:User is already exists Err",
		},
		{
			inputJson: `{
				"username": "test_name",
				"password": "test_password"
			}`,
			// service request
			inputUser: &model.User{
				Username: "test_name",
				Password: "test_password",
			},
			mockBehavior: func(s *mock_services.MockUserAuthService, user *model.User) {
				// service response
				outputUser := &model.User{}
				s.EXPECT().RegisterUser(user).Return(outputUser, e.New(errors.ErrDBNotExists))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"error":"No user with name 'test_name'"}
`,
			testName: "test-5-Service:User not found",
		},
		{
			inputJson: `{
				"username": "test_name",
				"password": "test_password"
			}`,
			// service request
			inputUser: &model.User{
				Username: "test_name",
				Password: "test_password",
			},
			mockBehavior: func(s *mock_services.MockUserAuthService, user *model.User) {
				// service response
				outputUser := &model.User{}
				s.EXPECT().RegisterUser(user).Return(outputUser, e.New("some db Err"))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"desc":"some db Err","error":"db response error"}
`,
			testName: "test-6-Service:Db resp Err",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			// Init mock controller
			c := gomock.NewController(t)
			defer c.Finish()
			// Init mock service
			auth := mock_services.NewMockUserAuthService(c)
			testCase.mockBehavior(auth, testCase.inputUser)
			// Init testing logger with "fatal" level (5)
			logger := l.NewLogger(&config.Config{Logger: config.Logger{LogLevel: 5}})
			loggingMiddleware := l.NewLoggerMiddleware(logger)
			// Init service
			service := &services.Services{Auth: auth}
			handler := NewHandler(service, loggingMiddleware)
			// Test server
			router := httprouter.New()
			router.POST(utils.Register, handler.LogMiddleware(handler.RegisterUser))
			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, utils.Register, bytes.NewBufferString(testCase.inputJson))
			// Make Request
			router.ServeHTTP(w, req)
			// Assert
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			require.Equal(t, testCase.expectedResponse, w.Body.String())
		})
	}
}

func TestHandler_GenerateToken(t *testing.T) {
	// Init mock func obj
	type mockBehavior func(s *mock_services.MockUserAuthService, userName, password string)

	testTable := []struct {
		inputJson          string
		userName           string
		password           string
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
		testName           string
	}{
		{
			inputJson: `{
				"username": "test_name",
				"password": "test_password"
			}`,
			// service request
			userName: "test_name",
			password: "test_password",
			mockBehavior: func(s *mock_services.MockUserAuthService, userName, password string) {
				// service response
				outputToken := "generatedToken"
				s.EXPECT().GenerateToken(userName, password).Return(outputToken, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: `{"token":"Bearer generatedToken"}
`,
			testName: "test-1-Handler:OK",
		},
		{
			inputJson: `{
				"username": "test_name
			}`,
			mockBehavior:       func(s *mock_services.MockUserAuthService, userName, password string) {},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `invalid character '\n' in string literal
`,
			testName: "test-2-Handler:Bad JSON",
		},
		{
			inputJson: `{
						"username": "test_name"
			}`,
			mockBehavior:       func(s *mock_services.MockUserAuthService, userName, password string) {},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `Key: 'UserAuth.Password' Error:Field validation for 'Password' failed on the 'required' tag
`,
			testName: "test-3-Handler:Validation Err",
		},
		{
			inputJson: `{
				"username": "test_name",
				"password": "test_password"
			}`,
			// service request
			userName: "test_name",
			password: "test_password",
			mockBehavior: func(s *mock_services.MockUserAuthService, userName, password string) {
				// service response
				s.EXPECT().GenerateToken(userName, password).Return("", e.New(errors.ErrDBNotExists))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"error":"No user with name 'test_name'"}
`,
			testName: "test-4-Service:User not found",
		},
		{
			inputJson: `{
				"username": "test_name",
				"password": "test_password"
			}`,
			// service request
			userName: "test_name",
			password: "test_password",
			mockBehavior: func(s *mock_services.MockUserAuthService, userName, password string) {
				// service response
				s.EXPECT().GenerateToken(userName, password).Return("", e.New("some db Err"))
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
			auth := mock_services.NewMockUserAuthService(c)
			testCase.mockBehavior(auth, testCase.userName, testCase.password)
			// Init testing logger with "fatal" level (5)
			logger := l.NewLogger(&config.Config{Logger: config.Logger{LogLevel: 5}})
			loggingMiddleware := l.NewLoggerMiddleware(logger)
			// Init service
			service := &services.Services{Auth: auth}
			handler := NewHandler(service, loggingMiddleware)
			// Test server
			router := httprouter.New()
			router.POST(utils.Login, handler.LogMiddleware(handler.GenerateToken))
			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, utils.Login, bytes.NewBufferString(testCase.inputJson))
			// Make Request
			router.ServeHTTP(w, req)
			// Assert
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			require.Equal(t, testCase.expectedResponse, w.Body.String())
		})
	}
}

func TestHandler_GetUserByID(t *testing.T) {
	// Init mock func obj
	type mockBehavior func(s *mock_services.MockUserService, userID string)

	testTable := []struct {
		userID             string
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
		testName           string
	}{
		{
			// service request
			userID: "1",
			mockBehavior: func(s *mock_services.MockUserService, userID string) {
				// service response
				outputUser := &dto.UserResp{
					ID:       "1",
					Username: "test_name",
				}
				s.EXPECT().GetUserByID(userID).Return(outputUser, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: `{"id":"1","username":"test_name"}
`,
			testName: "test-1-Handler:OK",
		},
		{
			// service request
			userID: "2",
			mockBehavior: func(s *mock_services.MockUserService, userID string) {
				// service response
				s.EXPECT().GetUserByID(userID).Return(nil, e.New(errors.ErrDBNotExists))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"error":"No user with id '2'"}
`,
			testName: "test-2-Service:User not found",
		},
		{
			// service request
			userID: "3",
			mockBehavior: func(s *mock_services.MockUserService, userID string) {
				// service response
				s.EXPECT().GetUserByID(userID).Return(nil, e.New("some db Err"))
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
			userSrv := mock_services.NewMockUserService(c)
			testCase.mockBehavior(userSrv, testCase.userID)
			// Init testing logger with "fatal" level (5)
			logger := l.NewLogger(&config.Config{Logger: config.Logger{LogLevel: 5}})
			loggingMiddleware := l.NewLoggerMiddleware(logger)
			// Init service
			service := &services.Services{User: userSrv}
			handler := NewHandler(service, loggingMiddleware)
			// Test server
			router := httprouter.New()
			router.GET(utils.UserURL, handler.LogMiddleware(handler.GetUserByID))
			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", testCase.userID), nil)
			// Make Request
			router.ServeHTTP(w, req)
			// Assert
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			require.Equal(t, testCase.expectedResponse, w.Body.String())
		})
	}
}

func TestHandler_GetAllUsers(t *testing.T) {
	// Init mock func obj
	type mockBehavior func(s *mock_services.MockUserService)

	testTable := []struct {
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
		testName           string
	}{
		{
			mockBehavior: func(s *mock_services.MockUserService) {
				// service response
				outputUser := []dto.UserResp{
					{
						ID:       "1",
						Username: "test_name1",
					},
					{
						ID:       "2",
						Username: "test_name2",
					},
				}
				s.EXPECT().GetAllUsers().Return(outputUser, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: `[{"id":"1","username":"test_name1"},{"id":"2","username":"test_name2"}]
`,
			testName: "test-1-Handler:OK",
		},
		{
			mockBehavior: func(s *mock_services.MockUserService) {
				// service response
				s.EXPECT().GetAllUsers().Return(nil, nil)
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"error":"no users"}
`,
			testName: "test-2-Service:Users not found",
		},
		{
			mockBehavior: func(s *mock_services.MockUserService) {
				// service response
				s.EXPECT().GetAllUsers().Return(nil, e.New("some db Err"))
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
			userSrv := mock_services.NewMockUserService(c)
			testCase.mockBehavior(userSrv)
			// Init testing logger with "fatal" level (5)
			logger := l.NewLogger(&config.Config{Logger: config.Logger{LogLevel: 5}})
			loggingMiddleware := l.NewLoggerMiddleware(logger)
			// Init service
			service := &services.Services{User: userSrv}
			handler := NewHandler(service, loggingMiddleware)
			// Test server
			router := httprouter.New()
			router.GET(utils.UsersURL, handler.LogMiddleware(handler.GetAllUsers))
			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, utils.UsersURL, nil)
			// Make Request
			router.ServeHTTP(w, req)
			// Assert
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			require.Equal(t, testCase.expectedResponse, w.Body.String())
		})
	}
}

func TestHandler_UpdateUser(t *testing.T) {
	// Init mock func obj
	type mockBehavior func(s *mock_services.MockUserService, user *dto.UserUpdate, userID string)
	userName := "test_name"
	password := "test_password"

	testTable := []struct {
		inputJson          string
		inputUser          *dto.UserUpdate
		userID             string
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
		testName           string
	}{
		{
			inputJson: `{
				"username": "test_name",
				"password": "test_password"
			}`,
			// service request
			inputUser: &dto.UserUpdate{
				Username: &userName,
				Password: &password,
			},
			userID: "1",
			mockBehavior: func(s *mock_services.MockUserService, user *dto.UserUpdate, userID string) {
				s.EXPECT().GetUserByID(userID).Return(nil, nil)
				s.EXPECT().UpdateUser(user, userID).Return(nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: `{"Updated user with id":"1"}
`,
			testName: "test-1-Handler:OK",
		},
		{
			inputJson: `{
				"username": "test_name
			}`,
			userID:             "1",
			inputUser:          &dto.UserUpdate{},
			mockBehavior:       func(s *mock_services.MockUserService, user *dto.UserUpdate, userID string) {},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `invalid character '\n' in string literal
`,
			testName: "test-2-Handler:Bad JSON",
		},
		{
			inputJson: `{}`,
			// service request
			userID: "2",
			mockBehavior: func(s *mock_services.MockUserService, user *dto.UserUpdate, userID string) {
				// service response
				s.EXPECT().GetUserByID(userID).Return(nil, e.New(errors.ErrDBNotExists))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"error":"No user with id '2'"}
`,
			testName: "test-3-Service:User not found",
		},
		{
			inputJson: `{
				"username": "test_name",
				"password": "test_password"
			}`,
			// service request
			inputUser: &dto.UserUpdate{
				Username: &userName,
				Password: &password,
			},
			userID: "1",
			mockBehavior: func(s *mock_services.MockUserService, user *dto.UserUpdate, userID string) {
				s.EXPECT().GetUserByID(userID).Return(nil, nil)
				s.EXPECT().UpdateUser(user, userID).Return(e.New("some db Err"))
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
			userSrv := mock_services.NewMockUserService(c)
			testCase.mockBehavior(userSrv, testCase.inputUser, testCase.userID)
			// Init testing logger with "fatal" level (5)
			logger := l.NewLogger(&config.Config{Logger: config.Logger{LogLevel: 5}})
			loggingMiddleware := l.NewLoggerMiddleware(logger)
			// Init service
			service := &services.Services{User: userSrv}
			handler := NewHandler(service, loggingMiddleware)
			// Test server
			router := httprouter.New()
			router.PUT(utils.UserURL, handler.LogMiddleware(handler.UpdateUser))
			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/users/%s", testCase.userID), bytes.NewBufferString(testCase.inputJson))
			// Make Request
			router.ServeHTTP(w, req)
			// Assert
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			require.Equal(t, testCase.expectedResponse, w.Body.String())
		})
	}
}

func TestHandler_DeleteUser(t *testing.T) {
	// Init mock func obj
	type mockBehavior func(s *mock_services.MockUserService, id string)

	testTable := []struct {
		userID             string
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
		testName           string
	}{
		{
			// service request
			userID: "1",
			mockBehavior: func(s *mock_services.MockUserService, id string) {
				// service response
				s.EXPECT().GetUserByID(id).Return(nil, nil)
				s.EXPECT().DeleteUser(id).Return(1, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: `{"Deleted user with id":1}
`,
			testName: "test-1-Handler:OK",
		},
		{
			// service request
			userID: "2",
			mockBehavior: func(s *mock_services.MockUserService, id string) {
				// service response
				s.EXPECT().GetUserByID(id).Return(nil, e.New(errors.ErrDBNotExists))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"error":"No user with id '2'"}
`,
			testName: "test-2-Service:User not found",
		},
		{
			// service request
			userID: "1",
			mockBehavior: func(s *mock_services.MockUserService, id string) {
				// service response
				s.EXPECT().GetUserByID(id).Return(nil, nil)
				s.EXPECT().DeleteUser(id).Return(0, e.New("some db Err"))
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
			userSrv := mock_services.NewMockUserService(c)
			testCase.mockBehavior(userSrv, testCase.userID)
			// Init testing logger with "fatal" level (5)
			logger := l.NewLogger(&config.Config{Logger: config.Logger{LogLevel: 5}})
			loggingMiddleware := l.NewLoggerMiddleware(logger)
			// Init service
			service := &services.Services{User: userSrv}
			handler := NewHandler(service, loggingMiddleware)
			// Test server
			router := httprouter.New()
			router.DELETE(utils.UserURL, handler.LogMiddleware(handler.DeleteUser))
			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/users/%s", testCase.userID), nil)
			// Make Request
			router.ServeHTTP(w, req)
			// Assert
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			require.Equal(t, testCase.expectedResponse, w.Body.String())
		})
	}
}
