package models

import "gorm.io/gorm"

type User struct{
	gorm.Model    // fields like ID, CreatedAt,UpdatedAt, DeletedAt
	Username string `gorm:"unique" json:"username"`
	Email string `gorm:"unique" json:"email"`
	Password string `json:"password"` 
}

