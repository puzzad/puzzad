package puzzad

import (
	"context"
	"github.com/greboid/puzzad/ent/game"
	"github.com/greboid/puzzad/ent/user"

	"github.com/greboid/puzzad/ent"
	"github.com/greboid/puzzad/ent/adventure"
)

func (db *DBClient) GetPaidAdventuresForUser(ctx context.Context, u *ent.User) ([]*ent.Adventure, error) {
	return db.GetAdventuresForUser(ctx, u, game.StatusPaid)
}

func (db *DBClient) GetAdventuresForUser(ctx context.Context, u *ent.User, status game.Status) ([]*ent.Adventure, error) {
	return u.QueryGame().Where(game.StatusEQ(status)).QueryAdventures().All(ctx)
}

func (db *DBClient) CreateAdventure(ctx context.Context, name string) (*ent.Adventure, error) {
	return db.entclient.Adventure.Create().SetName(name).Save(ctx)
}

func (db *DBClient) GetAdventure(ctx context.Context, name string) (*ent.Adventure, error) {
	return db.entclient.Adventure.Query().Where(adventure.Name(name)).Only(ctx)
}

func (db *DBClient) AddAdventureForUser(ctx context.Context, a *ent.Adventure, u *ent.User) error {
	_, err := a.QueryUser().Where(user.ID(u.ID)).Only(ctx)
	if err != nil {
		_, err = u.Update().AddAdventures(a).Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *DBClient) GetGameForUser(ctx context.Context, a *ent.Adventure, u *ent.User) (*ent.Game, error) {
	// TODO: This won't work if a user has multiple games of the same adventure
	return u.QueryGame().Where(game.And(game.UserID(u.ID), game.AdventureID(a.ID))).Only(ctx)
}

func (db *DBClient) GetUserStatusForAdventure(ctx context.Context, a *ent.Adventure, u *ent.User) (game.Status, error) {
	// TODO: This won't work if a user has multiple games of the same adventure
	ac, err := db.GetGameForUser(ctx, a, u)
	if err != nil {
		return game.StatusUnpaid, err
	}
	return ac.Status, err
}

func (db *DBClient) SetUserStatusForAdventure(ctx context.Context, a *ent.Adventure, u *ent.User, status game.Status) error {
	ac, err := db.GetGameForUser(ctx, a, u)
	if err != nil {
		return err
	}
	_, err = ac.Update().SetStatus(status).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (db *DBClient) VerifyAdventureCode(ctx context.Context, code string) error {
	_, err := db.entclient.Game.Query().Where(game.Code(code)).Only(ctx)
	return err
}
