package model

import (
	"fmt"
	"github.com/endymion/go-base/task-04/common/model/request"
	"gorm.io/gorm"
)

type Post struct {
	BaseEntity
	Title         string `gorm:"size:255" json:"title" validate:"min:1,max:255"`
	Content       string `gorm:"type:text" json:"content" validate:"min:1"`
	AuthorID      uint
	Author        User      `gorm:"foreignKey:AuthorID" json:"author" oig:"omitempty"`
	CommentCount  int       `gorm:"type:int;default:0" json:"commentCount"`
	CommentStatus int       `gorm:"type:tinyint"`
	Comments      []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty" oig:"omitempty"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) error {
	tx.Model(&User{}).Debug().
		Where("id = ?", post.AuthorID).
		Updates(map[string]interface{}{
			"PostsCount": gorm.Expr("posts_count + 1"),
		})
	return nil
}

func (post *Post) CreatePost() error {
	return db.Debug().Create(post).Error
}

func ListPosts(r *request.PostsSearchRequest) (posts []Post) {
	paginate := r.Paginate()
	db := paginate(db)
	condition := r.SearchCondition()
	db = condition(db)
	tx := db.Model(&Post{}).Find(&posts)
	fmt.Println(tx.Error)
	return posts
}

func CountPosts(r *request.PostsSearchRequest) (total int64) {
	paginate := r.Paginate()
	db := paginate(db)
	condition := r.SearchCondition()
	db = condition(db)
	tx := db.Debug().Model(&Post{}).Count(&total)
	fmt.Println(tx.Error)
	return
}

func DeletePost(id uint, userId uint) (bool, error) {
	tx := db.Where("id = ? AND author_id = ?", id, userId).Delete(&Post{})
	err := tx.Error
	affected := tx.RowsAffected
	return affected == 1, err
}

func GetPostDetail(id uint) Post {
	var post Post
	db.Where("id = ?", id).First(&post)
	return post
}

func EditPost(req *Post, id uint) (bool, error) {
	tx := db.Model(&Post{}).Where("id = ? and author_id = ?", req.ID, id).Updates(req)
	err := tx.Error
	affected := tx.RowsAffected
	return affected == 1, err
}
