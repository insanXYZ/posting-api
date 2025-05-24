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
	SetMiddleware()

	// api
	api := cfg.Echo.Group("/api")
	api.POST("/register", cfg.AuthController.Register)
	api.POST("/login", cfg.AuthController.Login)
	api.POST("/refresh", cfg.AuthController.Refresh, RefreshJWT)

	// api/users
	users := api.Group("/users", HasJWT)
	users.GET("", cfg.UserController.GetUser)
	users.PUT("", cfg.UserController.UpdateUser)
	users.DELETE("", cfg.UserController.DeleteUser)

	// api/users/posts
	userPosts := users.Group("/posts")
	userPosts.POST("", cfg.PostController.CreatePost)
	userPosts.PUT("/:postId", cfg.PostController.UpdatePost)
	userPosts.DELETE("/:postId", cfg.PostController.DeletePost)

	// api/posts
	posts := api.Group("/posts")
	posts.GET("", cfg.PostController.GetAllPosts)
	posts.GET("/:postId", cfg.PostController.GetPost)

	postGuard := posts.Group("", HasJWT)
	postGuard.PUT("/:postId/like", cfg.PostController.LikePost)
	postGuard.POST("/:postId/comment", cfg.PostController.CommentPost)
}
