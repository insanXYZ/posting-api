package service

import (
	"context"
	"posting-api/dto"
	"posting-api/entity"
	"posting-api/repository"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db             *gorm.DB
	validator      *validator.Validate
	userRepository *repository.UserRepository
}

func NewUserService(db *gorm.DB, validator *validator.Validate, userRepository *repository.UserRepository) *UserService {
	return &UserService{
		validator:      validator,
		db:             db,
		userRepository: userRepository,
	}
}

func (u *UserService) HandleGetUser(ctx context.Context, claims jwt.MapClaims) (*entity.User, error) {
	user := &entity.User{
		ID: claims["sub"].(string),
	}

	err := u.userRepository.Take(ctx, u.db, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) HandleUpdateUser(ctx context.Context, claims jwt.MapClaims, req *dto.UpdateUserRequest) error {
	err := u.validator.Struct(req)
	if err != nil {
		return err
	}

	user := &entity.User{
		ID: claims["sub"].(string),
	}

	err = u.userRepository.Take(ctx, u.db, user)
	if err != nil {
		return err
	}

	b, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Email = req.Email
	user.Password = string(b)
	user.Username = req.Username

	return u.userRepository.Update(ctx, u.db, user)
}

func (u *UserService) HandleDeleteUser(ctx context.Context, claims jwt.MapClaims) error {
	user := &entity.User{
		ID: claims["sub"].(string),
	}

	err := u.userRepository.Take(ctx, u.db, user)
	if err != nil {
		return err
	}

	return u.userRepository.Delete(ctx, u.db, user)
}
