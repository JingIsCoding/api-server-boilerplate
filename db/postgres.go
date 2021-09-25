package db

import (
	"web-server/config"
	"web-server/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresqlDialector() gorm.Dialector {
	dbConf := config.Get().DatabaseConfig
	logger.Infof("trying to connect to %v", dbConf)
	return postgres.New(postgres.Config{
		DSN: dbConf.DatabaseUrl,
	})
}
