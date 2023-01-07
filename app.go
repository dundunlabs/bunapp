package bunapp

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/uptrace/bun"
	"github.com/uptrace/bunrouter"
	"github.com/urfave/cli/v3"
)

type App struct {
	Router *bunrouter.Router
	DB     *bun.DB
}

func (app *App) NewServer(addr string) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: app.Router,
	}
}

func (app *App) WaitExitSignal() os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	return <-sig
}

func (app *App) NewCLI() *cli.App {
	return &cli.App{
		Name:  "bunapp",
		Usage: "A web application built with Bun ecosystems",
		Commands: []*cli.Command{
			{
				Name:    "start",
				Aliases: []string{"s"},
				Usage:   "Start application",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Value:   "8080",
					},
				},
				Action: func(ctx *cli.Context) error {
					addr := ":" + ctx.String("port")
					srv := app.NewServer(addr)

					go func() {
						if err := srv.ListenAndServe(); err != nil {
							fmt.Println("\n" + err.Error() + "\n")
						}
					}()
					fmt.Println("Application started on", addr)

					app.WaitExitSignal()

					return srv.Shutdown(ctx.Context)
				},
			},
		},
	}
}
