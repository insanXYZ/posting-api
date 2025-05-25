package controller

import (
	"posting-api/dto"
	"posting-api/dto/converter"
	"posting-api/dto/message"
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
		return util.HttpResponseError(ctx, message.FAILED_CREATE_POST, err)
	}

	return util.HttpResponseSuccess(ctx, message.SUCCESS_CREATE_POST, nil)

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
		return util.HttpResponseError(ctx, message.FAILED_UPDATE_POST, err)
	}

	return util.HttpResponseSuccess(ctx, message.SUCCESS_UPDATE_POST, nil)
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
		return util.HttpResponseError(ctx, message.FAILED_DELETE_POST, err)
	}

	return util.HttpResponseSuccess(ctx, message.SUCCESS_DELETE_POST, nil)
}

func (p *PostController) GetAllPosts(ctx echo.Context) error {
	req := new(dto.GetAllPostsRequest)
	_ = ctx.Bind(req)

	posts, err := p.postService.HandleGetAllPosts(ctx.Request().Context(), req)
	if err != nil {
		return util.HttpResponseError(ctx, message.FAILED_GET_ALL_POSTS, err)
	}

	return util.HttpResponseSuccess(ctx, message.SUCCESS_GET_ALL_POSTS, converter.PostsToReponseDto(posts))
}

func (p *PostController) GetPost(ctx echo.Context) error {
	req := new(dto.GetPostRequest)
	err := ctx.Bind(req)
	if err != nil {
		return err
	}

	post, err := p.postService.HandleGetPost(ctx.Request().Context(), req)
	if err != nil {
		return util.HttpResponseError(ctx, message.FAILED_GET_POST, err)
	}

	return util.HttpResponseSuccess(ctx, message.SUCCESS_GET_POST, converter.PostToResponseDto(post))
}

func (p *PostController) LikePost(ctx echo.Context) error {
	claims := util.GetClaims(ctx)
	req := new(dto.LikePostRequest)
	err := ctx.Bind(req)
	if err != nil {
		return err
	}

	msg := message.SUCCESS_UNLIKE_POST

	liked, err := p.postService.HandleLikePost(ctx.Request().Context(), claims, req)
	if err != nil {
		return util.HttpResponseError(ctx, message.FAILED_LIKED_POST, err)
	}

	if liked {
		msg = message.SUCCESS_LIKE_POST
	}

	return util.HttpResponseSuccess(ctx, msg, nil)
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
		return util.HttpResponseError(ctx, message.FAILED_COMMENT_POST, err)
	}

	return util.HttpResponseSuccess(ctx, message.SUCCESS_COMMENT_POST, nil)
}
