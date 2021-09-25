package restful

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"web-server/exceptions"
	mockservice "web-server/mock/service"
	"web-server/model"
	"web-server/restful/middlewares"
	"web-server/restful/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
	ctx             context.Context
	engine          *gin.Engine
	writer          *httptest.ResponseRecorder
	mockUserService *mockservice.MockUserService
	mockAuthService *mockservice.MockAuthService
	authChecker     middlewares.AuthChecker
}

func (suite *UserControllerTestSuite) SetupTest() {
	writer := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(writer)
	gin.SetMode(gin.TestMode)
	suite.writer = writer
	suite.ctx = ctx
	suite.engine = engine
	suite.mockUserService = new(mockservice.MockUserService)
	suite.mockAuthService = new(mockservice.MockAuthService)
	suite.authChecker = middlewares.NewAuthChecker(mockservice.NewMockAuthServiceWithContext(suite.mockAuthService))
}

func (suite *UserControllerTestSuite) TestGetUserByIdHandler() {
	tests := []unitTest{
		unitTest{
			name: "should return 404 of user is not found",
			run: func() {
				NewUserController(suite.engine.Group("api"), suite.authChecker, mockservice.NewMockUserServiceWithContext(suite.mockUserService))
				id := uuid.New()
				token := jwt.New(jwt.SigningMethodES256)
				tokenString, _ := token.SignedString("")
				request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/user/%s", id.String()), nil)
				header := http.Header{}
				header.Add("Authorization", fmt.Sprintf("Bearer %s", tokenString))
				request.Header = header
				suite.mockAuthService.On("Validate", mock.Anything, mock.Anything).Return(token, nil)
				suite.mockUserService.On("GetUserById", id).Return(model.User{}, exceptions.UserNotExists)

				suite.engine.ServeHTTP(suite.writer, request)
				suite.Equal(suite.writer.Code, http.StatusNotFound, "should return error if user does not exist")
			},
		},
		unitTest{
			name: "should return ok of user found",
			run: func() {
				NewUserController(suite.engine.Group("api"), suite.authChecker, mockservice.NewMockUserServiceWithContext(suite.mockUserService))
				id := uuid.New()
				token := jwt.New(jwt.SigningMethodES256)
				tokenString, _ := token.SignedString("")
				foundUser := model.User{
					Base: model.Base{
						ID: id,
					},
					Email: "abc",
				}
				request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/user/%s", id.String()), nil)
				header := http.Header{}
				header.Add("Authorization", fmt.Sprintf("Bearer %s", tokenString))
				request.Header = header
				suite.mockAuthService.On("Validate", mock.Anything, mock.Anything).Return(token, nil)
				suite.mockUserService.On("GetUserById", id).Return(foundUser, nil)

				suite.engine.ServeHTTP(suite.writer, request)
				suite.Equal(suite.writer.Code, http.StatusOK, "should return error if user does not exist")
				body, _ := io.ReadAll(suite.writer.Body)
				expectedResponse, _ := json.Marshal(response.NewGetUserResponse(foundUser))
				suite.Equal(body, expectedResponse, "should return user object")
			},
		},
	}
	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, test.run)
	}
}

func (suite *UserControllerTestSuite) TestGetUsersHandler() {
}

func (suite *UserControllerTestSuite) TestCreateUserHandler() {
}

func (suite *UserControllerTestSuite) TestUpdateUserHandler() {
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
