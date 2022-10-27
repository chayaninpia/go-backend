package gormx

import (
	"apps/src/db/orm"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
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

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
