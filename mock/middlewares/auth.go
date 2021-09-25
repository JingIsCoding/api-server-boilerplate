package middlewares

import (
	"web-server/restful/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockAuthCheck struct {
	mock.Mock
}

func (mock *MockAuthCheck) Check(ctx *gin.Context) {
	mock.Called(ctx)
}

func NewMockAuthChecker() middlewares.AuthChecker {
	return &MockAuthCheck{}
}
