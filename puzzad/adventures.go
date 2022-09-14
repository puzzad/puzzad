package puzzad

import (
	"context"

	"github.com/greboid/puzzad/ent"
)

type AdventureDatabase interface {
	GetAllPublicAdventures(ctx context.Context) ([]*ent.Adventure, error)
	GetAllAdventures(ctx context.Context) ([]*ent.Adventure, error)
	GetAdventureByID(ctx context.Context, id int) (*ent.Adventure, []*ent.Puzzle, error)
	CreateAdventure(ctx context.Context, name string) (*ent.Adventure, error)
}

type AdventureManager struct {
	db AdventureDatabase
}

func NewAdventureManager(db AdventureDatabase) *AdventureManager {
	return &AdventureManager{
		db: db,
	}
}

func (am *AdventureManager) GetAllPublicAdventures(ctx context.Context) ([]*ent.Adventure, error) {
	return am.db.GetAllPublicAdventures(ctx)
}

func (am *AdventureManager) GetAllAdventures(ctx context.Context) ([]*ent.Adventure, error) {
	return am.db.GetAllAdventures(ctx)
}

func (am *AdventureManager) GetAdventureByID(ctx context.Context, id int) (*ent.Adventure, []*ent.Puzzle, error) {
	return am.db.GetAdventureByID(ctx, id)
}

func (am *AdventureManager) CreateAdventure(ctx context.Context, name string) (*ent.Adventure, error) {
	return am.db.CreateAdventure(ctx, name)
}
