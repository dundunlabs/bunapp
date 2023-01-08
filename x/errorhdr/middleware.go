package errorhdr

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

func NewMiddleware[E error](errMap map[any]E, errFunc func(e E) bunrouter.HandlerFunc) bunrouter.MiddlewareFunc {
	return func(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
		return func(w http.ResponseWriter, req bunrouter.Request) error {
			err := next(w, req)

			if err == nil {
				return nil
			}

			if err, ok := err.(E); ok {
				return errFunc(err)(w, req)
			}

			if err, ok := errMap[err]; ok {
				return errFunc(err)(w, req)
			}

			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
	}
}
