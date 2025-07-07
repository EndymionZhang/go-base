package request

import (
	"gorm.io/gorm"
)

type PostsSearchRequest struct {
	PageRequest
	Title string `form:"keyword" json:"title"`
}

func (r *PostsSearchRequest) SearchCondition() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if keyword := r.Title; keyword != "" {
			db.Where("title like ?", "%"+keyword+"%")
		}
		return db
	}
}
