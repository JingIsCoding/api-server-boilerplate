package repo

import (
	"context"
	"web-server/model"
	"web-server/repo"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (repo *MockUserRepo) Create(user model.User) (model.User, error) {
	args := repo.Called(user)
	return args.Get(0).(model.User), args.Error(1)
}

func (repo *MockUserRepo) Save(user model.User) error {
	args := repo.Called(user)
	return args.Error(0)
}

func (repo *MockUserRepo) GetById(id uuid.UUID) (model.User, error) {
	args := repo.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

func (repo *MockUserRepo) GetByEmail(email string) (model.User, error) {
	args := repo.Called(email)
	return args.Get(0).(model.User), args.Error(1)
}

func (repo *MockUserRepo) QueryUsers(query repo.UserQuery) ([]model.User, int64, error) {
	args := repo.Called(query)
	return args.Get(0).([]model.User), args.Get(1).(int64), args.Error(2)
}

func (repo *MockUserRepo) RemoveUserById(id uuid.UUID) error {
	args := repo.Called(id)
	return args.Error(0)
}

func NewMockUserRepoWithContext(userRepo repo.UserRepo) repo.UserRepoWithContext {
	return func(context.Context) repo.UserRepo {
		return userRepo
	}
}

//type UserRepo interface {
//	Create(model.User) (model.User, error)
//	Save(model.User) error
//	GetById(uuid.UUID) (model.User, error)
//	QueryUsers(query UserQuery) ([]model.User, model.Pagination, error)
//	RemoveUserById(uuid.UUID) error
//}
