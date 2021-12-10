package config

import(
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

const (
	Host = "127.0.0.1"
	Port = 3307
	DbName = "todos"
	UserName = "root"
	Password = "123"
)

type DbConfig struct {
	Host string
	port int32
	DbName string
	UserName string
	Password string
}

func BuildDbConfig() *DbConfig {
	config := DbConfig{
		Host,
		Port,
		DbName,
		UserName,
		Password,
	}
	return &config
}

func GetDbUrl(config *DbConfig) string {

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.UserName,config.Password,config.Host,config.port,config.DbName,
		)
}