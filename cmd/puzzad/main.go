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
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	client, err := puzzad.CreateEntClient(*Debug)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed creating ent client")
	}
	defer func() {
		_ = client.Close()
	}()
	server := web.Webserver{}
	server.Init(*WebPort)
	log.Info().Msgf("Starting server: %d", *WebPort)
	err = server.RunAndWait()
	if err != nil {
		log.Error().Msgf("Error: %s", err)
	} else {
		log.Info().Msgf("Server stopped.")
	}
}
