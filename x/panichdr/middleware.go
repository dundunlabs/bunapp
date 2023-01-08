package panichdr

import (
	"net/http"
	"runtime/debug"

	"github.com/uptrace/bunrouter"
)

func NewMiddleware[E error](errMap map[any]E, errFunc func(e E) bunrouter.HandlerFunc) bunrouter.MiddlewareFunc {
	return func(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
		return func(w http.ResponseWriter, req bunrouter.Request) error {
			defer func() {
				if err := recover(); err != nil {
					if err, ok := err.(E); ok {
						errFunc(err)(w, req)
						return
					}

					if err, ok := errMap[err]; ok {
						errFunc(err)(w, req)
						return
					}

					debug.PrintStack()
					w.WriteHeader(http.StatusInternalServerError)
				}
			}()

			return next(w, req)
		}
	}
}
