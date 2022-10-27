package orm

import "github.com/spf13/viper"

var (
	Host      = ``
	HostLocal = ``
	Port      = 0
	User      = ``
	Password  = ``
	Dbname    = ``
)

func InitConfigDb() {

	Host = viper.GetString(`db.host`)
	HostLocal = viper.GetString(`db.hostLocal`)
	Port = viper.GetInt(`db.port`)
	User = viper.GetString(`db.user`)
	Password = viper.GetString(`db.password`)
	Dbname = viper.GetString(`db.dbname`)

}
