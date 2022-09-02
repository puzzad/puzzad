package main

import (
	"context"
	"log"

	"github.com/greboid/puzzad/ent"
	"github.com/greboid/puzzad/puzzad"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer func() {
		_ = client.Close()
	}()
	createTeams := []string{"#mdbot", "test🔴", "123456789012345678901"}
	lookupTeams := []string{"#mdbot", "test🔴", "test2"}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	for index := range createTeams {
		log.Printf("Creating team: %s", createTeams[index])
		_, err = puzzad.CreateTeam(context.Background(), client, createTeams[index], createTeams[index]+"@example.com")
		if err != nil {
			log.Printf("failed to create t: %v", err)
		}
	}
	for index := range lookupTeams {
		t, err := puzzad.QueryTeam(context.Background(), client, lookupTeams[index])
		if err != nil {
			log.Printf("Error finding team: %s", lookupTeams[index])
		} else {
			log.Printf("Team: %s", t.Name)
		}
	}
}
