package model

import (
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

func CreateNewUser(username, password string) *User {
	return &User{Username: username, Password: password}
}

func UserMigrate(db *gorm.DB) {
	user := &User{}
	if err := db.AutoMigrate(user); err != nil {
		log.Fatal(err)
	}
}
