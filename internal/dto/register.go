package dto

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jumaevkova04/posts/internal/helpers"
)

type OTPToRegisterRequest struct {
	Email string `json:"email" validate:"required" example:"example@gmail.com"`
}

func (r *OTPToRegisterRequest) Validate() error {
	return v.Errors{
		"email": v.Validate(r.Email, v.Required, is.Email),
	}.Filter()
}

type CheckRegisterOTPRequest struct {
	Email string `json:"email" validate:"required" example:"example@gmail.com"`
	Otp   string `json:"otp" validate:"required" minLength:"6" maxLength:"6" example:"123456"`
}

func (r *CheckRegisterOTPRequest) Validate() error {
	return v.Errors{
		"email": v.Validate(r.Email, v.Required, is.Email),
		"otp":   v.Validate(r.Otp, v.Required, is.Int, v.Length(6, 6)),
	}.Filter()
}

type CheckRegisterOTPResponse struct {
	Nonce string `json:"nonce"`
}

type RegisterRequest struct {
	Nonce      string `json:"nonce" validate:"required" example:"exampleNonce"`
	Name       string `json:"name" validate:"required" minLength:"2" maxLength:"100" example:"John"`
	Surname    string `json:"surname" validate:"required" minLength:"2" maxLength:"100" example:"Doe"`
	Patronymic string `json:"patronymic" minLength:"2" maxLength:"100" example:"Adams"`
	Email      string `json:"email" validate:"required" example:"example@gmail.com"`
	Password   string `json:"password" validate:"required" minLength:"8" maxLength:"64" example:"p@ssW0rd"` // `min: one uppercase letter, one lowercase letter, one digit, one special character`
}

func (r *RegisterRequest) Validate() error {
	return v.Errors{
		"nonce":      v.Validate(r.Nonce, v.Required, is.Base64),
		"name":       v.Validate(r.Name, v.Required, v.Length(2, 100)),
		"surname":    v.Validate(r.Surname, v.Required, v.Length(2, 100)),
		"patronymic": v.Validate(r.Patronymic, v.Length(2, 100)),
		"email":      v.Validate(r.Email, v.Required, is.Email),
		"password":   v.Validate(r.Password, v.Required, v.By(helpers.StrongPassword)),
	}.Filter()
}
