package converter

import (
	"posting-api/dto"
	"posting-api/entity"
)

func UserToResponseDto(entity *entity.User) *dto.UserDto {
	if entity != nil {
		return &dto.UserDto{
			ID:       entity.ID,
			Username: entity.Username,
			Email:    entity.Email,
			Posts:    PostsToReponseDto(entity.Posts),
		}
	}

	return nil
}

func UsersToResponseDto(entities []*entity.User) []*dto.UserDto {
	var response []*dto.UserDto

	for _, entity := range entities {
		response = append(response, UserToResponseDto(entity))
	}

	return response
}
