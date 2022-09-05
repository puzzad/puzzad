package puzzad

import (
	"context"
	"github.com/greboid/puzzad/ent/user"

	"github.com/greboid/puzzad/ent"
	"github.com/greboid/puzzad/ent/access"
	"github.com/greboid/puzzad/ent/adventure"
)

func (db *DBClient) GetPaidAdventuresForUser(ctx context.Context, u *ent.User) ([]*ent.Adventure, error) {
	return db.GetAdventuresForUser(ctx, u, access.StatusPaid)
}

func (db *DBClient) GetAdventuresForUser(ctx context.Context, u *ent.User, status access.Status) ([]*ent.Adventure, error) {
	return u.QueryAccess().Where(access.StatusEQ(status)).QueryAdventures().All(ctx)
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

func (db *DBClient) GetAdventureAccessForUser(ctx context.Context, a *ent.Adventure, u *ent.User) (*ent.Access, error) {
	return u.QueryAccess().Where(access.And(access.UserID(u.ID), access.AdventureID(a.ID))).Only(ctx)
}

func (db *DBClient) GetUserStatusForAdventure(ctx context.Context, a *ent.Adventure, u *ent.User) (access.Status, error) {
	ac, err := db.GetAdventureAccessForUser(ctx, a, u)
	if err != nil {
		return access.StatusUnpaid, err
	}
	return ac.Status, err
}

func (db *DBClient) SetUserStatusForAdventure(ctx context.Context, a *ent.Adventure, u *ent.User, status access.Status) error {
	ac, err := db.GetAdventureAccessForUser(ctx, a, u)
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
	_, err := db.entclient.Access.Query().Where(access.Code(code)).Only(ctx)
	return err
}
