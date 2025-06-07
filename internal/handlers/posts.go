package handlers

import (
	"github.com/jumaevkova04/posts/internal/dto"
	"github.com/jumaevkova04/posts/pkg/request"
	"github.com/jumaevkova04/posts/pkg/response"
	"net/http"
)

// CreatePost godoc
// @Security 	ApiKeyAuth
// @Tags    	posts
// @Accept  	json
// @Produce 	json
// @Param   	request 		body 	  		dto.CreatePostRequest	 	true "request body"
// @Success 	200     		"No Content"
// @Failure  	400  			{object}  		response.ErrorResponse
// @Failure  	422  			{object}    	response.ErrorResponse
// @Failure  	500  			{object}    	response.ErrorResponse
// @Router 		/posts		 	[post]
func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
	input := new(dto.CreatePostRequest)

	err = request.DecodeJSON(w, r, &input)
	if err != nil {
		return nil, response.NewBadRequestError(err)
	}

	input.UserID = contextGetAuthenticatedUserID(r)

	err = h.Service.CreatePost(input)

	return
}

// GetPosts godoc
// @Security 	ApiKeyAuth
// @Tags    	posts
// @Accept  	json
// @Produce 	json
// @Param   	request 		query 	  		dto.GetPostsRequest	 	true "request body"
// @Success 	200     		{array} 		models.Post
// @Failure  	400  			{object}  		response.ErrorResponse
// @Failure  	422  			{object}    	response.ErrorResponse
// @Failure  	500  			{object}    	response.ErrorResponse
// @Router 		/posts		[get]
func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
	input := new(dto.GetPostsRequest)

	err = request.DecodeFormQuery(r, &input)
	if err != nil {
		return nil, response.NewBadRequestError(err)
	}

	input.UserID = contextGetAuthenticatedUserID(r)

	data, err = h.Service.GetPosts(input)

	return
}
