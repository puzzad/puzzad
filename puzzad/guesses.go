package puzzad

import (
	"context"

	"github.com/greboid/puzzad/ent/puzzle"
	"github.com/greboid/puzzad/ent/user"

	"github.com/greboid/puzzad/ent"
	"github.com/greboid/puzzad/ent/guess"
)

func (db *DBClient) addGuess(ctx context.Context, p *ent.Puzzle, u *ent.User, guess string) (*ent.Guess, error) {
	return db.entclient.Guess.Create().
		SetContent(guess).
		AddPuzzle(p).
		AddTeam(u).
		Save(ctx)
}

func (db *DBClient) getGuessesForQuestion(ctx context.Context, p *ent.Puzzle, u *ent.User) ([]*ent.Guess, error) {
	return db.entclient.Guess.Query().
		Where(guess.And(guess.HasTeamWith(user.ID(u.ID)), guess.HasPuzzleWith(puzzle.ID(p.ID)))).
		All(ctx)
}
