package repository

import "posting-api/entity"

type CommentRepository struct {
	Repository[entity.Comment]
}

func NewCommentRepository() *CommentRepository {
	return &CommentRepository{}
}
