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
	AddPuzzle(ctx context.Context, a *ent.Adventure, order int, title, answer string) (*ent.Puzzle, error)
	GetPuzzleByID(ctx context.Context, id int) (*ent.Puzzle, error)
	DeletePuzzleByID(ctx context.Context, id int) error
	UpdatePuzzle(ctx context.Context, id int, title string, answer string)
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

func (am *AdventureManager) CreatePuzzle(ctx context.Context, adventureID int, puzzleTitle string) (*ent.Puzzle, error) {
	a, p, err := am.db.GetAdventureByID(ctx, adventureID)
	if err != nil {
		return nil, err
	}
	return am.db.AddPuzzle(ctx, a, len(p), puzzleTitle, "")
}

func (am *AdventureManager) GetPuzzleByID(ctx context.Context, id int) (*ent.Puzzle, error) {
	return am.db.GetPuzzleByID(ctx, id)
}

func (am *AdventureManager) DeletePuzzleByID(ctx context.Context, id int) error {
	return am.db.DeletePuzzleByID(ctx, id)
}
