package main

import (
	"demo/internal/app"
	"demo/internal/pkg/config"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
