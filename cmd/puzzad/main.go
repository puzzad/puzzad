package main

import (
	"context"
	"flag"
	"os"

	"github.com/csmith/envflag"
	"github.com/greboid/puzzad/puzzad"
	"github.com/greboid/puzzad/puzzad/database"
	"github.com/greboid/puzzad/web"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	WebURL        = flag.String("web-url", "http://localhost:3000", "Base URL used for the site")
	WebPort       = flag.Int("web-port", 3000, "Port for webserver")
	Debug         = flag.Bool("debug", true, "Enable debug logging")
	AdminEmail    = flag.String("admin-email", "", "Default admin email, only used if at least one admin does not exist, must be accompanied by admin-password")
	AdminPassword = flag.String("admin-password", "", "Default admin password, only used if at least one admin does not exist, must be accompanied by admin-email")
	SessionKey    = flag.String("session-key", "", "Encryption key for sessions, should be 32 bytes")

	SmtpUser     = flag.String("smtp_user", "", "SMTP Username")
	SmtpPassword = flag.String("smtp_password", "", "SMTP Password")
	SmtpServer   = flag.String("smtp_server", "", "SMTP Server")
	SmtpPort     = flag.Int("smtp_port", 25, "SMTP Server port")
	SmtpFrom     = flag.String("smtp_from", "", "SMTP From address")
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
	mailer := &puzzad.MailClient{
		SMTPUser:     *SmtpUser,
		SMTPPassword: *SmtpPassword,
		SMTPServer:   *SmtpServer,
		SMTPPort:     *SmtpPort,
		SMTPFrom:     *SmtpFrom,
		URL:          *WebURL,
	}
	userManager := puzzad.NewUserManager(client, mailer)
	adventureManager := puzzad.NewAdventureManager(client)

	if err := userManager.EnsureAdminExists(context.Background(), *AdminEmail, *AdminPassword); err != nil {
		log.Fatal().Err(err).Msg("Unable to create default admin account")
	}

	if *SessionKey == "" {
		log.Fatal().Msg("Session key must be set")
	}

	server := web.Webserver{
		Client:           client,
		UserManager:      userManager,
		SessionKey:       *SessionKey,
		AdventureManager: adventureManager,
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
