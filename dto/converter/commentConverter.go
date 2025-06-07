package converter

import (
	"posting-api/dto"
	"posting-api/entity"
)

func CommentToResponseDto(entity *entity.Comment) *dto.CommentDto {
	if entity != nil {
		return &dto.CommentDto{
			ID:        entity.ID,
			Comment:   entity.Comment,
			CreatedBy: UserToResponseDto(entity.User),
		}
	}

	return nil
}

func CommentsToResponseDto(entities []*entity.Comment) []*dto.CommentDto {
	var res []*dto.CommentDto

	for _, entity := range entities {
		res = append(res, CommentToResponseDto(entity))
	}

	return res
}
