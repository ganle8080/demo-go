package demo

import "gorm.io/gorm"

type Demo struct {
	gorm.Model
	Name string
	Age  int

	Ids []uint `gorm:"-"`
}
