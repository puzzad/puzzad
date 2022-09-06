package database

import (
	"context"

	"github.com/greboid/puzzad/ent"
	_ "github.com/mattn/go-sqlite3"
)

type DBClient struct {
	entclient *ent.Client
	debug     bool
}

func (db *DBClient) Init(debug bool) error {
	var dsn string
	if debug {
		dsn = "file:ent?mode=memory&cache=shared&_fk=1"
	} else {
		dsn = "file:test.db?mode=rwc&_fk=1&_auto_vacuum=incremental"
	}
	var err error
	db.entclient, err = ent.Open("sqlite3", dsn)
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
