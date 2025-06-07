package handlers

import (
	"github.com/jumaevkova04/posts/internal/dto"
	"github.com/jumaevkova04/posts/pkg/request"
	"github.com/jumaevkova04/posts/pkg/response"
	"net/http"
)

// OTPToRegister godoc
// @Tags    	register
// @Accept  	json
// @Produce 	json
// @Param   	request 		body 	  	dto.OTPToRegisterRequest		 			true "request body"
// @Success 	200  			"No Content"
// @Failure  	400  			{object}  	response.ErrorResponse
// @Failure  	422  			{object}    response.ErrorResponse
// @Failure  	500  			{object}    response.ErrorResponse
// @Router 		/register/otp 				[post]
func (h *Handler) OTPToRegister(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
	input := new(dto.OTPToRegisterRequest)

	err = request.DecodeJSON(w, r, &input)
	if err != nil {
		return nil, response.NewBadRequestError(err)
	}

	err = h.Service.OTPToRegister(input)

	return
}

// CheckRegisterOTP godoc
// @Tags    	register
// @Accept  	json
// @Produce 	json
// @Param   	request 		body 	  	dto.CheckRegisterOTPRequest		 			true "request body"
// @Success 	200  			{object}    dto.CheckRegisterOTPResponse
// @Failure  	400  			{object}  	response.ErrorResponse
// @Failure  	422  			{object}    response.ErrorResponse
// @Failure  	500  			{object}    response.ErrorResponse
// @Router 		/register/check_otp 		[post]
func (h *Handler) CheckRegisterOTP(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
	input := new(dto.CheckRegisterOTPRequest)

	err = request.DecodeJSON(w, r, &input)
	if err != nil {
		return nil, response.NewBadRequestError(err)
	}

	data, err = h.Service.CheckRegisterOTP(input)

	return
}

// Register godoc
// @Tags    	register
// @Accept  	json
// @Produce 	json
// @Param   	request 		body 	  	dto.RegisterRequest		 			true "request body"
// @Success 	200  			"No Content"
// @Failure  	400  			{object}  	response.ErrorResponse
// @Failure  	422  			{object}    response.ErrorResponse
// @Failure  	500  			{object}    response.ErrorResponse
// @Router 		/register 					[post]
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
	input := new(dto.RegisterRequest)

	err = request.DecodeJSON(w, r, &input)
	if err != nil {
		return nil, response.NewBadRequestError(err)
	}

	err = h.Service.Register(input)

	return
}
