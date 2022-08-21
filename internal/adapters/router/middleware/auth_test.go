package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/require"

	"web/internal/adapters/router/handlers/user"
	"web/internal/config"
	"web/internal/domain/services"
	"web/internal/domain/services/mocks"
	l "web/pkg/logger"
)

func TestCheckToken(t *testing.T) {
	// Init mock func obj
	type mockBehavior func(s *mock_services.MockUserAuthService, token string)

	testTable := []struct {
		headerName         string
		headerValue        string
		token              string
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedUserID     string
		expectedResponse   string
		testName           string
	}{
		{
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mock_services.MockUserAuthService, token string) {
				s.EXPECT().ParseToken(token).Return("1", nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedUserID:     "1",
			testName:           "test-1-OK",
		},
		{
			headerName:         "",
			mockBehavior:       func(s *mock_services.MockUserAuthService, token string) {},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse: `{"error":"empty auth header"}
`,
			testName: "test-2-Empty auth header",
		},
		{
			headerName:         "Authorization",
			headerValue:        "Bearrrrrr token",
			mockBehavior:       func(s *mock_services.MockUserAuthService, token string) {},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse: `{"error":"invalid auth header"}
`,
			testName: "test-3-Invalid Bearer",
		},
		{
			headerName:         "Authorization",
			headerValue:        "Bearer ",
			mockBehavior:       func(s *mock_services.MockUserAuthService, token string) {},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse: `{"error":"empty token"}
`,
			testName: "test-4-Empty token",
		},
		{
			headerName:         "Authorization",
			headerValue:        "Bearer",
			mockBehavior:       func(s *mock_services.MockUserAuthService, token string) {},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse: `{"error":"invalid auth header"}
`,
			testName: "test-5-Invalid auth header",
		},
		{
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mock_services.MockUserAuthService, token string) {
				s.EXPECT().ParseToken(token).Return("1", errors.New("failed to parse token"))
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse: `{"error":"failed to parse token"}
`,
			testName: "test-6-Service Err",
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			// Init mock controller
			c := gomock.NewController(t)
			defer c.Finish()
			// Init mock service
			auth := mock_services.NewMockUserAuthService(c)
			testCase.mockBehavior(auth, testCase.token)
			// Init testing logger with "fatal" level (5)
			logger := l.NewLogger(&config.Config{Logger: config.Logger{LogLevel: 5}})
			loggingMiddleware := l.NewLoggerMiddleware(logger)
			// Init service
			service := &services.Services{Auth: auth}
			handler := user.NewHandler(service, loggingMiddleware)
			// actual user_id from r header
			var actualUserID string
			// Test server
			router := httprouter.New()
			router.GET("/protected",
				handler.LogMiddleware(
					CheckToken(
						func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
							actualUserID = r.Header.Get("user_id")
						},
						service.Auth),
				),
			)
			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/protected", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)
			// Make Request
			router.ServeHTTP(w, req)
			// Assert
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			require.Equal(t, testCase.expectedResponse, w.Body.String())
			require.Equal(t, testCase.expectedUserID, actualUserID)
		})
	}
}
