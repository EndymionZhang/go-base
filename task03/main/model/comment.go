package model

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID       int
	Content  string `gorm:"type:text"`
	PostID   int
	Post     Post
	UserId   int
	User     User
	CreateAt time.Time `gorm:"autoCreateTime"`
	UpdateAt time.Time `gorm:"autoUpdateTime"`
}

func (comment *Comment) AfterCreate(tx *gorm.DB) error {
	tx.Model(&Post{}).Debug().
		Where("id = ?", comment.PostID).
		Updates(map[string]interface{}{
			"CommentCount":  gorm.Expr("comment_count + 1"),
			"CommentStatus": 1,
		})
	return nil
}

func (comment *Comment) BeforeDelete(tx *gorm.DB) error {
	tx.Model(&Post{}).Debug().
		Where("id = ?", comment.PostID).
		Updates(map[string]interface{}{
			"CommentCount":   gorm.Expr("comment_count - 1"),
			"comment_status": gorm.Expr("CASE WHEN comment_count - 1 <= 0 THEN 0 ELSE comment_status END"),
		})
	return nil
}
