package handlers

import (
	"github.com/jumaevkova04/posts/internal/dto"
	"github.com/jumaevkova04/posts/pkg/request"
	"github.com/jumaevkova04/posts/pkg/response"
	"net/http"
)

// UserInfo godoc
// @Security 	ApiKeyAuth
// @Tags    	users
// @Accept  	json
// @Produce 	json
// @Success 	200     		{object} 		models.User
// @Failure  	400  			{object}  		response.ErrorResponse
// @Failure  	422  			{object}    	response.ErrorResponse
// @Failure  	500  			{object}    	response.ErrorResponse
// @Router 		/users/me						[get]
func (h *Handler) UserInfo(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
	input := new(dto.UserInfoRequest)

	input.ID = contextGetAuthenticatedUserID(r)

	data, err = h.Service.UserInfo(input)

	return
}

// UpdateUser godoc
// @Security 	ApiKeyAuth
// @Tags    	users
// @Accept  	json
// @Produce 	json
// @Param   	request 		body 	  		dto.UpdateUserRequest	 	true "request body"
// @Success 	200     		"No Content"
// @Failure  	400  			{object}  		response.ErrorResponse
// @Failure  	422  			{object}    	response.ErrorResponse
// @Failure  	500  			{object}    	response.ErrorResponse
// @Router 		/users/update					[put]
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
	input := new(dto.UpdateUserRequest)

	err = request.DecodeJSON(w, r, &input)
	if err != nil {
		return nil, response.NewBadRequestError(err)
	}

	input.ID = contextGetAuthenticatedUserID(r)

	err = h.Service.UpdateUserInfo(input)

	return
}
