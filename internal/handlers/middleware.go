package handlers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jumaevkova04/posts/pkg/response"
	"net/http"
	"strings"
)

func (h *Handler) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var res response.Response

		defer func() {
			if err := recover(); err != nil {
				res.Error = response.NewServerError(fmt.Errorf("%s", err))
				res.Send(w, r)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var res response.Response

		w.Header().Add("Vary", "Authorization")

		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader != "" {
			headerParts := strings.Split(authorizationHeader, " ")

			if len(headerParts) == 2 && headerParts[0] == "Bearer" {
				headerToken := headerParts[1]

				token, err := jwt.Parse(headerToken, func(token *jwt.Token) (interface{}, error) {
					return []byte(h.Config.Jwt.SecretKey), nil
				})
				if err != nil {
					res.Error = response.NewInvalidAuthenticationTokenError(response.ErrInvalidAuthenticationToken)
					res.Send(w, r)
					return
				}

				claims := token.Claims

				claimsIssuer, err := claims.GetIssuer()
				if err != nil {
					res.Error = response.NewInvalidAuthenticationTokenError(response.ErrInvalidAuthenticationToken)
					res.Send(w, r)
					return
				}

				if claimsIssuer != h.Config.BaseURL {
					res.Error = response.NewInvalidAuthenticationTokenError(response.ErrInvalidAuthenticationToken)
					res.Send(w, r)
					return
				}

				userID, err := claims.GetSubject()
				if err != nil {
					res.Error = response.NewInvalidAuthenticationTokenError(response.ErrInvalidAuthenticationToken)
					res.Send(w, r)
					return
				}

				user, err := h.Service.Repository.GetUser(userID)
				if err != nil {
					res.Error = response.NewServerError(err)
					res.Send(w, r)
					return
				}

				if user != nil {
					r = contextSetAuthenticatedUserID(r, user.ID)
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var res response.Response

		authenticatedUserID := contextGetAuthenticatedUserID(r)

		if authenticatedUserID == "" {
			res.Error = response.NewAuthenticationRequiredError(response.ErrAuthenticationRequired)
			res.Send(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
