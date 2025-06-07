package service

import (
	"github.com/jumaevkova04/posts/internal/dto"
	"github.com/jumaevkova04/posts/internal/helpers"
	"github.com/jumaevkova04/posts/internal/models"
	"github.com/jumaevkova04/posts/pkg/response"
)

func (s *Service) UserInfo(input *dto.UserInfoRequest) (*models.User, error) {
	user, err := s.Repository.GetUser(input.ID)
	if err != nil {
		return nil, response.NewServerError(err)
	}

	return user, nil
}

func (s *Service) UpdateUserInfo(input *dto.UpdateUserRequest) error {
	if err := input.Validate(); err != nil {
		return response.NewFailedValidationError(err)
	}

	user, err := s.Repository.GetUser(input.ID)
	if err != nil {
		return response.NewServerError(err)
	}

	user.Name = input.Name
	user.Surname = input.Surname
	user.Patronymic = helpers.NullString(input.Patronymic)

	err = s.Repository.UpdateUser(user)
	if err != nil {
		return response.NewServerError(err)
	}

	return nil
}
