package repository

import (
	"github.com/jumaevkova04/posts/internal/models"
)

func (r *Repository) CreatePost(post *models.Post) error {
	_, err := r.Exec(
		`INSERT INTO posts (id, user_id, image_url, text, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		post.ID,
		post.UserID,
		post.ImageURL,
		post.Text,
		post.CreatedAt,
		post.UpdatedAt,
	)

	return err
}

//func (r *Repository) GetPosts(userID string) ([]*models.Post, error) {
//	var posts []*models.Post
//
//	query, args, err := sqlx.In(
//		`SELECT * FROM posts WHERE user_id IN
//        (SELECT following_id FROM followers WHERE follower_id = $1)`, userID)
//	if err != nil {
//		return nil, err
//	}
//
//	// Rebind query for the target driver (PostgreSQL = $1, $2, ...)
//	query = r.Rebind(query)
//
//	err = r.Select(&posts, query, args...)
//
//	return posts, err
//}

func (r *Repository) GetPosts(userID string, limit, offset int) ([]*models.Post, int, error) {
	var posts []*models.Post
	var count int

	countQuery := `
		SELECT COUNT(*) 
		FROM posts 
		WHERE user_id IN (
			SELECT following_id 
			FROM followers 
			WHERE follower_id = $1
		)`
	err := r.Get(&count, countQuery, userID)
	if err != nil {
		return nil, 0, err
	}

	query := `
		SELECT * 
		FROM posts 
		WHERE user_id IN (
			SELECT following_id 
			FROM followers 
			WHERE follower_id = $1
		)
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3`

	err = r.Select(&posts, query, userID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return posts, count, nil
}
