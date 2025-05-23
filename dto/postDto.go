package dto

type CreatePostRequest struct {
	Content string `json:"content" validate:"required,max=255"`
}

type UpdatePostRequest struct {
	ID      string `param:"id" validate:"required"`
	Content string `json:"content" validate:"required,max=255"`
}

type DeletePostRequest struct {
	ID string `param:"id" validate:"required"`
}

type GetAllPostsRequest struct {
	Page int `query:"page"`
}
