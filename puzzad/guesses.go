package puzzad

import (
	"context"

	"github.com/greboid/puzzad/ent"
	"github.com/greboid/puzzad/ent/guess"
	"github.com/greboid/puzzad/ent/question"
	"github.com/greboid/puzzad/ent/team"
)

func (db *DBClient) addGuess(ctx context.Context, q *ent.Question, t *ent.Team, guess string) (*ent.Guess, error) {
	g, err := db.entclient.Guess.Create().
		SetContent(guess).
		AddQuestion(q).
		AddTeam(t).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (db *DBClient) getGuessesForQuestionAndTeam(ctx context.Context, q *ent.Question, t *ent.Team) ([]*ent.Guess, error) {
	return db.entclient.Guess.Query().
		Where(guess.And(guess.HasTeamWith(team.ID(t.ID)), guess.HasQuestionWith(question.ID(q.ID)))).
		All(ctx)
}
