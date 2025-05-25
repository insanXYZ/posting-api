package dto

type UserDto struct {
	ID        string     `json:"id,omitempty"`
	Username  string     `json:"username,omitempty"`
	Email     string     `json:"email,omitempty"`
	Posts     []*PostDto `json:"posts,omitempty"`
	LikePosts []*PostDto `json:"like_posts,omitempty"`
}

type UpdateUserRequest struct {
	Username string `json:"username" validate:"omitempty,max=50"`
	Email    string `json:"email" validate:"omitempty,email,max=50"`
	Password string `json:"password" validate:"omitempty,min=6,max=100"`
}
