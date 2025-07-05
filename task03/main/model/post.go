package model

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID            int
	Title         string `gorm:"size:255"`
	Content       string `gorm:"type:text"`
	AuthorID      int
	Author        User
	CommentCount  int
	CommentStatus int `gorm:"type:tinyint"`
	Comments      []Comment
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdateAt      time.Time `gorm:"autoUpdateTime"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) error {
	tx.Model(&User{}).Debug().
		Where("id = ?", post.AuthorID).
		Updates(map[string]interface{}{
			"PostsCount": gorm.Expr("posts_count + 1"),
		})
	return nil
}
