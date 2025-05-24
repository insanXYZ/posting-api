package service

import (
	"context"
	"posting-api/dto"
	"posting-api/entity"
	"posting-api/repository"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostService struct {
	db                *gorm.DB
	validator         *validator.Validate
	userRepository    *repository.UserRepository
	commentRepository *repository.CommentRepository
	postRepository    *repository.PostRepository
}

func NewPostService(db *gorm.DB, validator *validator.Validate, userRepository *repository.UserRepository, postRepository *repository.PostRepository, commentRepository *repository.CommentRepository) *PostService {
	return &PostService{
		postRepository:    postRepository,
		db:                db,
		userRepository:    userRepository,
		validator:         validator,
		commentRepository: commentRepository,
	}
}

func (p *PostService) HandleCreatePost(ctx context.Context, claims jwt.MapClaims, req *dto.CreatePostRequest) error {
	err := p.validator.Struct(req)
	if err != nil {
		return err
	}

	post := &entity.Post{
		ID:        uuid.NewString(),
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

	if req.Page <= 1 {
		req.Page = 0
	}

	err := p.postRepository.TakeDetailPostsWithPagination(ctx, p.db, &posts, req.Page)

	return posts, err
}

func (p *PostService) HandleGetPost(ctx context.Context, req *dto.GetPostRequest) (*entity.Post, error) {
	post := &entity.Post{
		ID: req.ID,
	}

	err := p.postRepository.TakeDetailPost(ctx, p.db, post)

	return post, err
}

func (p *PostService) HandleLikePost(ctx context.Context, claims jwt.MapClaims, req *dto.LikePostRequest) (bool, error) {
	var liked bool

	err := p.db.Transaction(func(tx *gorm.DB) error {
		user := &entity.User{
			ID: claims["sub"].(string),
		}

		post := &entity.Post{
			ID: req.ID,
		}

		err := p.userRepository.Take(ctx, tx, user)
		if err != nil {
			return err
		}

		err = p.postRepository.Take(ctx, tx, post)
		if err != nil {
			return err
		}

		if count := p.postRepository.CountUserLikePost(ctx, tx, post, user); count == 0 {
			liked = true
			return p.postRepository.Liked(ctx, tx, post, user)
		}
		liked = false
		return p.postRepository.Unliked(ctx, tx, post, user)
	})
	return liked, err
}

func (p *PostService) HandleCommentPost(ctx context.Context, claims jwt.MapClaims, req *dto.CommentPostRequest) error {
	err := p.validator.Struct(req)
	if err != nil {
		return err
	}

	err = p.db.Transaction(func(tx *gorm.DB) error {
		user := &entity.User{
			ID: claims["sub"].(string),
		}

		post := &entity.Post{
			ID: req.ID,
		}

		err = p.userRepository.Take(ctx, tx, user)
		if err != nil {
			return err
		}

		err = p.postRepository.Take(ctx, tx, post)
		if err != nil {
			return err
		}

		return p.commentRepository.Create(ctx, tx, &entity.Comment{
			Comment: req.Comment,
			UserID:  user.ID,
			PostID:  post.ID,
		})
	})
	return err
}
