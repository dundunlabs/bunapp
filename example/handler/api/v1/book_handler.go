package v1

import (
	"net/http"

	"github.com/dundunlabs/bunapp/example/handler"
	"github.com/uptrace/bunrouter"
)

type BookHandler struct {
	handler.GenericHandler
}

func (h BookHandler) Index(w http.ResponseWriter, req bunrouter.Request) error {
	return h.App().DB.Ping()
}
