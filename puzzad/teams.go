package puzzad

import (
	"context"
	"fmt"

	"github.com/greboid/puzzad/ent"
	"github.com/greboid/puzzad/ent/team"
	_ "github.com/mattn/go-sqlite3"
)

func (db *DBClient) CreateTeam(ctx context.Context, name string, email string) (*ent.Team, error) {
	t, err := db.entclient.Team.
		Create().
		SetName(name).
		SetEmail(email).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating team: %w", err)
	}
	return t, nil
}

func (db *DBClient) QueryTeam(ctx context.Context, name string) (*ent.Team, error) {
	t, err := db.entclient.Team.
		Query().
		Where(team.Name(name)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying team: %w", err)
	}
	return t, nil
}
