package response

import (
	"net/http"
)

type Response struct {
	Error   error `json:"error,omitempty"`
	Payload any   `json:"payload,omitempty"`
}

func (res *Response) Send(w http.ResponseWriter, r *http.Request) {
	if res.Error != nil {
		res.ErrorResponse(w, r)
		return
	}

	if err := JSON(w, http.StatusOK, res.Payload); err != nil {
		res.ErrorResponse(w, r)
		return
	}
}

func (res *Response) ErrorResponse(w http.ResponseWriter, r *http.Request) {
	switch res.Error.(type) {
	case BadRequestError:
		res.badRequest(w, r)
	case FailedValidationError:
		res.failedValidation(w, r)
	case InvalidAuthenticationTokenError:
		res.invalidAuthenticationToken(w, r)
	case AuthenticationRequiredError:
		res.authenticationRequired(w, r)
	case ServerError:
		res.serverError(w, r, res.Error)
	default:
		res.serverError(w, r, res.Error)
	}
}
