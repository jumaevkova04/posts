package models

import "github.com/jumaevkova04/posts/internal/helpers"

type Post struct {
	ID        string `json:"id" db:"id"`
	UserID    string `json:"user_id" db:"user_id"`
	ImageURL  string `json:"image_url" db:"image_url"`
	Text      string `json:"text" db:"text"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

func NewPost(
	userID string,
	imageURL string,
	text string,
) *Post {
	return &Post{
		ID:        helpers.NewUUID(),
		UserID:    userID,
		ImageURL:  imageURL,
		Text:      text,
		CreatedAt: helpers.TimeNowWithRFC3339Format(),
		UpdatedAt: helpers.TimeNowWithRFC3339Format(),
	}
}
