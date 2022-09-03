package main

import (
	"flag"
	"log"

	"github.com/csmith/envflag"
	"github.com/greboid/puzzad/puzzad"
	"github.com/greboid/puzzad/web"
)

var (
	WebPort = flag.Int("web-port", 3000, "Port for webserver")
	Debug   = flag.Bool("debug", true, "Enable debug logging")
)

func main() {
	envflag.Parse()
	client, err := puzzad.CreateEntClient(*Debug)
	if err != nil {
		log.Fatalf("failed creating ent client: %v", err)
	}
	defer func() {
		_ = client.Close()
	}()
	server := web.Webserver{}
	server.Init(*WebPort)
	log.Printf("Starting server: %d", *WebPort)
	err = server.RunAndWait()
	if err != nil {
		log.Printf("Error: %s", err)
	} else {
		log.Printf("Server stopped.")
	}
}
