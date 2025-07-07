package request

import "gorm.io/gorm"

type PageRequest struct {
	Page     int `form:"page" json:"page" validate:"min:1"`
	PageSize int `form:"page_size" json:"page_size" validate:"min:1"`
}

func (r *PageRequest) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if r.Page <= 0 {
			r.Page = 1
		}
		switch {
		case r.PageSize > 100:
			r.PageSize = 100
		case r.PageSize <= 0:
			r.PageSize = 10
		}
		offset := (r.Page - 1) * r.PageSize
		return db.Offset(offset).Limit(r.PageSize)
	}
}

type RequestType int

const (
	Get RequestType = iota
	Post
	Put
	Delete
)

type IdRequest struct {
	Id uint `json:"id" validate:"gt=0"`
}
