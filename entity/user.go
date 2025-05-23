package entity

import (
	"time"

	"github.com/google/uuid"
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
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.NewString()
	return nil
}
