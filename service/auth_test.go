package service

import (
	"context"
	"errors"
	"testing"
	"time"
	"web-server/config"
	"web-server/exceptions"
	mockRepo "web-server/mock/repo"
	"web-server/model"
	"web-server/repo"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/suite"
)

type AuthServiceTestSuite struct {
	suite.Suite
	userRepo               *mockRepo.MockUserRepo
	userRepoWithContext    repo.UserRepoWithContext
	authServiceWithContext AuthServiceWithContext
}

func (suite *AuthServiceTestSuite) SetupTest() {
	suite.userRepo = &mockRepo.MockUserRepo{}
	suite.userRepoWithContext = mockRepo.NewMockUserRepoWithContext(suite.userRepo)
	suite.authServiceWithContext = NewAuthServiceWithContext(suite.userRepoWithContext)
}

func (suite *AuthServiceTestSuite) TestAuthenticateByEmail() {
	tests := []unitTest{
		unitTest{
			name: "should return token if user is found",
			run: func() {
				user := model.User{
					EncryptedPassword: model.Password("123").Encrypt(),
				}
				suite.userRepo.On("GetByEmail", "abc@gmail.com").Return(user, nil)
				token, err := suite.authServiceWithContext(context.Background()).AuthenticateByEmail("abc@gmail.com", "123")
				suite.Nil(err, "should not have any error")
				suite.NotNil(token, "should return token")
			},
		},
		unitTest{
			name: "should return error if user is not found",
			run: func() {
				user := model.User{
					EncryptedPassword: model.Password("123").Encrypt(),
				}
				suite.userRepo.On("GetByEmail", "abc@gmail.com").Return(user, exceptions.UserNotExists)
				_, err := suite.authServiceWithContext(context.Background()).AuthenticateByEmail("abc@gmail.com", "123")
				suite.NotNil(err, "should return error")
			},
		},
		unitTest{
			name: "should return error if password not matched",
			run: func() {
				user := model.User{
					EncryptedPassword: model.Password("123").Encrypt(),
				}
				suite.userRepo.On("GetByEmail", "abc@gmail.com").Return(user, exceptions.UserNotExists)
				_, err := suite.authServiceWithContext(context.Background()).AuthenticateByEmail("abc@gmail.com", "456")
				suite.NotNil(err, "should return error")
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, test.run)
	}
}

func (suite *AuthServiceTestSuite) TestValidate() {
	tests := []unitTest{
		unitTest{
			name: "should not return error if validate successeed",
			run: func() {
				conf := config.Get()
				signedToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"exp": time.Now().Add(time.Duration(30 * time.Second)).Unix(),
				}).SignedString([]byte(conf.ServerConfig.HMACSecret))
				token, err := suite.authServiceWithContext(context.Background()).Validate(signedToken, jwt.MapClaims{})
				suite.Nil(err, "should not have error")
				suite.NotNil(token, "should return a parsed token")
			},
		},
		unitTest{
			name: "should not return error if token expires",
			run: func() {
				conf := config.Get()
				signedToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"exp": time.Now().Add(time.Duration(-90 * time.Minute)).Unix(),
				}).SignedString([]byte(conf.ServerConfig.HMACSecret))
				_, err := suite.authServiceWithContext(context.Background()).Validate(signedToken, jwt.MapClaims{})
				suite.NotNil(err, "should have error")
				suite.Equalf(errors.Is(err, exceptions.TokenExpires), true, "should return token expires")
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, test.run)
	}
}

func (suite *AuthServiceTestSuite) TestRefreshToken() {}

func TestAuthServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}
