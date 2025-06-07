package handlers

import (
	"errors"
	"github.com/jumaevkova04/posts/internal/config"
	"github.com/jumaevkova04/posts/internal/service"
	"github.com/jumaevkova04/posts/pkg/leveledlog"
	"github.com/jumaevkova04/posts/pkg/response"
	"net/http"
)

type Handler struct {
	Config  *config.Config
	Logger  *leveledlog.Logger
	Service *service.Service
}

func (h *Handler) handle(handle func(w http.ResponseWriter, r *http.Request) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res response.Response

		res.Payload, res.Error = handle(w, r)

		var serverError response.ServerError
		ok := errors.As(res.Error, &serverError)
		if ok {
			h.Logger.Error(res.Error)
		}

		res.Send(w, r)
	}
}
