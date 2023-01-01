package bunapp

import (
	"net/http"
	"time"

	"github.com/uptrace/bunrouter"
)

type Router struct {
	*bunrouter.Router
}

func newRouter(opts ...bunrouter.Option) *Router {
	return &Router{bunrouter.New(opts...)}
}

func (router *Router) newHTTPServer(addr string) *http.Server {
	return &http.Server{
		Addr:         addr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}
}
