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
	return db.entclient.Question.DeleteOne(q).Exec(ctx)
}

func (db *DBClient) GetQuestionsForAdventure(ctx context.Context, a *ent.Adventure) ([]*ent.Question, error) {
	return db.entclient.Adventure.QueryQuestions(a).All(ctx)
}
