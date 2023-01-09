package handler

import "github.com/dundunlabs/bunapp"

type GenericHandler struct {
	app *bunapp.App
}

func (h GenericHandler) App() *bunapp.App {
	return h.app
}
