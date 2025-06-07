package handlers

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/jumaevkova04/posts/docs"
	"github.com/jumaevkova04/posts/pkg/cors"
	"github.com/jumaevkova04/posts/pkg/response"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"
)

// Routes godoc
// @title POSTS API
// @version 1.0
// host localhost:3939
// @BasePath /api
// @Schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func (h *Handler) Routes() http.Handler {
	router := chi.NewRouter()

	router.NotFound(response.NotFound)
	router.MethodNotAllowed(response.MethodNotAllowed)

	router.Use(cors.Cors)
	router.Use(h.recoverPanic)

	router.Get("/api/status", response.Status)
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	router.Group(func(r chi.Router) {
		r.Use(h.authenticate)

		r.Mount("/api", r)

		r.Route("/register", func(r chi.Router) {
			r.Post("/otp", h.handle(h.OTPToRegister))
			r.Post("/check_otp", h.handle(h.CheckRegisterOTP))
			r.Post("/", h.handle(h.Register))
		})

		r.Post("/login", h.handle(h.Login))

		r.Group(func(r chi.Router) {
			r.Use(h.authentication)

			r.Route("/users", func(r chi.Router) {
				r.Get("/me", h.handle(h.UserInfo))
				r.Put("/update", h.handle(h.UpdateUser))
			})

			r.Route("/posts", func(r chi.Router) {
				r.Post("/", h.handle(h.CreatePost))
				r.Get("/", h.handle(h.GetPosts))
			})

			r.Post("/follow", h.handle(h.CreateFollower))

			r.Get("/followings", h.handle(h.GetFollowingsID))
		})
	})

	return router
}
