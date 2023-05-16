package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/sprokhorov/cint-test/internal/config"
	"github.com/sprokhorov/cint-test/internal/server"
	"github.com/sprokhorov/cint-test/internal/storage"
)

func main() {
	log := zerolog.New(os.Stdout)

	cfg, err := config.New()
	if err != nil {
		log.Fatal().Msgf("Failed to load config, %v", err)
	}

	logLevel, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatal().Msgf("Failed to parse log level, %v", err)
	}
	log.Level(logLevel)

	srv := server.New(storage.NewInMemoryStorage(), &log)
	if err := srv.Listen(cfg.ServeHost, cfg.ServePort); err != nil {
		log.Fatal().Msg(err.Error())
	}
}
