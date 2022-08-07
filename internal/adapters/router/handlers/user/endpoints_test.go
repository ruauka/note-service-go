package user

import (
	"bytes"
	e "errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/require"

	"web/internal/config"
	"web/internal/domain/enteties/model"
	"web/internal/domain/errors"
	"web/internal/domain/services"
	"web/internal/domain/services/mocks"
	"web/internal/utils"
	l "web/pkg/logger"
)

func TestHandler_RegisterUser(t *testing.T) {
	// Init mock func obj
	type mockBehavior func(s *mock_services.MockUserAuthService, user model.User)

	testTable := []struct {
		inputJson          string
		inputUser          model.User
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
			inputUser: model.User{
				Username: "test_name",
				Password: "test_password",
			},
			mockBehavior: func(s *mock_services.MockUserAuthService, user model.User) {
				// service response
				outputUser := model.User{
					ID:       "1",
					Username: "test_name",
					Password: "test_password",
				}
				s.EXPECT().RegisterUser(&user).Return(&outputUser, nil)
			},
			expectedStatusCode: http.StatusCreated,
			expectedResponse: `{"Created new user 'test_name' with id":"1"}
`,
			testName: "test-1-OK",
		},
		{
			inputJson: `{
				"username": "test_name
			}`,
			inputUser:          model.User{},
			mockBehavior:       func(s *mock_services.MockUserAuthService, user model.User) {},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `invalid character '\n' in string literal
`,
			testName: "test-2-Bad JSON",
		},
		{
			inputJson: `{
				"username": "test_name"
			}`,
			inputUser:          model.User{},
			mockBehavior:       func(s *mock_services.MockUserAuthService, user model.User) {},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `Key: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag
`,
			testName: "test-3-Validation Err",
		},
		{
			inputJson: `{
				"username": "test_name",
				"password": "test_password"
			}`,
			// service request
			inputUser: model.User{
				Username: "test_name",
				Password: "test_password",
			},
			mockBehavior: func(s *mock_services.MockUserAuthService, user model.User) {
				// service response
				outputUser := model.User{}
				s.EXPECT().RegisterUser(&user).Return(&outputUser, e.New(errors.ErrDbDuplicate))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"error":"user 'test_name' is already exists"}
`,
			testName: "test-4-Service-user is already exists Err",
		},
		{
			inputJson: `{
				"username": "test_name",
				"password": "test_password"
			}`,
			// service request
			inputUser: model.User{
				Username: "test_name",
				Password: "test_password",
			},
			mockBehavior: func(s *mock_services.MockUserAuthService, user model.User) {
				// service response
				outputUser := model.User{}
				s.EXPECT().RegisterUser(&user).Return(&outputUser, e.New(errors.ErrDbNotExists))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"error":"No user with name 'test_name'"}
`,
			testName: "test-5-Service-user not found",
		},
		{
			inputJson: `{
				"username": "test_name",
				"password": "test_password"
			}`,
			// service request
			inputUser: model.User{
				Username: "test_name",
				Password: "test_password",
			},
			mockBehavior: func(s *mock_services.MockUserAuthService, user model.User) {
				// service response
				outputUser := model.User{}
				s.EXPECT().RegisterUser(&user).Return(&outputUser, e.New("some db Err"))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: `{"desc":"some db Err","error":"db response error"}
`,
			testName: "test-6-Service-db resp Err",
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
