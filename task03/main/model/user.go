package model

import "time"

type User struct {
	ID         int
	Name       string `gorm:"size:50"`
	Email      string `gorm:"size:50"`
	Posts      []Post `gorm:"foreignKey:AuthorID"`
	PostsCount int
	CreateAt   time.Time `gorm:"autoCreateTime"`
	UpdateAt   time.Time `gorm:"autoUpdateTime"`
}
