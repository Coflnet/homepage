package main

import (
	"log/slog"
	"os"

	"github.com/Coflnet/homepage/internal/api"
	"github.com/Coflnet/homepage/internal/usecase"
)

func main() {
	setupLogger()
	slog.Info("starting application..")

	config, err := usecase.NewConfig()
	if err != nil {
		panic(err)
	}
	slog.Info("config initialized")

	translator := usecase.NewTranslator()
	slog.Info("translator initialized")

	slog.Info("setting up webserver..")
	api := api.NewWebServer(config, translator)

	slog.Info("starting webserver..")
	err = api.StartServer()

	slog.Error("webserver stopped")
	if err != nil {
		panic(err)
	}
}

func setupLogger() {
	opts := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	textHandler := slog.NewTextHandler(os.Stdout, &opts)
	logger := slog.New(textHandler)
	slog.SetDefault(logger)
}
