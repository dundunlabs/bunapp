package main

import (
	"context"
	"log"
	"net/http"

	"github.com/dundunlabs/bunapp"
	"github.com/dundunlabs/bunapp/x/httperror"
	_ "github.com/joho/godotenv/autoload"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

func main() {
	ctx := context.Background()

	app := bunapp.New(ctx, bunapp.AppConfig{
		DB: newDB(),
		Middlewares: []bunrouter.MiddlewareFunc{
			reqlog.NewMiddleware(),
			httperror.NewMiddleware(),
		},
	})

	app.Router().GET("/ping", func(w http.ResponseWriter, req bunrouter.Request) error {
		_, err := w.Write([]byte("pong"))
		return err
	})

	log.Fatal(app.Run(":8080"))
}
