package dto

type PostDto struct {
	ID            string        `json:"id,omitempty"`
	Content       string        `json:"content,omitempty"`
	CreatedBy     *UserDto      `json:"created_by,omitempty"`
	LikeNumber    int           `json:"like_number"`
	CommentNumber int           `json:"comment_number"`
	Liked         []*UserDto    `json:"liked,omitempty"`
	Comments      []*CommentDto `json:"comments,omitempty"`
}

type CreatePostRequest struct {
	Content string `json:"content" validate:"required"`
}

type UpdatePostRequest struct {
	ID      string `param:"postId" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type DeletePostRequest struct {
	ID string `param:"postId" validate:"required"`
}

type GetAllPostsRequest struct {
	Page int `query:"page"`
}

type GetPostRequest struct {
	ID string `param:"postId"`
}

type LikePostRequest struct {
	ID string `param:"postId"`
}

type CommentPostRequest struct {
	ID      string `param:"postId" validate:"required"`
	Comment string `json:"comment" validate:"required,max=255"`
}
