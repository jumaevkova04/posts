package dto

import v "github.com/go-ozzo/ozzo-validation/v4"

type UserInfoRequest struct {
	ID string `json:"-" swaggerignore:"true"`
}

type UpdateUserRequest struct {
	ID         string `json:"-" swaggerignore:"true"`
	Name       string `json:"name" validate:"required" minLength:"2" maxLength:"100" example:"John"`
	Surname    string `json:"surname" validate:"required" minLength:"2" maxLength:"100" example:"Doe"`
	Patronymic string `json:"patronymic" minLength:"2" maxLength:"100" example:"Adams"`
}

func (r *UpdateUserRequest) Validate() error {
	return v.Errors{
		"name":       v.Validate(r.Name, v.Required, v.Length(2, 100)),
		"surname":    v.Validate(r.Surname, v.Required, v.Length(2, 100)),
		"patronymic": v.Validate(r.Patronymic, v.Length(2, 100)),
	}.Filter()
}
