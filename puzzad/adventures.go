package puzzad

import (
	"context"

	"github.com/greboid/puzzad/ent"
	"github.com/greboid/puzzad/ent/access"
	"github.com/greboid/puzzad/ent/adventure"
	"github.com/greboid/puzzad/ent/team"
)

func (db *DBClient) GetTeamPaidAdventures(ctx context.Context, team *ent.Team) ([]*ent.Adventure, error) {
	return db.GetTeamdAventures(ctx, team, access.StatusPaid)
}

func (db *DBClient) GetTeamdAventures(ctx context.Context, team *ent.Team, status access.Status) ([]*ent.Adventure, error) {
	return team.QueryAccess().Where(access.StatusEQ(status)).QueryAdventures().All(ctx)
}

func (db *DBClient) CreateAdventure(ctx context.Context, name string) (*ent.Adventure, error) {
	return db.entclient.Adventure.Create().SetName(name).Save(ctx)
}

func (db *DBClient) GetAdventure(ctx context.Context, name string) (*ent.Adventure, error) {
	return db.entclient.Adventure.Query().Where(adventure.Name(name)).Only(ctx)
}

func (db *DBClient) AddAdventureToTeam(ctx context.Context, a *ent.Adventure, t *ent.Team) error {
	_, err := a.QueryTeam().Where(team.ID(t.ID)).Only(ctx)
	if err != nil {
		_, err = t.Update().AddAdventures(a).Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *DBClient) GetTeamAdventureAccess(ctx context.Context, a *ent.Adventure, t *ent.Team) (*ent.Access, error) {
	return t.QueryAccess().Where(access.And(access.TeamID(t.ID), access.AdventureID(a.ID))).Only(ctx)
}

func (db *DBClient) CheckTeamAdventureStatus(ctx context.Context, a *ent.Adventure, t *ent.Team) (access.Status, error) {
	ac, err := db.GetTeamAdventureAccess(ctx, a, t)
	if err != nil {
		return access.StatusUnpaid, err
	}
	return ac.Status, err
}

func (db *DBClient) SetTeamAdventureStatus(ctx context.Context, a *ent.Adventure, t *ent.Team, status access.Status) error {
	ac, err := db.GetTeamAdventureAccess(ctx, a, t)
	if err != nil {
		return err
	}
	_, err = ac.Update().SetStatus(status).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
