package service

import (
	"github.com/jumaevkova04/posts/internal/dto"
	"github.com/jumaevkova04/posts/internal/models"
	"github.com/jumaevkova04/posts/pkg/response"
)

func (s *Service) CreateFollower(input *dto.CreateFollowerRequest) error {
	if err := input.Validate(); err != nil {
		return response.NewFailedValidationError(err)
	}

	if input.UserID == input.FollowerID {
		return response.NewBadRequestError(response.ErrYouCantFollowYourself)
	}

	exists, err := s.Repository.FollowerExists(input.FollowerID, input.UserID)
	if err != nil {
		return response.NewServerError(err)
	}

	if exists {
		return response.NewBadRequestError(response.ErrFollowerExists)
	}

	newFollower := models.NewFollower(
		input.FollowerID,
		input.UserID,
	)

	err = s.Repository.CreateFollower(newFollower)
	if err != nil {
		return response.NewServerError(err)
	}

	return nil
}

func (s *Service) GetFollowingsID(input *dto.GetFollowingsIDRequest) ([]string, error) {
	followingsID, err := s.Repository.GetFollowingsID(input.UserID)
	if err != nil {
		return nil, response.NewServerError(err)
	}

	return followingsID, nil
}
