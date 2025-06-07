package service

import (
	"github.com/jumaevkova04/posts/internal/dto"
	"github.com/jumaevkova04/posts/internal/models"
	"github.com/jumaevkova04/posts/pkg/response"
)

func (s *Service) CreatePost(input *dto.CreatePostRequest) error {
	if err := input.Validate(); err != nil {
		return response.NewFailedValidationError(err)
	}

	newPost := models.NewPost(
		input.UserID,
		input.ImageURL,
		input.Text,
	)

	err := s.Repository.CreatePost(newPost)
	if err != nil {
		return response.NewServerError(err)
	}

	return nil
}

func (s *Service) GetPosts(input *dto.GetPostsRequest) (*dto.GetPostsResponse, error) {
	if err := input.Validate(); err != nil {
		return nil, response.NewFailedValidationError(err)
	}

	offset := (input.Page - 1) * input.PerPage

	posts, count, err := s.Repository.GetPosts(input.UserID, input.PerPage, offset)
	if err != nil {
		return nil, response.NewServerError(err)
	}

	postsResponse := &dto.GetPostsResponse{
		Posts: posts,
		Count: count,
	}

	return postsResponse, nil
}
