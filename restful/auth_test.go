package restful

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"web-server/exceptions"
	mockservice "web-server/mock/service"
	"web-server/restful/request"
	"web-server/restful/response"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type unitTest struct {
	name string
	run  func()
}

type AuthControllerTestSuite struct {
	suite.Suite
	ctx             context.Context
	engine          *gin.Engine
	writer          *httptest.ResponseRecorder
	mockAuthService *mockservice.MockAuthService
}

func (suite *AuthControllerTestSuite) SetupTest() {
	writer := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(writer)
	gin.SetMode(gin.TestMode)
	suite.writer = writer
	suite.ctx = ctx
	suite.engine = engine
	suite.mockAuthService = new(mockservice.MockAuthService)
}

func (suite *AuthControllerTestSuite) TestLoginHandler() {
	tests := []unitTest{
		unitTest{
			name: "should response StatusUnprocessableEntity if parse json failed",
			run: func() {
				NewAuthController(suite.engine.Group("api"), mockservice.NewMockAuthServiceWithContext(suite.mockAuthService))
				request := httptest.NewRequest(http.MethodPost, "/api/v1/login", nil)
				suite.engine.ServeHTTP(suite.writer, request)
				suite.Equal(suite.writer.Code, http.StatusUnprocessableEntity, "should return error if parse json failed")
			},
		},
		unitTest{
			name: "should response StatusNotFound if email does not exist",
			run: func() {
				NewAuthController(suite.engine.Group("api"), mockservice.NewMockAuthServiceWithContext(suite.mockAuthService))
				loginRequest := request.EmailLoginRequest{
					Email:    "abc@def.com",
					Password: "123456",
				}
				suite.mockAuthService.On("AuthenticateByEmail", "abc@def.com", "123456").Return("", exceptions.UserNotExists)
				json, _ := json.Marshal(loginRequest)
				request := httptest.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewReader(json))
				suite.engine.ServeHTTP(suite.writer, request)
				suite.Equal(suite.writer.Code, http.StatusNotFound, "should return error if user does not exist")
			},
		},
		unitTest{
			name: "should response StatusUnauthorized if authenticate failed",
			run: func() {
				NewAuthController(suite.engine.Group("api"), mockservice.NewMockAuthServiceWithContext(suite.mockAuthService))
				loginRequest := request.EmailLoginRequest{
					Email:    "abc@def.com",
					Password: "123456",
				}
				suite.mockAuthService.On("AuthenticateByEmail", "abc@def.com", "123456").Return("", exceptions.AuthFailed)
				loginRequestJson, _ := json.Marshal(loginRequest)
				request := httptest.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewReader(loginRequestJson))
				suite.engine.ServeHTTP(suite.writer, request)
				suite.Equal(suite.writer.Code, http.StatusUnauthorized, "should return error if auth failed")
				body, _ := io.ReadAll(suite.writer.Body)
				authFailedJson, _ := json.Marshal(exceptions.AuthFailed)
				suite.Equal(body, authFailedJson, "should return authFailedJson")
			},
		},
		unitTest{
			name: "should response Ok if authenticate succeeded",
			run: func() {
				NewAuthController(suite.engine.Group("api"), mockservice.NewMockAuthServiceWithContext(suite.mockAuthService))
				loginRequest := request.EmailLoginRequest{
					Email:    "abc@def.com",
					Password: "123456",
				}
				suite.mockAuthService.On("AuthenticateByEmail", "abc@def.com", "123456").Return("123", nil)
				loginRequestJson, _ := json.Marshal(loginRequest)
				request := httptest.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewReader(loginRequestJson))
				suite.engine.ServeHTTP(suite.writer, request)
				suite.Equal(suite.writer.Code, http.StatusOK, "should not return error")
				body, _ := io.ReadAll(suite.writer.Body)
				loginResponse, _ := json.Marshal(response.NewLoginResponse("123"))
				suite.Equal(body, loginResponse, "should return login response")
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, test.run)
	}
}

func TestAuthControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthControllerTestSuite))
}
