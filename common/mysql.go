package common

import (
	"demo-go/config"

	"gorm.io/gorm"
)

func LoadMysql(config *config.MysqlConfig) {

}

var instance *gorm.DB

func GetDB() *gorm.DB {
	return instance
}
