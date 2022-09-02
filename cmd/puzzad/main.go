package main

import (
	"log"

	"github.com/greboid/puzzad/puzzad"
	"github.com/greboid/puzzad/web"
)

func main() {
	client, err := puzzad.CreateEntClient(false)
	if err != nil {
		log.Fatalf("failed creating ent client: %v", err)
	}
	defer func() {
		_ = client.Close()
	}()
	h := web.CreateServer(3000)
	web.RunAndWait(h)
}
