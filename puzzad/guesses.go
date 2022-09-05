package puzzad

import (
	"context"
	"github.com/greboid/puzzad/ent/user"

	"github.com/greboid/puzzad/ent"
	"github.com/greboid/puzzad/ent/guess"
	"github.com/greboid/puzzad/ent/question"
)

func (db *DBClient) addGuess(ctx context.Context, q *ent.Question, u *ent.User, guess string) (*ent.Guess, error) {
	return db.entclient.Guess.Create().
		SetContent(guess).
		AddQuestion(q).
		AddTeam(u).
		Save(ctx)
}

func (db *DBClient) getGuessesForQuestion(ctx context.Context, q *ent.Question, u *ent.User) ([]*ent.Guess, error) {
	return db.entclient.Guess.Query().
		Where(guess.And(guess.HasTeamWith(user.ID(u.ID)), guess.HasQuestionWith(question.ID(q.ID)))).
		All(ctx)
}
