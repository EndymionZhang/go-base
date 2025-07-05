package question

import (
	"fmt"
	"github.com/endymion/go-base/task03/main/model"
	"gorm.io/gorm"
)

type User = model.User
type Post = model.Post
type Comment = model.Comment

func RunBlog(db *gorm.DB) {
	initBlobData(db)
	//findUserData(db)
	//findMaxComments(db)
	//deleteComment(db)
}

func initBlobData(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	err = db.AutoMigrate(&Post{})
	err = db.AutoMigrate(&Comment{})
	if err != nil {
		return
	}
	users := &[]User{
		{Name: "张三", Email: "zhangsan@163.com"},
		{Name: "李四", Email: "lisi@163.com"},
		{Name: "王五", Email: "wangwu@163.com"}}
	db.Create(users)

	post := &Post{Title: "Go 语言", Content: "Go 语言是一个开源的编程语言 -- post", Author: User{ID: 1}}
	db.Create(post)
	comment := &[]Comment{
		{Content: "Go 语言是一个开源的编程语言 -- comment", User: User{ID: 2}, Post: *post},
		{Content: "Go 语言是一个开源的编程语言 -- comment", User: User{ID: 3}, Post: *post},
	}
	db.Create(comment)
}

func findUserData(db *gorm.DB) {
	var user User
	user.ID = 1
	// 查询 user 时，带出 所有关联信息
	tx := db.Debug().Preload("Posts").Preload("Posts.Comments").Where(user).First(&user)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	} else {
		fmt.Println(user)
	}
}

func findMaxComments(db *gorm.DB) {
	var post Post
	// 查询评论最多的一篇文章
	tx := db.Debug().Order("comment_count desc").First(&post)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	} else {
		fmt.Println(post)
	}
}

func deleteComment(db *gorm.DB) {
	var comment Comment
	comment.UserId = 2
	comment.PostID = 2
	tx := db.Debug().Where(comment).Delete(&comment)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	} else {
		fmt.Println(comment)
	}
}
