package constant

import (
	"database/sql"
	"io"
	"net/http"

	"github.com/uptrace/bunrouter"
)

type HTTPError struct {
	statusCode int

	Code    string `json:"code"`
	Message string `json:"message"`
}

func (err HTTPError) Error() string {
	return err.Message
}

func HTTPErrorFunc(err HTTPError) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		w.WriteHeader(err.statusCode)
		return bunrouter.JSON(w, err)
	}
}

var HTTPErrors = map[any]HTTPError{
	io.EOF: {
		statusCode: http.StatusBadRequest,

		Code:    "eof",
		Message: "EOF reading HTTP request body",
	},
	sql.ErrNoRows: {
		statusCode: http.StatusNotFound,

		Code:    "not_found",
		Message: "Record Not Found",
	},
}
