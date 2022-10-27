package xormx

import (
	"apps/src/db/orm"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

func Init() (*xorm.Engine, error) {

	psqlInfo := ``
	hostName, _ := os.Hostname()
	if strings.Contains(hostName, `local`) {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			orm.HostLocal, orm.Port, orm.User, orm.Password, orm.Dbname)
	} else {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			orm.Host, orm.Port, orm.User, orm.Password, orm.Dbname)
	}

	e, err := xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	e.Logger().SetLevel(log.LOG_INFO)
	e.ShowSQL(true)

	return e, err
}
