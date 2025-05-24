package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id" gorm:"primaryKey;column:id"`
	Username  string         `json:"username" gorm:"column:username"`
	Email     string         `json:"email" gorm:"column:email"`
	Password  string         `json:"password" gorm:"column:password"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;autoCreatetime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Posts     []*Post        `json:"posts" gorm:"foreignKey:created_by;references:id"`
	LikePosts []*Post        `json:"like_posts" gorm:"many2many:user_like_post;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:post_id"`
}
