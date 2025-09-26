package models

import "gorm.io/gorm"

type User struct{
	gorm.Model    // fields like ID, CreatedAt,UpdatedAt, DeletedAt
	Username string `gorm:"unique" json:"username"`
	Email string `gorm:"unique" json:"email"`
	Password string `json:"-"` 
}

type Post struct{
	gorm.Model
	Title string
	Content string
	UserID uint
	User    User `gorm:"foreignKey:UserID"`
}

type Tag struct{
	gorm.Model
	Name string `gorm:"unique"`
	Posts []*Post `gorm:"many2many:post_tags;"`

}