package response

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound                                = errors.New("NOT_FOUND")
	ErrIncorrectData                           = errors.New("INCORRECT_DATA")
	ErrRequiresAuthentication                  = errors.New("REQUIRES_AUTHENTICATION")
	ErrMethodNotAllowed                        = errors.New("METHOD_NOT_ALLOWED")
	ErrThisMethodIsNotSupportedForThisResource = errors.New("THIS_METHOD_IS_NOT_SUPPORTED_FOR_THIS_RESOURCE")
	ErrServerError                             = errors.New("SERVER_ERROR")
	ErrAuthenticationRequired                  = errors.New("AUTHENTICATION_REQUIRED")
	ErrInvalidAuthenticationToken              = errors.New("INVALID_AUTHENTICATION_TOKEN")
	ErrBadRequest                              = errors.New("BAD_REQUEST")
	ErrResourceNotFound                        = errors.New("RESOURCE_NOT_FOUND")
	ErrEmailAlreadyInUse                       = errors.New("EMAIL_ALREADY_IN_USE")
	ErrInvalidOTP                              = errors.New("INVALID_OTP")
	ErrInvalidEmail                            = errors.New("INVALID_EMAIL")
	ErrInvalidNonce                            = errors.New("INVALID_NONCE")
	ErrInvalidCredentials                      = errors.New("INVALID_EMAIL_OR_PASSWORD")
	ErrYouCantFollowYourself                   = errors.New("YOU_CANT_FOLLOW_YOURSELF")
	ErrFollowerExists                          = errors.New("FOLLOWER_EXISTS")
)

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func (res *Response) errorMessage(w http.ResponseWriter, r *http.Request, status int, err error, message string, headers http.Header) {
	data := ErrorResponse{
		Message: message,
		Error:   err.Error(),
	}

	if err := JSONWithHeaders(w, status, data, headers); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// ---------------------------------------------------------------------------------------

func NotFound(w http.ResponseWriter, r *http.Request) {
	message := ErrNotFound.Error()
	err := ErrResourceNotFound

	data := struct {
		Message string `json:"message"`
		Error   string `json:"error"`
	}{
		Message: message,
		Error:   err.Error(),
	}

	if err := JSON(w, http.StatusNotFound, data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// ---------------------------------------------------------------------------------------

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := ErrMethodNotAllowed.Error()
	err := ErrThisMethodIsNotSupportedForThisResource

	data := struct {
		Message string `json:"message"`
		Error   string `json:"error"`
	}{
		Message: message,
		Error:   err.Error(),
	}

	if err := JSON(w, http.StatusMethodNotAllowed, data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// ---------------------------------------------------------------------------------------

type BadRequestError struct {
	Err error `json:"err"`
}

func (e BadRequestError) Error() string {
	return e.Err.Error()
}

func NewBadRequestError(err error) BadRequestError {
	return BadRequestError{Err: err}
}

func (res *Response) badRequest(w http.ResponseWriter, r *http.Request) {
	message := ErrBadRequest.Error()
	err := res.Error

	res.errorMessage(w, r, http.StatusBadRequest, err, message, nil)
}

// ---------------------------------------------------------------------------------------

type FailedValidationError struct {
	Err error `json:"err"`
}

func (e FailedValidationError) Error() string {
	return e.Err.Error()
}

func NewFailedValidationError(err error) FailedValidationError {
	return FailedValidationError{Err: err}
}

func (res *Response) failedValidation(w http.ResponseWriter, r *http.Request) {
	message := ErrIncorrectData.Error()
	err := res.Error

	res.errorMessage(w, r, http.StatusUnprocessableEntity, err, message, nil)
}

// ---------------------------------------------------------------------------------------

type InvalidAuthenticationTokenError struct {
	Err error `json:"err"`
}

func (e InvalidAuthenticationTokenError) Error() string {
	return e.Err.Error()
}

func NewInvalidAuthenticationTokenError(err error) InvalidAuthenticationTokenError {
	return InvalidAuthenticationTokenError{Err: err}
}

func (res *Response) invalidAuthenticationToken(w http.ResponseWriter, r *http.Request) {
	headers := make(http.Header)
	headers.Set("WWW-Authenticate", "Bearer")

	message := ErrInvalidAuthenticationToken.Error()
	err := ErrInvalidAuthenticationToken

	res.errorMessage(w, r, http.StatusUnauthorized, err, message, headers)
}

// ---------------------------------------------------------------------------------------

type AuthenticationRequiredError struct {
	Err error `json:"err"`
}

func (e AuthenticationRequiredError) Error() string {
	return e.Err.Error()
}

func NewAuthenticationRequiredError(err error) AuthenticationRequiredError {
	return AuthenticationRequiredError{Err: err}
}

func (res *Response) authenticationRequired(w http.ResponseWriter, r *http.Request) {
	message := ErrRequiresAuthentication.Error()
	err := ErrAuthenticationRequired

	res.errorMessage(w, r, http.StatusUnauthorized, err, message, nil)
}

// ---------------------------------------------------------------------------------------

type ServerError struct {
	Err error `json:"err"`
}

func (e ServerError) Error() string {
	return e.Err.Error()
}

func NewServerError(err error) ServerError {
	return ServerError{Err: err}
}

func (res *Response) serverError(w http.ResponseWriter, r *http.Request, err error) {
	message := ErrServerError.Error()

	var serverError ServerError
	errors.As(res.Error, &serverError)
	err = ErrServerError

	res.errorMessage(w, r, http.StatusInternalServerError, err, message, nil)
}
