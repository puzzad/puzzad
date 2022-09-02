package puzzad

import (
	"context"
	"fmt"

	"github.com/greboid/puzzad/ent"
	"github.com/greboid/puzzad/ent/team"
	_ "github.com/mattn/go-sqlite3"
)

func CreateTeam(ctx context.Context, client *ent.Client, name string, email string) (*ent.Team, error) {
	t, err := client.Team.
		Create().
		SetName(name).
		SetEmail(email).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating team: %w", err)
	}
	return t, nil
}

func QueryTeam(ctx context.Context, client *ent.Client, name string) (*ent.Team, error) {
	t, err := client.Team.
		Query().
		Where(team.Name(name)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying team: %w", err)
	}
	return t, nil
}
