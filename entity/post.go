package entity

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        string         `gorm:"primaryKey;column:id"`
	Content   string         `gorm:"column:content"`
	CreatedBy string         `gorm:"column:created_by"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User      *User          `gorm:"foreignKey:created_by;references:id"`
	Liked     []*User        `gorm:"many2many:user_like_post;foreignKey:id;joinForeignKey:post_id;references:id;joinReferences:user_id"`
	Comments  []*Comment     `gorm:"foreignKey:post_id;references:id"`
}
