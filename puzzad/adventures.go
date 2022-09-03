package puzzad

import (
	"context"

	"github.com/greboid/puzzad/ent"
)

func (db *DBClient) GetTeamAdventures(ctx context.Context, team ent.Team) ([]*ent.Access, error) {
	access, err := team.QueryAccess().All(ctx)
	if err != nil {
		return nil, err
	}
	return access, err
}
