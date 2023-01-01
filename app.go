package bunapp

import (
	"context"
	"errors"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bunrouter"
)

type App struct {
	ctx    context.Context
	router *Router
	db     *DB
}

type AppConfig struct {
	DB          *bun.DB
	Middlewares []bunrouter.MiddlewareFunc
}

func New(ctx context.Context, config AppConfig) *App {
	db := newDB(config.DB)
	router := newRouter(
		bunrouter.Use(config.Middlewares...),
	)

	return &App{
		ctx:    ctx,
		db:     db,
		router: router,
	}
}

func (app *App) Router() *Router {
	return app.router
}

func (app *App) DB() *DB {
	return app.db
}

func (app *App) Run(addr string) error {
	srv := app.router.newHTTPServer(addr)

	go srv.ListenAndServe()
	log.Println("app started on", srv.Addr)

	waitExitSignal()

	srv.Shutdown(app.ctx)
	log.Println("shutted down HTTP server")

	app.db.Close()
	log.Println("closed DB connection")

	return errors.New("app stopped")
}
