package models

import (
	"github.com/jumaevkova04/posts/internal/helpers"
	"gopkg.in/guregu/null.v3"
)

type User struct {
	ID         string      `json:"id" db:"id"`
	Name       string      `json:"name" db:"name"`
	Surname    string      `json:"surname" db:"surname"`
	Patronymic null.String `json:"patronymic" db:"patronymic"`
	Email      string      `json:"email" db:"email"`
	Password   string      `json:"password,omitempty" db:"password"`
	CreatedAt  string      `json:"created_at" db:"created_at"`
	UpdatedAt  string      `json:"updated_at" db:"updated_at"`
}

func NewUser(
	name string,
	surname string,
	patronymic string,
	email string,
	password string,
) *User {
	return &User{
		ID:         helpers.NewUUID(),
		Name:       name,
		Surname:    surname,
		Patronymic: helpers.NullString(patronymic),
		Email:      email,
		Password:   password,
		CreatedAt:  helpers.TimeNowWithRFC3339Format(),
		UpdatedAt:  helpers.TimeNowWithRFC3339Format(),
	}
}
