package service

import (
	"context"
	"posting-api/dto"
	"posting-api/entity"
	"posting-api/repository"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type PostService struct {
	db             *gorm.DB
	validator      *validator.Validate
	postRepository *repository.PostRepository
}

func NewPostService(db *gorm.DB, validator *validator.Validate, postRepository *repository.PostRepository) *PostService {
	return &PostService{
		postRepository: postRepository,
		db:             db,
		validator:      validator,
	}
}

func (p *PostService) HandleCreatePost(ctx context.Context, claims jwt.MapClaims, req *dto.CreatePostRequest) error {
	err := p.validator.Struct(req)
	if err != nil {
		return err
	}

	post := &entity.Post{
		Content:   req.Content,
		CreatedBy: claims["sub"].(string),
	}

	return p.postRepository.Create(ctx, p.db, post)
}

func (p *PostService) HandleUpdatePost(ctx context.Context, claims jwt.MapClaims, req *dto.UpdatePostRequest) error {
	err := p.validator.Struct(req)
	if err != nil {
		return err
	}

	post := &entity.Post{
		ID: req.ID,
	}

	err = p.postRepository.Take(ctx, p.db, post)
	if err != nil {
		return err
	}

	post.Content = req.Content

	return p.postRepository.Save(ctx, p.db, post)
}

func (p *PostService) HandleDeletePost(ctx context.Context, claims jwt.MapClaims, req *dto.DeletePostRequest) error {
	err := p.validator.Struct(req)
	if err != nil {
		return err
	}

	post := &entity.Post{
		ID:        req.ID,
		CreatedBy: claims["sub"].(string),
	}

	err = p.postRepository.Take(ctx, p.db, post)
	if err != nil {
		return err
	}

	return p.postRepository.Delete(ctx, p.db, post)
}

func (p *PostService) HandleGetAllPosts(ctx context.Context, req *dto.GetAllPostsRequest) ([]entity.Post, error) {
	var posts []entity.Post

	if req.Page == 0 {
		req.Page = 1
	}

	err := p.postRepository.PaginationPost(ctx, p.db, &posts, req.Page)

	return posts, err
}
