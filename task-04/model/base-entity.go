package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseEntity struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoCreateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"delete_at,omitempty"`
}

func (b *BaseEntity) BeforeDelete(tx *gorm.DB) error {
	now := time.Now()
	tx.Statement.SetColumn("DeleteAt", now)
	return nil
}
