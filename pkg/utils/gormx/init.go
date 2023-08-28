package gormx

import (
	"apps/pkg/utils/configx"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init() (*gorm.DB, error) {

	dbConf := configx.InitConfigDb()

	db, err := gorm.Open(postgres.Open(dbConf.GetPostgresConnection()), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
