package db

import (
	"context"
	"time"
)

type DatabaseWithCtx func(context.Context) Database

// Database defines the interface for Relational database
// This is the same interface as what gorm provided
type Database interface {
	Debug() Database
	Scopes(funcs ...func(Database) Database) Database
	Unscoped() Database
	Save(value interface{}) Database
	Take(out interface{}, where ...interface{}) Database
	First(out interface{}, where ...interface{}) Database
	Find(out interface{}, where ...interface{}) Database
	Preload(column string, conditions ...interface{}) Database
	Scan(dest interface{}) Database
	Last(out interface{}, where ...interface{}) Database
	Model(value interface{}) Database
	Table(name string) Database
	Update(column string, value interface{}) Database
	Updates(values interface{}) Database
	Create(value interface{}) Database
	Select(query interface{}, args ...interface{}) Database
	Where(query interface{}, args ...interface{}) Database
	Or(query interface{}, args ...interface{}) Database
	Not(query interface{}, args ...interface{}) Database
	Raw(sql string, values ...interface{}) Database
	Exec(sql string, values ...interface{}) Database
	Delete(value interface{}, where ...interface{}) Database
	Joins(query string, args ...interface{}) Database
	Order(value interface{}) Database
	Group(fields string) Database
	Limit(limit int) Database
	Offset(offset int) Database
	Count(count *int64) Database
	Commit() Database
	Rollback() Database
	Error() error
	Begin() Database
	Transaction(transaction func(Database) error) error
	WithTimeout(function func(Database) Database, timeout time.Duration) Database
	Distinct() Database
	Pluck(column string, value interface{}) Database
}
