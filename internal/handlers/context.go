package handlers

import (
	"context"
	"net/http"
)

type contextKey string

const (
	authenticatedUserIDContextKey = contextKey("user_id")
)

func contextSetAuthenticatedUserID(r *http.Request, userID string) *http.Request {
	ctx := context.WithValue(r.Context(), authenticatedUserIDContextKey, userID)
	return r.WithContext(ctx)
}

func contextGetAuthenticatedUserID(r *http.Request) string {
	userID, ok := r.Context().Value(authenticatedUserIDContextKey).(string)
	if !ok {
		return ""
	}

	return userID
}
