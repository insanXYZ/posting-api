package entity

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        string         `json:"id" gorm:"primaryKey;column:id"`
	Content   string         `json:"content" gorm:"column:content"`
	CreatedBy string         `json:"created_by" gorm:"column:created_by"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	User      *User          `json:"user" gorm:"foreignKey:created_by;references:id"`
	Liked     []*User        `json:"liked" gorm:"many2many:user_like_post;foreignKey:id;joinForeignKey:post_id;references:id;joinReferences:user_id"`
}
