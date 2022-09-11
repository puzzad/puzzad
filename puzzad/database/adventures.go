package database

import (
	"context"

	"github.com/greboid/puzzad/ent"
	"github.com/greboid/puzzad/ent/adventure"
	"github.com/greboid/puzzad/ent/game"
	"github.com/greboid/puzzad/ent/puzzle"
)

func (db *DBClient) GetPaidAdventuresForUser(ctx context.Context, u *ent.User) ([]*ent.Adventure, error) {
	return db.GetAdventuresForUser(ctx, u, game.StatusPaid)
}

func (db *DBClient) GetAdventuresForUser(ctx context.Context, u *ent.User, status game.Status) ([]*ent.Adventure, error) {
	return u.QueryGame().Where(game.StatusEQ(status)).QueryAdventure().All(ctx)
}

func (db *DBClient) CreateAdventure(ctx context.Context, name string) (*ent.Adventure, error) {
	return db.entclient.Adventure.Create().SetName(name).Save(ctx)
}

func (db *DBClient) GetAdventure(ctx context.Context, name string) (*ent.Adventure, error) {
	return db.entclient.Adventure.Query().Where(adventure.Name(name)).Only(ctx)
}

func (db *DBClient) CreateGame(ctx context.Context, a *ent.Adventure, u *ent.User) (*ent.Game, error) {
	firstPuzzle, err := db.GetNextPuzzleForAdventure(ctx, a, -1)
	if err != nil {
		return nil, err
	}

	return db.entclient.Game.Create().
		SetAdventure(a).
		SetUser(u).
		SetCurrentPuzzle(firstPuzzle).
		Save(ctx)
}

func (db *DBClient) SetStatusForGame(ctx context.Context, g *ent.Game, status game.Status) error {
	_, err := g.Update().SetStatus(status).Save(ctx)
	return err
}

func (db *DBClient) VerifyAdventureCode(ctx context.Context, code string) error {
	_, err := db.entclient.Game.Query().Where(game.Code(code)).Only(ctx)
	return err
}

func (db *DBClient) GetNextPuzzleForAdventure(ctx context.Context, a *ent.Adventure, current int) (*ent.Puzzle, error) {
	return a.QueryPuzzles().Order(ent.Asc("order")).Where(puzzle.OrderGT(current)).First(ctx)
}

func (db *DBClient) GetAllPublicAdventures(ctx context.Context) ([]*ent.Adventure, error) {
	return db.entclient.Adventure.Query().Where(adventure.Public(true)).All(ctx)
}
