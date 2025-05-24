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

func (a *App) GetEcho() *echo.Echo {
	return a.echo
}

func (a *App) GetDb() *gorm.DB {
	return a.db
}

func Init() *App {
	app := &App{
		echo:      config.NewEcho(),
		db:        config.NewGorm(),
		validator: config.NewValidator(),
	}

	// repository
	userRepository := repository.NewUserRepository()
	postRepository := repository.NewPostRepository()
	commentRepository := repository.NewCommentRepository()

	// service
	authService := service.NewAuthService(app.db, app.validator, userRepository)
	userService := service.NewUserService(app.db, app.validator, userRepository)
	postService := service.NewPostService(app.db, app.validator, userRepository, postRepository, commentRepository)

	// controller
	authController := controller.NewAuthController(authService)
	userController := controller.NewUserController(userService)
	postController := controller.NewPostController(postService)

	route.SetRoute(&route.RouteConfig{
		Echo:           app.echo,
		AuthController: authController,
		UserController: userController,
		PostController: postController,
	})

	return app
}

func (a *App) Run() error {
	return a.echo.Start(fmt.Sprintf(":%v", os.Getenv("APP_PORT")))
}
