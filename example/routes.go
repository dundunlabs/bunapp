package main

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

var routes = map[string]bunrouter.HandlerFunc{
	"foo": func(w http.ResponseWriter, req bunrouter.Request) error {
		return nil
	},
}
