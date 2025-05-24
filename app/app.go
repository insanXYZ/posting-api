package app

import (
	"fmt"
	"os"
	"posting-api/config"
	"posting-api/controller"
	"posting-api/repository"
	"posting-api/route"
	"posting-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type App struct {
	echo      *echo.Echo
	db        *gorm.DB
	validator *validator.Validate
}

func Init() *App {
	return &App{
		echo:      config.NewEcho(),
		db:        config.NewGorm(),
		validator: config.NewValidator(),
	}
}

func (a *App) Run() error {
	// repository
	userRepository := repository.NewUserRepository()
	postRepository := repository.NewPostRepository()
	commentRepository := repository.NewCommentRepository()

	// service
	authService := service.NewAuthService(a.db, a.validator, userRepository)
	userService := service.NewUserService(a.db, a.validator, userRepository)
	postService := service.NewPostService(a.db, a.validator, userRepository, postRepository, commentRepository)

	// controller
	authController := controller.NewAuthController(authService)
	userController := controller.NewUserController(userService)
	postController := controller.NewPostController(postService)

	route.SetRoute(&route.RouteConfig{
		Echo:           a.echo,
		AuthController: authController,
		UserController: userController,
		PostController: postController,
	})

	return a.echo.Start(fmt.Sprintf(":%v", os.Getenv("APP_PORT")))
}
