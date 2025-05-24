// CREATE TABLE comments (
//     id serial not null,
//     comment varchar(255) not null,
//     user_id varchar(100) not null,
//     post_id varchar(100) not null,
//     PRIMARY KEY(id),
//     CONSTRAINT fk_user
//         FOREIGN KEY(user_id) REFERENCES users(id),
//     CONSTRAINT fk_post
//         FOREIGN KEY(post_id) REFERENCES posts(id)
// )

package entity

type Comment struct {
	ID      int    `gorm:"primaryKey;autoIncrement;column:id"`
	Comment string `gorm:"column:comment"`
	UserID  string `gorm:"column:user_id"`
	PostID  string `gorm:"column:post_id"`
	User    *User  `gorm:"foreignKey:user_id;references:id"`
	Post    *Post  `gorm:"foreignKey:post_id;references:id"`
}
