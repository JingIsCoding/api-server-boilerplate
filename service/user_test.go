package service

import (
	"context"
	"testing"
	"web-server/exceptions"
	"web-server/mock/repo"
	"web-server/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
}

type unitTest struct {
	name string
	run  func()
}

func (suite *UserServiceTestSuite) TestCreateUser() {
	tests := []unitTest{
		unitTest{
			name: "should not return error if create successfully",
			run: func() {
				ctx := context.Background()
				user := model.User{}
				mockUserRepo := new(repo.MockUserRepo)
				mockUserService := NewUserServiceWithContext(repo.NewMockUserRepoWithContext(mockUserRepo))

				mockUserRepo.On("Create", user).Return(user, nil)

				returnedUser, err := mockUserService(ctx).Create(user)
				suite.Equal(returnedUser, user)
				suite.Nil(err, "should not return error")
			},
		},
		unitTest{
			name: "should return error if create failed",
			run: func() {
				ctx := context.Background()
				user := model.User{}
				mockUserRepo := new(repo.MockUserRepo)
				mockUserService := NewUserServiceWithContext(repo.NewMockUserRepoWithContext(mockUserRepo))

				mockUserRepo.On("Create", user).Return(user, exceptions.UserAlreadyExist)

				returnedUser, err := mockUserService(ctx).Create(user)
				suite.Equal(returnedUser, user)
				suite.NotNil(err, "should return error")
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, test.run)
	}
}
func (suite *UserServiceTestSuite) TestUpdateUser() {
	tests := []unitTest{
		unitTest{
			name: "should not return error if update successfully",
			run: func() {
				ctx := context.Background()
				user := model.User{}
				mockUserRepo := new(repo.MockUserRepo)
				mockUserService := NewUserServiceWithContext(repo.NewMockUserRepoWithContext(mockUserRepo))

				mockUserRepo.On("Save", user).Return(nil)

				err := mockUserService(ctx).Update(user)
				suite.Nil(err, "should not return error")
			},
		},
		unitTest{
			name: "should return error if repo save returns error",
			run: func() {
				ctx := context.Background()
				user := model.User{}
				mockUserRepo := new(repo.MockUserRepo)
				mockUserService := NewUserServiceWithContext(repo.NewMockUserRepoWithContext(mockUserRepo))

				mockUserRepo.On("Save", user).Return(exceptions.UserNotExists)

				err := mockUserService(ctx).Update(user)
				suite.NotNil(err, "should return error")
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, test.run)
	}
}
func (suite *UserServiceTestSuite) TestGetUserById() {
	tests := []unitTest{
		unitTest{
			name: "should get user successfully",
			run: func() {
				ctx := context.Background()
				user := model.User{}
				id := uuid.New()

				mockUserRepo := new(repo.MockUserRepo)
				mockUserService := NewUserServiceWithContext(repo.NewMockUserRepoWithContext(mockUserRepo))

				mockUserRepo.On("GetById", id).Return(user, nil)

				returnedUser, err := mockUserService(ctx).GetUserById(id)
				suite.Equal(returnedUser, user)
				suite.Nil(err, "should not return error")
			},
		},
		unitTest{
			name: "should return user not exists repo returns error",
			run: func() {
				ctx := context.Background()
				user := model.User{}
				id := uuid.New()

				mockUserRepo := new(repo.MockUserRepo)
				mockUserService := NewUserServiceWithContext(repo.NewMockUserRepoWithContext(mockUserRepo))

				mockUserRepo.On("GetById", id).Return(user, exceptions.UserNotExists)

				returnedUser, err := mockUserService(ctx).GetUserById(id)
				suite.Equal(returnedUser, user)
				suite.NotNil(err, "should return error")
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, test.run)
	}
}
func (suite *UserServiceTestSuite) TestListUsers() {
	tests := []unitTest{
		unitTest{
			name: "should list user successfully",
			run: func() {
				ctx := context.Background()
				users := []model.User{
					model.User{},
				}

				mockUserRepo := new(repo.MockUserRepo)
				mockUserService := NewUserServiceWithContext(repo.NewMockUserRepoWithContext(mockUserRepo))

				mockUserRepo.On("QueryUsers", mock.Anything).Return(users, int64(1), nil)

				returnedUsers, count, err := mockUserService(ctx).ListUsers(0, 25)
				suite.Equal(returnedUsers, users)
				suite.Equal(count, int64(1))
				suite.Nil(err, "should return error")
			},
		},
		unitTest{
			name: "should return error if repo returns error",
			run: func() {
				ctx := context.Background()
				users := []model.User{
					model.User{},
				}

				mockUserRepo := new(repo.MockUserRepo)
				mockUserService := NewUserServiceWithContext(repo.NewMockUserRepoWithContext(mockUserRepo))

				mockUserRepo.On("QueryUsers", mock.Anything).Return(users, int64(1), exceptions.UserNotExists)

				returnedUsers, returnedPage, err := mockUserService(ctx).ListUsers(0, 25)
				suite.Equal(returnedUsers, users)
				suite.Equal(returnedPage, int64(1))
				suite.NotNil(err, "should return error")
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, test.run)
	}
}

func (suite *UserServiceTestSuite) TestDeleteUserById() {
	tests := []unitTest{
		unitTest{
			name: "should delete user successfully",
			run: func() {
				ctx := context.Background()
				id := uuid.New()

				mockUserRepo := new(repo.MockUserRepo)
				mockUserService := NewUserServiceWithContext(repo.NewMockUserRepoWithContext(mockUserRepo))

				mockUserRepo.On("RemoveUserById", id).Return(nil)

				err := mockUserService(ctx).DeleteUserById(id)
				suite.Nil(err, "should not return error")
			},
		},
		unitTest{
			name: "should fail to delete if repo returns error",
			run: func() {
				ctx := context.Background()
				id := uuid.New()

				mockUserRepo := new(repo.MockUserRepo)
				mockUserService := NewUserServiceWithContext(repo.NewMockUserRepoWithContext(mockUserRepo))

				mockUserRepo.On("RemoveUserById", id).Return(exceptions.UserNotExists)

				err := mockUserService(ctx).DeleteUserById(id)
				suite.NotNil(err, "should return error")
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, test.run)
	}
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
