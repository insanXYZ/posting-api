package dto

type CommentDto struct {
	ID        int      `json:"id,omitempty"`
	Comment   string   `json:"comment,omitempty"`
	CreatedBy *UserDto `json:"created_by,omitempty"`
}
