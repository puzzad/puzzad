package puzzad

import (
	"context"

	"github.com/greboid/puzzad/ent"
	"github.com/greboid/puzzad/ent/team"
	_ "github.com/mattn/go-sqlite3"
)

func (db *DBClient) CreateTeam(ctx context.Context, name string, email string) (*ent.Team, error) {
	return db.entclient.Team.
		Create().
		SetName(name).
		SetEmail(email).
		Save(ctx)
}

func (db *DBClient) GetTeam(ctx context.Context, name string) (*ent.Team, error) {
	return db.entclient.Team.
		Query().
		Where(team.Name(name)).
		Only(ctx)
}

func (db *DBClient) UpdateTeamStatus(ctx context.Context, t *ent.Team, newStatus team.Status) error {
	_, err := t.Update().SetStatus(newStatus).Save(ctx)
	return err
}
