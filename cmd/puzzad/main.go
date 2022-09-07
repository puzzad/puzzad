package main

import (
	"context"
	"flag"
	"net/mail"
	"os"

	"github.com/csmith/envflag"
	"github.com/greboid/puzzad/puzzad"
	"github.com/greboid/puzzad/puzzad/database"
	"github.com/greboid/puzzad/web"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	WebPort       = flag.Int("web-port", 3000, "Port for webserver")
	Debug         = flag.Bool("debug", true, "Enable debug logging")
	AdminEmail    = flag.String("admin-email", "", "Default admin email, only used if at least one admin does not exist, must be accompanied by admin-password")
	AdminPassword = flag.String("admin-password", "", "Default admin password, only used if at least one admin does not exist, must be accompanied by admin-email")
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
		// TODO: This should use the user manager
		admins, err := client.GetAdmins(ctx)
		if err != nil {
			log.Fatal().Err(err).Msg("Unable to check if admins exist")
		}
		if len(admins) == 0 {
			log.Info().Str("email", *AdminEmail).Msg("Creating initial admin user.")
			_, err = mail.ParseAddress(*AdminEmail)
			if err != nil {
				log.Fatal().Err(err).Msg("Admin email must be valid")
			}
			if len(*AdminPassword) == 0 {
				log.Fatal().Err(err).Msg("Unable to hash admin password, not set")
			}
			passHash, err := puzzad.GetHash(*AdminPassword)
			if err != nil {
				log.Fatal().Err(err).Msg("Unable to hash admin password")
			}
			user, err := client.CreateUser(ctx, *AdminEmail)
			if err != nil {
				log.Fatal().Err(err).Msg("Unable to create admin user")
			}
			err = client.SetPassword(ctx, user, passHash)
			if err != nil {
				log.Fatal().Err(err).Msg("Unable to set admin password")
			}
			err = client.SetAdmin(ctx, user, true)
			if err != nil {
				log.Fatal().Err(err).Msg("Unable to set user as admin")
			}
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
