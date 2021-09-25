package service

import (
	"context"
	"web-server/model"
	"web-server/service"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (service *MockUserService) Create(user model.User) (model.User, error) {
	args := service.Called(user)
	return args.Get(0).(model.User), args.Error(1)
}

func (service *MockUserService) Update(user model.User) error {
	args := service.Called(user)
	return args.Error(0)
}

func (service *MockUserService) GetUserById(id uuid.UUID) (model.User, error) {
	args := service.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}
func (service *MockUserService) ListUsers(page int, pageSize int) ([]model.User, int64, error) {
	args := service.Called(page, pageSize)
	return args.Get(0).([]model.User), args.Get(1).(int64), args.Error(2)
}
func (service *MockUserService) DeleteUserById(id uuid.UUID) error {
	args := service.Called(id)
	return args.Error(0)
}

func NewMockUserServiceWithContext(mockUserService *MockUserService) service.UserServiceWithContext {
	return func(ctx context.Context) service.UserService {
		return mockUserService
	}
}

//	Create(model.User) (model.User, error)
//	Update(model.User) error
//	GetUserById(uuid.UUID) (model.User, error)
//	ListUsers(page int, pageSize int) ([]model.User, model.Pagination, error)
//	DeleteUserById(uuid.UUID) error
