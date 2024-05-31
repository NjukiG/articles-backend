package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Articles []Article `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:UserID"`
}

type Article struct {
	gorm.Model
	Title         string `gorm:"unique"`
	SubTitle      string
	Image         string
	Body          string
	MinutesToRead int64
	UserID        uint
	Comments      []Comment `gorm:"foreignKey:ArticleID"`
}

type Comment struct {
	gorm.Model
	Content   string `gorm:"unique"`
	UserID    uint
	ArticleID uint
}
