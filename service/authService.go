package service

import (
	"context"
	"errors"
	"posting-api/dto"
	"posting-api/entity"
	"posting-api/repository"
	"posting-api/util"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db             *gorm.DB
	validator      *validator.Validate
	userRepository *repository.UserRepository
}

func NewAuthService(db *gorm.DB, validator *validator.Validate, userRepository *repository.UserRepository) *AuthService {
	return &AuthService{
		db:             db,
		validator:      validator,
		userRepository: userRepository,
	}
}

func (a *AuthService) HandleRegister(ctx context.Context, req *dto.RegisterRequest) error {
	err := a.validator.Struct(req)
	if err != nil {
		return err
	}

	user := &entity.User{
		Email: req.Email,
	}

	err = a.userRepository.Take(ctx, a.db, user)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("email already in use")
	}

	bc, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := &entity.User{
		ID:       uuid.NewString(),
		Username: req.Username,
		Email:    req.Email,
		Password: string(bc),
	}

	return a.userRepository.Create(ctx, a.db, newUser)
}

func (a *AuthService) HandleLogin(ctx context.Context, req *dto.LoginRequest) (string, error) {
	err := a.validator.Struct(req)
	if err != nil {
		return "", err
	}

	user := &entity.User{
		Email: req.Email,
	}

	err = a.userRepository.Take(ctx, a.db, user)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.New("email or password wrong")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("email or password wrong")
	}

	return util.CreateJWT(jwt.MapClaims{
		"sub":  user.ID,
		"name": user.Username,
		"exp":  time.Now().Add(2 * time.Minute).Unix(),
	})
}

func (a *AuthService) HandleRefresh(ctx context.Context, claims jwt.MapClaims) (string, error) {
	user := &entity.User{
		ID: claims["sub"].(string),
	}

	err := a.userRepository.Take(ctx, a.db, user)
	if err != nil {
		return "", err
	}

	return util.CreateJWT(jwt.MapClaims{
		"sub":  user.ID,
		"name": user.Username,
		"exp":  time.Now().Add(15 * time.Minute).Unix(),
	})
}
