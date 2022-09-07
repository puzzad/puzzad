package main

import (
	"context"
	"flag"
	"os"

	"github.com/csmith/envflag"
	"github.com/greboid/puzzad/ent/game"
	"github.com/greboid/puzzad/ent/user"
	"github.com/greboid/puzzad/puzzad/database"
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
	client := &database.DBClient{}
	err := client.Init(*Debug)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed creating ent client")
	}
	defer func() {
		_ = client.Close()
	}()
	go func() {
		ctx := context.Background()
		u, err := client.GetOrCreateUser(ctx, "test@test.test")
		if err != nil {
			panic(err)
		}
		log.Printf("Verify Code: %s", u.VerifyCode)
		ad, err := client.CreateAdventure(ctx, "Test Adventure")
		if err != nil {
			panic(err)
		}
		ad, err = client.GetAdventure(ctx, "Test Adventure")
		if err != nil {
			panic(err)
		}
		_, err = client.AddPuzzle(ctx, ad, 0, "First puzzle", "one")
		if err != nil {
			panic(err)
		}
		_, err = client.AddPuzzle(ctx, ad, 1, "Second puzzle", "two")
		if err != nil {
			panic(err)
		}
		_, err = client.AddPuzzle(ctx, ad, 2, "Third puzzle", "three")
		if err != nil {
			panic(err)
		}
		g, err := client.CreateGame(ctx, ad, u)
		if err != nil {
			panic(err)
		}
		err = client.SetStatusForGame(ctx, g, game.StatusPaid)
		if err != nil {
			panic(err)
		}
		adventures, err := client.GetPaidAdventuresForUser(ctx, u)
		if err != nil {
			panic(err)
		}
		for index := range adventures {
			ac, err := adventures[index].QueryGame().Where(game.HasUserWith(user.ID(u.ID))).Only(ctx)
			if err != nil {
				panic(err)
			}
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
