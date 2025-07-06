package model

import "gorm.io/gorm"

type Comment struct {
	BaseEntity
	Content string `gorm:"type:text" json:"content"`
	PostID  int
	Post    Post `gorm:"foreignKey:PostID" json:"post"`
	UserId  int
	User    User `gorm:"foreignKey:UserId" json:"user"`
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
