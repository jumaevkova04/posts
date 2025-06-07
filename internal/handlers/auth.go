package handlers

import (
	"github.com/jumaevkova04/posts/internal/dto"
	"github.com/jumaevkova04/posts/pkg/request"
	"github.com/jumaevkova04/posts/pkg/response"
	"net/http"
)

// Login godoc
// @Tags    	auth
// @Accept  	json
// @Produce 	json
// @Param   	request 		body 	  	dto.LoginRequest		 			true "request body"
// @Success 	200  			{object}   	dto.LoginResponse
// @Failure  	400  			{object}  	response.ErrorResponse
// @Failure  	422  			{object}    response.ErrorResponse
// @Failure  	500  			{object}    response.ErrorResponse
// @Router 		/login 						[post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
	input := new(dto.LoginRequest)

	err = request.DecodeJSON(w, r, &input)
	if err != nil {
		return nil, response.NewBadRequestError(err)
	}

	data, err = h.Service.Login(input)

	return
}
