package puzzad

import (
	"context"

	"github.com/greboid/puzzad/ent"
)

type AdventureDatabase interface {
	GetAllAdventures(ctx context.Context) ([]*ent.Adventure, error)
}

type AdventureManager struct {
	db AdventureDatabase
}

func NewAdventureManager(db AdventureDatabase) *AdventureManager {
	return &AdventureManager{
		db: db,
	}
}

func (am *AdventureManager) GetAllAdventures(ctx context.Context) ([]*ent.Adventure, error) {
	return am.db.GetAllAdventures(ctx)
}
