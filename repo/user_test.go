package repo

import (
	"context"
	"testing"
	mockDb "web-server/mock/db"
	"web-server/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type UserRepoTestSuite struct {
	suite.Suite
}

type unitTest struct {
	name string
	run  func()
}

func (suite *UserRepoTestSuite) TestCreateUser() {
	tests := []unitTest{
		unitTest{
			name: "should save successfully if db return no error",
			run: func() {
				ctx := context.Background()
				user := model.User{}

				db := mockDb.NewMockDatabase().(*mockDb.MockDatabase)
				db.On("Create", &user).Return(db)
				db.On("Error").Return(nil)

				userRepo := NewUserRepoWithContext(mockDb.NewMockDatabaseWithContext(db))
				returnedUser, err := userRepo(ctx).Create(user)
				suite.Equal(returnedUser, user, "should return user")
				suite.Nil(err, "should not return error")
			},
		},
		unitTest{
			name: "should return err db returns error",
			run: func() {
				ctx := context.Background()
				user := model.User{}

				db := mockDb.NewMockDatabase().(*mockDb.MockDatabase)
				db.On("Create", &user).Return(db)
				db.On("Error").Return(gorm.ErrInvalidData)

				userRepo := NewUserRepoWithContext(mockDb.NewMockDatabaseWithContext(db))
				_, err := userRepo(ctx).Create(user)
				suite.NotEmpty(err, "should return error")
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, test.run)
	}
}

func (suite *UserRepoTestSuite) TestGetById() {
	tests := []unitTest{
		unitTest{
			name: "should get successfully if db return no error",
			run: func() {
				ctx := context.Background()
				id := uuid.New()
				user := model.User{}

				db := mockDb.NewMockDatabase().(*mockDb.MockDatabase)
				db.On("First", &user, mock.Anything).Return(db)
				db.On("Error").Return(nil)

				userRepo := NewUserRepoWithContext(mockDb.NewMockDatabaseWithContext(db))
				returnedUser, err := userRepo(ctx).GetById(id)
				suite.Equal(returnedUser, user, "should return user")
				suite.Nil(err, "should not return error")
			},
		},
		unitTest{
			name: "should return err if db returns error",
			run: func() {
				ctx := context.Background()
				id := uuid.New()
				user := model.User{}

				db := mockDb.NewMockDatabase().(*mockDb.MockDatabase)
				db.On("First", &user, mock.Anything).Return(db)
				db.On("Error").Return(gorm.ErrInvalidData)

				userRepo := NewUserRepoWithContext(mockDb.NewMockDatabaseWithContext(db))
				_, err := userRepo(ctx).GetById(id)
				suite.NotEmpty(err, "should return error")
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, test.run)
	}
}

func (suite *UserRepoTestSuite) TestQueryUsers() {
	tests := []unitTest{
		unitTest{
			name: "should list users successfully if db returns no error",
			run: func() {
				ctx := context.Background()
				user := model.User{}
				query := UserQuery{
					Offset: 0,
					Limit:  25,
				}

				db := mockDb.NewMockDatabase().(*mockDb.MockDatabase)
				db.On("Model", mock.Anything).Return(db)
				db.On("Offset", mock.Anything).Return(db)
				db.On("Limit", mock.Anything).Return(db)
				db.On("Count", mock.Anything).Return(db)
				db.On("Find", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
					users := args[0].(*[]model.User)
					*users = []model.User{
						user,
					}
				}).Return(db)
				db.On("Error").Return(nil)

				userRepo := NewUserRepoWithContext(mockDb.NewMockDatabaseWithContext(db))
				users, _, err := userRepo(ctx).QueryUsers(query)
				suite.Equal(len(users), 1, "should return 1 user")
				suite.Equal(users[0], user, "should return user")
				suite.Nil(err, "should not return error")
			},
		},
		unitTest{
			name: "should return error if db returns no error",
			run: func() {
				ctx := context.Background()
				user := model.User{}
				query := UserQuery{
					Offset: 0,
					Limit:  25,
				}

				db := mockDb.NewMockDatabase().(*mockDb.MockDatabase)
				db.On("Model", mock.Anything).Return(db)
				db.On("Offset", mock.Anything).Return(db)
				db.On("Limit", mock.Anything).Return(db)
				db.On("Count", mock.Anything).Return(db)
				db.On("Find", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
					users := args[0].(*[]model.User)
					*users = []model.User{
						user,
					}
				}).Return(db)
				db.On("Error").Return(gorm.ErrRecordNotFound)

				userRepo := NewUserRepoWithContext(mockDb.NewMockDatabaseWithContext(db))
				_, _, err := userRepo(ctx).QueryUsers(query)
				suite.NotEmpty(err, "should return error")
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, test.run)
	}
}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}
