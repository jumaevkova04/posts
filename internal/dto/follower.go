package dto

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreateFollowerRequest struct {
	UserID     string `json:"-" swaggerignore:"true"`
	FollowerID string `json:"follower_id" validate:"required,uuid" example:"00000000-0000-0000-0000-000000000000"`
}

func (r *CreateFollowerRequest) Validate() error {
	return v.Errors{
		"follower_id": v.Validate(r.FollowerID, v.Required, is.UUID),
	}.Filter()
}

type GetFollowingsIDRequest struct {
	UserID string `json:"-" swaggerignore:"true"`
}
