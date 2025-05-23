package route

import (
	"posting-api/controller"

	"github.com/labstack/echo/v4"
)

type RouteConfig struct {
	Echo           *echo.Echo
	AuthController *controller.AuthController
	UserController *controller.UserController
	PostController *controller.PostController
}

func SetRoute(cfg *RouteConfig) {
	// api
	api := cfg.Echo.Group("/api")
	api.POST("/register", cfg.AuthController.Register)
	api.POST("/login", cfg.AuthController.Login)

	// api/users
	users := api.Group("/users", HasJWT)
	users.GET("", cfg.UserController.GetUser)
	users.PATCH("", cfg.UserController.UpdateUser)

	// api/users/posts
	userPosts := users.Group("/posts")
	userPosts.POST("", cfg.PostController.CreatePost)
	userPosts.PUT("/:id", cfg.PostController.UpdatePost)
	userPosts.DELETE("/:id", cfg.PostController.DeletePost)

	// api/posts
	posts := api.Group("/posts")
	posts.GET("", cfg.PostController.GetAllPosts)
}
