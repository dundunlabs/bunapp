package main

import (
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

func initRouter() *bunrouter.Router {
	return bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware(reqlog.FromEnv("BUNDEBUG"))),
	)
}
