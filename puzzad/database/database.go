package database

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/greboid/puzzad/ent"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type DBClient struct {
	entclient *ent.Client
	debug     bool
	Dialect   string
	DSN       string
}

func (db *DBClient) Init() error {
	var err error
	db.entclient, err = ent.Open(db.Dialect, db.DSN)
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
