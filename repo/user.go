package repo

import (
	"context"
	"web-server/db"
	"web-server/model"

	"github.com/google/uuid"
)

type UserRepoWithContext func(ctx context.Context) UserRepo

type UserQuery struct {
	FistName *string
	LastName *string
	Email    *string
	Phone    *string
	Offset   int
	Limit    int
}

type UserRepo interface {
	Create(model.User) (model.User, error)
	Save(model.User) error
	GetById(uuid.UUID) (model.User, error)
	GetByEmail(string) (model.User, error)
	QueryUsers(query UserQuery) (users []model.User, total int64, err error)
	RemoveUserById(uuid.UUID) error
}

type userRepoImpl struct {
	ctx context.Context
	db  db.Database
}

func NewUserRepoWithContext(db db.DatabaseWithCtx) UserRepoWithContext {
	return func(ctx context.Context) UserRepo {
		return &userRepoImpl{
			ctx: ctx,
			db:  db(ctx),
		}
	}
}

func (repo *userRepoImpl) Create(user model.User) (model.User, error) {
	err := repo.db.Create(&user).Error()
	return user, err
}

func (repo *userRepoImpl) Save(user model.User) error {
	return repo.db.Save(&user).Error()
}

func (repo *userRepoImpl) GetById(id uuid.UUID) (model.User, error) {
	var user model.User
	err := repo.db.First(&user, "id = ?", id).Error()
	return user, err
}

func (repo *userRepoImpl) GetByEmail(email string) (model.User, error) {
	var user model.User
	err := repo.db.First(&user, "email = ?", email).Error()
	return user, err
}

func (repo *userRepoImpl) QueryUsers(query UserQuery) ([]model.User, int64, error) {
	var count int64
	users := []model.User{}
	db := repo.db.Model(&model.User{}).
		Offset(query.Offset).
		Limit(query.Limit).
		Count(&count)
	err := db.Find(&users).Error()
	return users, count, err
}

func (repo *userRepoImpl) RemoveUserById(id uuid.UUID) error {
	return repo.db.Model(&model.User{}).Delete(&model.User{}, id).Error()
}
