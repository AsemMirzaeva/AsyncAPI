package main

import (
	"asyncapi/apiserver"
	"asyncapi/config"
	"asyncapi/store"
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	conf, err := config.New()
	if err != nil {
		return err
	}

	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(jsonHandler)

	db, err := store.NewPostgresDb(conf)
	if err != nil {
		return err
	}
	dataStore := store.New(db)
	server := apiserver.New(conf, logger, dataStore)
	if err := server.Start(ctx); err != nil {
		return err
	}

	return nil
}
