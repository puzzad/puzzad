package main

import (
	"flag"
	"os"

	"github.com/csmith/envflag"
	"github.com/greboid/puzzad/puzzad"
	"github.com/greboid/puzzad/web"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	WebPort = flag.Int("web-port", 3000, "Port for webserver")
	Debug   = flag.Bool("debug", true, "Enable debug logging")
)

func main() {
	envflag.Parse()
	logger := createLogger(*Debug)
	client, err := puzzad.CreateEntClient(*Debug)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed creating ent client")
	}
	defer func() {
		_ = client.Close()
	}()
	server := web.Webserver{}
	server.Init(*WebPort, logger)
	log.Info().Msgf("Starting server: %d", *WebPort)
	err = server.RunAndWait()
	if err != nil {
		log.Error().Msgf("Error: %s", err)
	} else {
		log.Info().Msgf("Server stopped.")
	}
}

func createLogger(debug bool) *zerolog.Logger {
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = logger
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	return &logger
}
