package service

import (
	"context"
	"web-server/service"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (mock *MockAuthService) AuthenticateByEmail(email, password string) (token string, err error) {
	results := mock.Called(email, password)
	return results.Get(0).(string), results.Error(1)
}

func (mock *MockAuthService) Validate(email string, claims jwt.MapClaims) (token *jwt.Token, err error) {
	results := mock.Called(email, claims)
	return results.Get(0).(*jwt.Token), results.Error(1)
}

func (mock *MockAuthService) RefreshToken(token string) (newToken string, err error) {
	results := mock.Called(token)
	return results.Get(0).(string), results.Error(1)
}

func NewMockAuthServiceWithContext(mockAuthService service.AuthService) service.AuthServiceWithContext {
	return func(context.Context) service.AuthService {
		return mockAuthService
	}
}
