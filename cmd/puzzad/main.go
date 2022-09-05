package main

import (
	"context"
	"flag"
	"os"

	"github.com/csmith/envflag"
	"github.com/greboid/puzzad/ent/game"
	"github.com/greboid/puzzad/ent/user"
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
	client := &puzzad.DBClient{}
	err := client.Init(*Debug)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed creating ent client")
	}
	defer func() {
		_ = client.Close()
	}()
	go func() {
		ctx := context.Background()
		u, _ := client.GetOrCreateUser(ctx, "test@test.test")
		log.Printf("Verify Code: %s", u.VerifyCode)
		ad, _ := client.CreateAdventure(ctx, "Test Adventure")
		ad, _ = client.GetAdventure(ctx, "Test Adventure")
		_, _ = client.AddPuzzle(ctx, ad, 0, "First puzzle", "one")
		_, _ = client.AddPuzzle(ctx, ad, 1, "Second puzzle", "two")
		_, _ = client.AddPuzzle(ctx, ad, 2, "Third puzzle", "three")
		g, _ := client.CreateGame(ctx, ad, u)
		_ = client.SetStatusForGame(ctx, g, game.StatusPaid)
		adventures, _ := client.GetPaidAdventuresForUser(ctx, u)
		for index := range adventures {
			ac, _ := adventures[index].QueryGame().Where(game.HasUserWith(user.ID(u.ID))).Only(ctx)
			log.Printf("Adventure codes: %s: %s", adventures[index].Name, ac.Code)
		}
	}()
	server := web.Webserver{
		Client: client,
	}
	server.Init(*WebPort, logger)
	log.Info().Int("port", *WebPort).Msg("Starting web server.")
	err = server.RunAndWait()
	if err != nil {
		log.Error().Err(err).Msg("Failed running web server.")
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
