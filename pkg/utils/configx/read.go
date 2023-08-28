package configx

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func GetPort() string {

	hostName, _ := os.Hostname()
	if strings.Contains(hostName, `local`) {
		return fmt.Sprintf(":%v", viper.GetString(`portLocal`))
	}
	return fmt.Sprintf(":%v", viper.GetString(`port`))
}

func GetMongoURI() string {
	uri, _ := base64.RawStdEncoding.DecodeString(viper.GetString("mongodb.uri"))
	return string(uri)
}

func GetMongoDbName() string {
	return viper.GetString("mongodb.dbname")
}

func GetJWTSignature() string {
	uri, _ := base64.RawStdEncoding.DecodeString(viper.GetString("jwt.signature"))
	return string(uri)
}

type PostgresConf struct {
	Host      string
	HostLocal string
	Port      string
	User      string
	Password  string
	Dbname    string
}

func InitConfigDb() *PostgresConf {
	host, _ := base64.RawStdEncoding.DecodeString(viper.GetString("db.host"))
	hostLocal, _ := base64.RawStdEncoding.DecodeString(viper.GetString("db.hostLocal"))
	port, _ := base64.RawStdEncoding.DecodeString(viper.GetString("db.port"))
	user, _ := base64.RawStdEncoding.DecodeString(viper.GetString("db.user"))
	password, _ := base64.RawStdEncoding.DecodeString(viper.GetString("db.password"))
	return &PostgresConf{
		Host:      string(host),
		HostLocal: string(hostLocal),
		Port:      string(port),
		User:      string(user),
		Password:  string(password),
		Dbname:    viper.GetString(`db.dbname`),
	}
}

func (conf PostgresConf) GetPostgresConnection() string {
	hostName, _ := os.Hostname()
	if strings.Contains(hostName, `local`) {
		return fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			conf.HostLocal, conf.Port, conf.User, conf.Password, conf.Dbname)
	}
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Dbname)

}
