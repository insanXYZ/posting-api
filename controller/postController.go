package controller

import (
	"posting-api/dto"
	"posting-api/dto/converter"
	"posting-api/service"
	"posting-api/util"

	"github.com/labstack/echo/v4"
)

type PostController struct {
	postService *service.PostService
}

func NewPostController(postService *service.PostService) *PostController {
	return &PostController{
		postService: postService,
	}
}

func (p *PostController) CreatePost(ctx echo.Context) error {
	claims := util.GetClaims(ctx)
	req := new(dto.CreatePostRequest)
	err := ctx.Bind(req)
	if err != nil {
		return err
	}

	err = p.postService.HandleCreatePost(ctx.Request().Context(), claims, req)
	if err != nil {
		return util.HttpResponseError(ctx, "failed create post", err)
	}

	return util.HttpResponseSuccess(ctx, "success create post", nil)

}

func (p *PostController) UpdatePost(ctx echo.Context) error {
	claims := util.GetClaims(ctx)

	req := new(dto.UpdatePostRequest)
	err := ctx.Bind(req)
	if err != nil {
		return err
	}

	err = p.postService.HandleUpdatePost(ctx.Request().Context(), claims, req)

	if err != nil {
		return util.HttpResponseError(ctx, "failed update post", err)
	}

	return util.HttpResponseSuccess(ctx, "success update post", nil)
}

func (p *PostController) DeletePost(ctx echo.Context) error {
	claims := util.GetClaims(ctx)

	req := new(dto.DeletePostRequest)
	err := ctx.Bind(req)
	if err != nil {
		return err
	}

	err = p.postService.HandleDeletePost(ctx.Request().Context(), claims, req)
	if err != nil {
		return util.HttpResponseError(ctx, "failed delete post", err)
	}

	return util.HttpResponseSuccess(ctx, "success delete post", nil)
}

func (p *PostController) GetAllPosts(ctx echo.Context) error {
	req := new(dto.GetAllPostsRequest)
	_ = ctx.Bind(req)

	posts, err := p.postService.HandleGetAllPosts(ctx.Request().Context(), req)
	if err != nil {
		return util.HttpResponseError(ctx, "failed get all posts", err)
	}

	return util.HttpResponseSuccess(ctx, "success get all posts", converter.PostsToReponseDto(posts))
}

func (p *PostController) GetPost(ctx echo.Context) error {
	req := new(dto.GetPostRequest)
	err := ctx.Bind(req)
	if err != nil {
		return err
	}

	post, err := p.postService.HandleGetPost(ctx.Request().Context(), req)
	if err != nil {
		return util.HttpResponseError(ctx, "failed get post", err)
	}

	return util.HttpResponseSuccess(ctx, "success get post", converter.PostToResponseDto(post))
}

func (p *PostController) LikePost(ctx echo.Context) error {
	claims := util.GetClaims(ctx)
	req := new(dto.LikePostRequest)
	err := ctx.Bind(req)
	if err != nil {
		return err
	}

	message := "post unliked successfully"

	liked, err := p.postService.HandleLikePost(ctx.Request().Context(), claims, req)
	if err != nil {
		return util.HttpResponseError(ctx, "failed like post", err)
	}

	if liked {
		message = "post liked successfully"
	}

	return util.HttpResponseSuccess(ctx, message, nil)
}

func (p *PostController) CommentPost(ctx echo.Context) error {
	claims := util.GetClaims(ctx)
	req := new(dto.CommentPostRequest)
	err := ctx.Bind(req)
	if err != nil {
		return err
	}

	err = p.postService.HandleCommentPost(ctx.Request().Context(), claims, req)
	if err != nil {
		return util.HttpResponseError(ctx, "failed send comment", err)
	}

	return util.HttpResponseSuccess(ctx, "success send comment", nil)
}
