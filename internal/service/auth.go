package service

import (
	"github.com/jumaevkova04/posts/internal/dto"
	"github.com/jumaevkova04/posts/internal/helpers"
	"github.com/jumaevkova04/posts/pkg/jwt"
	"github.com/jumaevkova04/posts/pkg/password"
	"github.com/jumaevkova04/posts/pkg/response"
	"time"
)

func (s *Service) Login(input *dto.LoginRequest) (*dto.LoginResponse, error) {
	if err := input.Validate(); err != nil {
		return nil, response.NewFailedValidationError(err)
	}

	user, err := s.Repository.GetUserByEmail(input.Email)
	if err != nil {
		return nil, response.NewServerError(err)
	}

	if user == nil {
		return nil, response.NewBadRequestError(response.ErrInvalidCredentials)
	}

	passwordMatches, err := password.Matches(input.Password, user.Password)
	if err != nil {
		return nil, response.NewServerError(err)
	}

	if !passwordMatches {
		return nil, response.NewBadRequestError(response.ErrInvalidCredentials)
	}

	expiry := 24 * time.Hour

	token, tokenExpiry, err := jwt.GenerateJwtToken(s.Config.BaseURL, s.Config.Jwt.SecretKey, user.ID, expiry)
	if err != nil {
		return nil, response.NewServerError(err)
	}

	data := &dto.LoginResponse{
		Token:       token,
		TokenExpiry: tokenExpiry.Format(helpers.TimeRFC3339Format),
	}

	return data, nil
}
