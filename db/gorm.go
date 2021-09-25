package db

import (
	"context"
	"sync"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"web-server/config"
	"web-server/logger"
	// Import support for postgres
)

var once sync.Once
var db *gorm.DB

type GormDatabase struct {
	ctx context.Context
	db  *gorm.DB
}

func NewDatabase() DatabaseWithCtx {
	once.Do(func() {
		var err error
		retries := 3
		// change to other dialector if needed
		dialector := NewPostgresqlDialector()
		gormConfig := &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			PrepareStmt: true,
		}
		db, err = gorm.Open(dialector, gormConfig)
		sql, _ := db.DB()
		err = sql.Ping()
		for err != nil {
			if retries == 0 {
				logger.Infof("Failed to connect to database %v", err)
			}
			logger.Infof("Failed to connect to %s, retry later", config.Get().DatabaseConfig.DatabaseUrl)
			time.Sleep(5 * time.Second)
			err = sql.Ping()
			retries--
		}
	})
	return func(ctx context.Context) Database {
		return &GormDatabase{
			db:  db,
			ctx: ctx,
		}
	}
}

func (g *GormDatabase) WithTimeout(function func(Database) Database, timeout time.Duration) Database {
	timeoutCtx, cancelFunc := context.WithTimeout(g.ctx, timeout)
	defer cancelFunc()
	return function(newGormDatabase(timeoutCtx, g.db))
}

// Begin return a new instance of Database
func (g *GormDatabase) Begin() Database {
	return newGormDatabase(g.ctx, g.db.Begin())
}

func (g *GormDatabase) Transaction(transaction func(Database) error) error {
	err := g.db.Transaction(func(tx *gorm.DB) error {
		err := transaction(newGormDatabase(g.ctx, tx))
		return err
	})
	return err
}

func (g *GormDatabase) Commit() Database {
	return newGormDatabase(g.ctx, g.db.Commit())
}

func (g *GormDatabase) Rollback() Database {
	return newGormDatabase(g.ctx, g.db.Rollback())
}

func (g *GormDatabase) Debug() Database {
	return newGormDatabase(g.ctx, g.db.Debug())
}

func (g *GormDatabase) Scopes(funcs ...func(Database) Database) Database {
	var db Database
	for _, f := range funcs {
		db = f(g).(*GormDatabase)
	}
	return db
}

func (g *GormDatabase) Unscoped() Database {
	return newGormDatabase(g.ctx, g.db.Unscoped())
}

func (g *GormDatabase) Save(value interface{}) Database {
	return newGormDatabase(g.ctx, g.db.Save(value))
}

func (g *GormDatabase) Model(value interface{}) Database {
	return newGormDatabase(g.ctx, g.db.Model(value))
}

func (g *GormDatabase) Table(name string) Database {
	return newGormDatabase(g.ctx, g.db.Table(name))
}

func (g *GormDatabase) Update(column string, value interface{}) Database {
	return newGormDatabase(g.ctx, g.db.Update(column, value))
}

func (g *GormDatabase) Updates(values interface{}) Database {
	return newGormDatabase(g.ctx, g.db.Updates(values))
}

func (g *GormDatabase) Create(value interface{}) Database {
	return newGormDatabase(g.ctx, g.db.Create(value))
}

func (g *GormDatabase) Select(query interface{}, args ...interface{}) Database {
	return newGormDatabase(g.ctx, g.db.Select(query, args...))
}

func (g *GormDatabase) Where(query interface{}, args ...interface{}) Database {
	return newGormDatabase(g.ctx, g.db.Where(query, args...))
}

func (g *GormDatabase) Or(query interface{}, args ...interface{}) Database {
	return newGormDatabase(g.ctx, g.db.Or(query, args...))
}

func (g *GormDatabase) Not(query interface{}, args ...interface{}) Database {
	return newGormDatabase(g.ctx, g.db.Not(query, args...))
}

func (g *GormDatabase) Raw(sql string, values ...interface{}) Database {
	return newGormDatabase(g.ctx, g.db.Raw(sql, values...))
}

func (g *GormDatabase) Exec(sql string, values ...interface{}) Database {
	return newGormDatabase(g.ctx, g.db.Exec(sql, values...))
}

func (g *GormDatabase) Joins(query string, args ...interface{}) Database {
	return newGormDatabase(g.ctx, g.db.Joins(query, args...))
}

func (g *GormDatabase) Delete(value interface{}, where ...interface{}) Database {
	return newGormDatabase(g.ctx, g.db.Delete(value, where...))
}

func (g *GormDatabase) Last(out interface{}, where ...interface{}) Database {
	return newGormDatabase(g.ctx, g.db.Last(out, where...))
}

func (g *GormDatabase) Find(out interface{}, where ...interface{}) Database {
	return newGormDatabase(g.ctx, g.db.Find(out, where...))
}

func (g *GormDatabase) Group(fields string) Database {
	return newGormDatabase(g.ctx, g.db.Group(fields))
}

func (g *GormDatabase) Preload(column string, conditions ...interface{}) Database {
	var newConditions []interface{}
	for _, condition := range conditions {
		conditionFun, ok := condition.(func(db Database) Database)
		if ok {
			condition = func(db *gorm.DB) *gorm.DB {
				database := conditionFun(newGormDatabase(g.ctx, db))
				gorm := database.(*GormDatabase)
				return gorm.GormDB()
			}
		}
		newConditions = append(newConditions, condition)
	}
	return newGormDatabase(g.ctx, g.db.Preload(column, newConditions...))
}

func (g *GormDatabase) Scan(dest interface{}) Database {
	return newGormDatabase(g.ctx, g.db.Scan(dest))
}

func (g *GormDatabase) Take(out interface{}, where ...interface{}) Database {
	return newGormDatabase(g.ctx, g.db.Take(out, where...))
}

func (g *GormDatabase) First(out interface{}, where ...interface{}) Database {
	return newGormDatabase(g.ctx, g.db.First(out, where...))
}

func (g *GormDatabase) Order(value interface{}) Database {
	return newGormDatabase(g.ctx, g.db.Order(value))
}

func (g *GormDatabase) Limit(limit int) Database {
	return newGormDatabase(g.ctx, g.db.Limit(limit))
}

func (g *GormDatabase) Offset(offset int) Database {
	return newGormDatabase(g.ctx, g.db.Offset(offset))
}

func (g *GormDatabase) Count(count *int64) Database {
	return newGormDatabase(g.ctx, g.db.Count(count))
}

func (g *GormDatabase) Distinct() Database {
	return newGormDatabase(g.ctx, g.db.Distinct())
}
func (g *GormDatabase) Pluck(column string, value interface{}) Database {
	return newGormDatabase(g.ctx, g.db.Pluck(column, value))
}

//DB return gorm db
func (g *GormDatabase) GormDB() *gorm.DB {
	return g.db
}

// Error returns the error
func (g *GormDatabase) Error() error {
	return g.db.Error
}

func newGormDatabase(ctx context.Context, db *gorm.DB) Database {
	return &GormDatabase{
		db:  db.WithContext(ctx),
		ctx: ctx,
	}
}
