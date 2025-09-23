package models

import "gorm.io/gorm"

type User struct{
	gorm.Model    // fields like ID, CreatedAt,UpdatedAt, DeletedAt
	Username string `gorm:"unique"`
	Email string `gorm:"unique"`
	Password string
}