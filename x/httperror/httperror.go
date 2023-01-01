package httperror

import (
	"database/sql"
	"io"
	"net/http"

	"github.com/uptrace/bunrouter"
)

type HTTPError struct {
	statusCode int
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (e HTTPError) Error() string {
	return e.Message
}

func NewHTTPError(err error) HTTPError {
	switch err {
	case io.EOF:
		return HTTPError{
			statusCode: http.StatusBadRequest,

			Code:    "eof",
			Message: "EOF reading HTTP request body",
		}
	case sql.ErrNoRows:
		return HTTPError{
			statusCode: http.StatusNotFound,

			Code:    "not_found",
			Message: "Record Not Found",
		}
	}

	return HTTPError{
		statusCode: http.StatusInternalServerError,

		Code:    "internal",
		Message: "Internal server error",
	}
}

func NewMiddleware() bunrouter.MiddlewareFunc {
	return func(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
		return func(w http.ResponseWriter, req bunrouter.Request) error {
			err := next(w, req)

			switch err := err.(type) {
			case nil:
			case HTTPError:
				w.WriteHeader(err.statusCode)
				_ = bunrouter.JSON(w, err)
			default:
				httpErr := NewHTTPError(err)
				w.WriteHeader(httpErr.statusCode)
				_ = bunrouter.JSON(w, httpErr)
			}

			return err
		}
	}
}
