package dto

type CreatePostRequest struct {
	Content string `json:"content" validate:"required"`
}

type UpdatePostRequest struct {
	ID      string `param:"id" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type DeletePostRequest struct {
	ID string `param:"id" validate:"required"`
}

type GetAllPostsRequest struct {
	Page int `query:"page"`
}

type GetPostRequest struct {
	ID string `param:"id"`
}

type LikePostRequest struct {
	ID string `param:"id"`
}

type CommentPostRequest struct {
	ID      string `param:"id" validate:"required"`
	Comment string `json:"comment" validate:"required,max=255"`
}
