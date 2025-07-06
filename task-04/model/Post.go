package model

import "gorm.io/gorm"

type Post struct {
	BaseEntity
	Title         string `gorm:"size:255"`
	Content       string `gorm:"type:text"`
	AuthorID      int
	Author        User      `gorm:"foreignKey:AuthorID" json:"author"`
	CommentCount  int       `gorm:"type:int;default:0" json:"commentCount"`
	CommentStatus int       `gorm:"type:tinyint"`
	Comments      []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) error {
	tx.Model(&User{}).Debug().
		Where("id = ?", post.AuthorID).
		Updates(map[string]interface{}{
			"PostsCount": gorm.Expr("posts_count + 1"),
		})
	return nil
}
