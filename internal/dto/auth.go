package dto

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jumaevkova04/posts/internal/helpers"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required" example:"example@gmail.com"`
	Password string `json:"password" validate:"required" minLength:"8" maxLength:"64" example:"p@ssW0rd"` // `min: one uppercase letter, one lowercase letter, one digit, one special character`
}

func (l *LoginRequest) Validate() error {
	return v.Errors{
		"email":    v.Validate(l.Email, v.Required, is.Email),
		"password": v.Validate(l.Password, v.Required, v.By(helpers.StrongPassword)),
	}.Filter()
}

type LoginResponse struct {
	Token       string `json:"token"`
	TokenExpiry string `json:"token_expiry"`
}
