package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `gorm:"primaryKey;column:id"`
	Username  string         `gorm:"column:username"`
	Email     string         `gorm:"column:email"`
	Password  string         `gorm:"column:password"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreatetime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Posts     []*Post        `gorm:"foreignKey:created_by;references:id"`
	LikePosts []*Post        `gorm:"many2many:user_like_post;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:post_id"`
}
