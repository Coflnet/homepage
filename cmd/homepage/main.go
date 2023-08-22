package main

import (
	"github.com/Coflnet/homepage/internal/api"
	"log/slog"
)

func main() {
	slog.Info("starting application..")

	err := api.StartWebserver()
	if err != nil {
		panic(err)
	}
}
