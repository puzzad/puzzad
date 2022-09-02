package puzzad

import (
	"github.com/greboid/puzzad/ent"
	_ "github.com/mattn/go-sqlite3"
)

func CreateEntClient(debug bool) (*ent.Client, error) {
	var client *ent.Client
	var err error
	if debug {
		client, err = ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	} else {
		client, err = ent.Open("sqlite3", "file:test.db?mode=rwc&_fk=1&_auto_vacuum=incremental")
	}
	return client, err
}
