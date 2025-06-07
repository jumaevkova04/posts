package dto

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jumaevkova04/posts/internal/models"
)

type CreatePostRequest struct {
	UserID   string `json:"-" swaggerignore:"true"`
	ImageURL string `json:"image_url" validate:"required,url" example:"https://example.com/image.jpg"`
	Text     string `json:"text" validate:"required" minLength:"10" maxLength:"300" example:"This is a post"`
}

func (r *CreatePostRequest) Validate() error {
	return v.Errors{
		"image_url": v.Validate(r.ImageURL, v.Required, is.URL),
		"text":      v.Validate(r.Text, v.Required, v.Length(10, 300)),
	}.Filter()
}

type GetPostsRequest struct {
	Page    int    `json:"page" validate:"required" min:"1" example:"1"`
	PerPage int    `json:"per_page" validate:"required" min:"1" max:"100" example:"10"`
	UserID  string `json:"-"`
}

func (r *GetPostsRequest) Validate() error {
	return v.Errors{
		"page":     v.Validate(r.Page, v.Required, v.Min(1)),
		"per_page": v.Validate(r.PerPage, v.Required, v.Min(1), v.Max(100)),
	}.Filter()
}

type GetPostsResponse struct {
	Posts []*models.Post `json:"posts"`
	Count int            `json:"count"`
}
