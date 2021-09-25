package service

import (
	"context"
	"errors"
	"web-server/exceptions"
	"web-server/model"
	"web-server/repo"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserServiceWithContext func(ctx context.Context) UserService

type UserService interface {
	Create(model.User) (model.User, error)
	Update(model.User) error
	GetUserById(uuid.UUID) (model.User, error)
	ListUsers(page int, pageSize int) ([]model.User, int64, error)
	DeleteUserById(uuid.UUID) error
}

type userServiceImpl struct {
	ctx      context.Context
	userRepo repo.UserRepoWithContext
}

func (userService *userServiceImpl) Create(user model.User) (model.User, error) {
	user, err := userService.userRepo(userService.ctx).Create(user)
	if err != nil {
		return user, exceptions.UserCreateFaled.Wrap(err).SetMessage(err.Error())
	}
	return user, nil
}

func (userService *userServiceImpl) Update(user model.User) error {
	err := userService.userRepo(userService.ctx).Save(user)
	if err != nil {
		return exceptions.UserCreateFaled.Wrap(err)
	}
	return nil
}

func (userService *userServiceImpl) GetUserById(id uuid.UUID) (model.User, error) {
	user, err := userService.userRepo(userService.ctx).GetById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, exceptions.UserNotExists.Wrap(err)
	}
	return user, err
}

func (userService *userServiceImpl) ListUsers(page int, pageSize int) ([]model.User, int64, error) {
	query := repo.UserQuery{Offset: page * pageSize, Limit: pageSize}
	return userService.userRepo(userService.ctx).QueryUsers(query)
}

func (userService *userServiceImpl) DeleteUserById(id uuid.UUID) error {
	return userService.userRepo(userService.ctx).RemoveUserById(id)
}

func NewUserServiceWithContext(userRepo repo.UserRepoWithContext) UserServiceWithContext {
	return func(ctx context.Context) UserService {
		return &userServiceImpl{
			ctx:      ctx,
			userRepo: userRepo,
		}
	}
}
