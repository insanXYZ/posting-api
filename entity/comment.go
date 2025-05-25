package entity

type Comment struct {
	ID      int    `gorm:"primaryKey;autoIncrement;column:id"`
	Comment string `gorm:"column:comment"`
	UserID  string `gorm:"column:user_id"`
	PostID  string `gorm:"column:post_id"`
	User    *User  `gorm:"foreignKey:user_id;references:id"`
	Post    *Post  `gorm:"foreignKey:post_id;references:id"`
}
