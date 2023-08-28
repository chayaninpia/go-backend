package xormx

import (
	"apps/pkg/utils/configx"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

func Init() (*xorm.Engine, error) {

	dbConf := configx.InitConfigDb()

	e, err := xorm.NewEngine("postgres", dbConf.GetPostgresConnection())
	if err != nil {
		return nil, err
	}

	e.Logger().SetLevel(log.LOG_INFO)
	e.ShowSQL(true)

	return e, err
}
