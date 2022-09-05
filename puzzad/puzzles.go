package puzzad

import (
	"context"

	"github.com/greboid/puzzad/ent"
)

func (db *DBClient) AddPuzzle(ctx context.Context, a *ent.Adventure, order int, title, answer string) (*ent.Puzzle, error) {
	return db.entclient.Puzzle.Create().
		SetTitle(title).
		SetAdventure(a).
		SetOrder(order).
		SetAnswer(answer).
		Save(ctx)
}

func (db *DBClient) UpdatePuzzle(ctx context.Context, p *ent.Puzzle, title, answer string) error {
	_, err := p.Update().
		SetTitle(title).
		SetAnswer(answer).
		Save(ctx)
	return err
}

func (db *DBClient) DeletePuzzle(ctx context.Context, p *ent.Puzzle) error {
	return db.entclient.Puzzle.DeleteOne(p).Exec(ctx)
}

func (db *DBClient) GetPuzzlesForAdventure(ctx context.Context, a *ent.Adventure) ([]*ent.Puzzle, error) {
	return db.entclient.Adventure.QueryPuzzles(a).All(ctx)
}
