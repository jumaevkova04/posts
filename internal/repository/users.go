package repository

import (
	"database/sql"
	"errors"
	"github.com/jumaevkova04/posts/internal/helpers"
	"github.com/jumaevkova04/posts/internal/models"
)

func (r *Repository) CreateUser(user *models.User) error {
	_, err := r.Exec(
		`INSERT INTO users (id, name, surname, patronymic, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		user.ID,
		user.Name,
		user.Surname,
		user.Patronymic,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

func (r *Repository) UpdateUser(user *models.User) error {
	_, err := r.Exec(
		`UPDATE users SET name = $2, surname = $3, patronymic = $4, updated_at = $5
		WHERE id = $1`,
		user.ID,
		user.Name,
		user.Surname,
		user.Patronymic,
		helpers.TimeNowWithRFC3339Format(),
	)

	return err
}

func (r *Repository) GetUser(id string) (*models.User, error) {
	var user models.User

	err := r.Get(&user,
		`SELECT 
			id, name, surname, patronymic, email, created_at, updated_at 
		FROM users WHERE id = $1`, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &user, err
}

func (r *Repository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.Get(&user,
		`SELECT 
			id, name, surname, patronymic, email, password, created_at, updated_at 
		FROM users WHERE email = $1`, email)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &user, err
}
