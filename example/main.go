package main

import (
	"os"

	"github.com/dundunlabs/bunapp"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := &bunapp.App{
		Router: initRouter(),
		DB:     initDB(),
	}
	cli := app.NewCLI()
	cli.Run(os.Args)
}
