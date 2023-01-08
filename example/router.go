package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dundunlabs/bunapp/example/constant"
	"github.com/dundunlabs/bunapp/x/errorhdr"
	"github.com/dundunlabs/bunapp/x/panichdr"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

func initRouter() *bunrouter.Router {
	router := bunrouter.New(
		bunrouter.Use(
			reqlog.NewMiddleware(reqlog.FromEnv("BUNDEBUG")),
			panichdr.NewMiddleware(constant.HTTPErrors, constant.HTTPErrorFunc),
			errorhdr.NewMiddleware(constant.HTTPErrors, constant.HTTPErrorFunc),
		),
	)

	router.GET("/ping", func(w http.ResponseWriter, req bunrouter.Request) error {
		return bunrouter.JSON(w, "pong")
	})

	router.WithGroup("/panic", func(g *bunrouter.Group) {
		g.GET("/eof", func(w http.ResponseWriter, req bunrouter.Request) error {
			var v any
			panic(json.NewDecoder(req.Body).Decode(&v)) // return io.EOF error
		})

		g.GET("/not-found", func(w http.ResponseWriter, req bunrouter.Request) error {
			panic(sql.ErrNoRows)
		})

		g.GET("/internal", func(w http.ResponseWriter, req bunrouter.Request) error {
			panic("Oops")
		})
	})

	router.WithGroup("/errors", func(g *bunrouter.Group) {
		g.GET("/eof", func(w http.ResponseWriter, req bunrouter.Request) error {
			var v any
			return json.NewDecoder(req.Body).Decode(&v) // return io.EOF error
		})

		g.GET("/not-found", func(w http.ResponseWriter, req bunrouter.Request) error {
			return sql.ErrNoRows
		})

		g.GET("/internal", func(w http.ResponseWriter, req bunrouter.Request) error {
			return errors.New("Unhandled error")
		})
	})

	return router
}
