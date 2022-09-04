package puzzad

import (
	"context"

	"github.com/greboid/puzzad/ent"
)

func (db *DBClient) AddQuestion(ctx context.Context, a *ent.Adventure, title string) (*ent.Question, error) {
	q, err := db.entclient.Question.Create().
		SetTitle(title).
		SetAdventure(a).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return q, nil
}

func (db *DBClient) UpdateQuestion(ctx context.Context, q *ent.Question, title, answer string) error {
	_, err := q.Update().
		SetTitle(title).
		SetAnswer(answer).
		Save(ctx)
	return err
}

func (db *DBClient) DeleteQuestion(ctx context.Context, q *ent.Question) error {
	a, err := q.QueryAdventure().Only(ctx)
	if err != nil {
		return err
	}
	_, err = a.Update().RemoveQuestions(q).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
