package handlers

import (
	"github.com/jumaevkova04/posts/internal/dto"
	"github.com/jumaevkova04/posts/pkg/request"
	"github.com/jumaevkova04/posts/pkg/response"
	"net/http"
)

// CreateFollower godoc
// @Security 	ApiKeyAuth
// @Tags    	followers
// @Accept  	json
// @Produce 	json
// @Param   	request 		body 	  		dto.CreateFollowerRequest	 	true "request body"
// @Success 	200     		"No Content"
// @Failure  	400  			{object}  		response.ErrorResponse
// @Failure  	422  			{object}    	response.ErrorResponse
// @Failure  	500  			{object}    	response.ErrorResponse
// @Router 		/follow		 	[post]
func (h *Handler) CreateFollower(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
	input := new(dto.CreateFollowerRequest)

	err = request.DecodeJSON(w, r, &input)
	if err != nil {
		return nil, response.NewBadRequestError(err)
	}

	input.UserID = contextGetAuthenticatedUserID(r)

	err = h.Service.CreateFollower(input)

	return
}

// GetFollowingsID godoc
// @Security 	ApiKeyAuth
// @Tags    	followers
// @Accept  	json
// @Produce 	json
// @Success 	200     		{array} 		string
// @Failure  	400  			{object}  		response.ErrorResponse
// @Failure  	422  			{object}    	response.ErrorResponse
// @Failure  	500  			{object}    	response.ErrorResponse
// @Router 		/followings		[get]
func (h *Handler) GetFollowingsID(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
	input := new(dto.GetFollowingsIDRequest)

	input.UserID = contextGetAuthenticatedUserID(r)

	data, err = h.Service.GetFollowingsID(input)

	return
}
