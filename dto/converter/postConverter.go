package converter

import (
	"posting-api/dto"
	"posting-api/entity"
)

func PostsToReponseDto(entities []*entity.Post) []*dto.PostDto {
	var responseDto []*dto.PostDto

	for _, entity := range entities {
		responseDto = append(responseDto, PostToResponseDto(entity))
	}

	return responseDto
}

func PostToResponseDto(entity *entity.Post) *dto.PostDto {
	if entity != nil {
		return &dto.PostDto{
			ID:         entity.ID,
			Content:    entity.Content,
			CreatedBy:  UserToResponseDto(entity.User),
			LikeNumber: len(entity.Liked),
			Liked:      UsersToResponseDto(entity.Liked),
			Comments:   CommentsToResponseDto(entity.Comments),
		}
	}
	return nil
}
