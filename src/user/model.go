package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex;not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"uniqueIndex;"`
	Phone    string `json:"phone" gorm:"uniqueIndex;not null"`
	Gender   string
}

func (m *User) TableName() string {
	return "user"
}
