package puzzad

import (
	"context"

	"github.com/greboid/puzzad/ent"
	_ "github.com/mattn/go-sqlite3"
)

type DBClient struct {
	entclient *ent.Client
}

func (db *DBClient) Init() error {
	var err error
	db.entclient, err = ent.Open("sqlite3", "file:test.db?mode=rwc&_fk=1&_auto_vacuum=incremental")
	if err != nil {
		return err
	}
	err = db.entclient.Schema.Create(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (db *DBClient) Close() error {
	return db.entclient.Close()
}
