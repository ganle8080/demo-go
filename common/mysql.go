package common

import "gorm.io/gorm"

var db *gorm.DB

func LoadMysql() {

}

func Get() *gorm.DB {
	return db
}
