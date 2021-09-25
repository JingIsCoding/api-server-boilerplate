package db

import (
	"context"
	"time"
	"web-server/db"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDatabase struct {
	mock.Mock
}

func (mock *MockDatabase) Distinct() db.Database {
	args := mock.Called()
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Pluck(column string, value interface{}) db.Database {
	args := mock.Called(column, value)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) GormDB() *gorm.DB {
	args := mock.Called()
	return args.Get(0).(*gorm.DB)
}

func (mock *MockDatabase) WithTimeout(function func(database db.Database) db.Database, timeout time.Duration) db.Database {
	args := mock.Called(function, timeout)
	return args.Get(0).(db.Database)
}

// Debug return a new instance of Database
func (mock *MockDatabase) Debug() db.Database {
	args := mock.Called()
	return args.Get(0).(db.Database)
}

// Begin return a new instance of Database
func (mock *MockDatabase) Begin() db.Database {
	args := mock.Called()
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Commit() db.Database {
	args := mock.Called()
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Rollback() db.Database {
	args := mock.Called()
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Scopes(funcs ...func(db.Database) db.Database) db.Database {
	args := mock.Called(funcs)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Unscoped() db.Database {
	args := mock.Called()
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Save(value interface{}) db.Database {
	args := mock.Called(value)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Model(value interface{}) db.Database {
	args := mock.Called(value)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Table(name string) db.Database {
	args := mock.Called(name)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Update(column string, value interface{}) db.Database {
	args := mock.Called(column, value)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Updates(values interface{}) db.Database {
	args := mock.Called(values)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Create(value interface{}) db.Database {
	args := mock.Called(value)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Select(query interface{}, arguments ...interface{}) db.Database {
	args := mock.Called(query, arguments)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Where(query interface{}, arguments ...interface{}) db.Database {
	args := mock.Called(query, arguments)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Or(query interface{}, arguments ...interface{}) db.Database {
	args := mock.Called(query, arguments)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Not(query interface{}, arguments ...interface{}) db.Database {
	args := mock.Called(query, arguments)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Raw(sql string, values ...interface{}) db.Database {
	args := mock.Called(sql, values)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Exec(sql string, values ...interface{}) db.Database {
	args := mock.Called(sql, values)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Limit(limit int) db.Database {
	args := mock.Called(limit)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Offset(offset int) db.Database {
	args := mock.Called(offset)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Group(fields string) db.Database {
	args := mock.Called(fields)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Joins(query string, arguments ...interface{}) db.Database {
	args := mock.Called(query, arguments)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Delete(value interface{}, where ...interface{}) db.Database {
	args := mock.Called(value, where)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Last(out interface{}, where ...interface{}) db.Database {
	args := mock.Called(out, where)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Find(out interface{}, where ...interface{}) db.Database {
	args := mock.Called(out, where)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Preload(column string, conditions ...interface{}) db.Database {
	args := mock.Called(column, conditions)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Preloads(out interface{}) db.Database {
	args := mock.Called(out)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Scan(dest interface{}) db.Database {
	args := mock.Called(dest)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Take(out interface{}, where ...interface{}) db.Database {
	args := mock.Called(out, where)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) First(out interface{}, where ...interface{}) db.Database {
	args := mock.Called(out, where)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Order(value interface{}) db.Database {
	args := mock.Called(value)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Count(count *int64) db.Database {
	args := mock.Called(count)
	return args.Get(0).(db.Database)
}

func (mock *MockDatabase) Transaction(transaction func(db.Database) error) error {
	args := mock.Called(transaction)
	return args.Error(0)
}

// Error returns the error
func (mock *MockDatabase) Error() error {
	args := mock.Called()
	return args.Error(0)
}

func NewMockDatabase() db.Database {
	return new(MockDatabase)
}

func NewMockDatabaseWithContext(mockDb db.Database) db.DatabaseWithCtx {
	return func(ctx context.Context) db.Database {
		return mockDb
	}
}
