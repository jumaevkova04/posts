package repository

import "github.com/jumaevkova04/posts/internal/models"

func (r *Repository) CreateFollower(follower *models.Follower) error {
	_, err := r.Exec(
		`INSERT INTO followers (id, following_id, follower_id, created_at)
		VALUES ($1, $2, $3, $4)`,
		follower.ID,
		follower.FollowingID,
		follower.FollowerID,
		follower.CreatedAt,
	)

	return err
}

func (r *Repository) FollowerExists(followerID, followingID string) (bool, error) {
	var exists bool

	query := `
		SELECT EXISTS (
			SELECT 1 FROM followers 
			WHERE follower_id = $1 AND following_id = $2
		)`

	err := r.Get(&exists, query, followerID, followingID)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *Repository) GetFollowingsID(userID string) ([]string, error) {
	var followingsID []string

	err := r.Select(&followingsID,
		`SELECT id 
		FROM followers WHERE following_id = $1`, userID)

	return followingsID, err
}
