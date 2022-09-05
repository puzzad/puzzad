package puzzad

import (
	"context"
	"github.com/google/uuid"
	"github.com/greboid/puzzad/ent/user"
	"time"

	"github.com/greboid/puzzad/ent"
	_ "github.com/mattn/go-sqlite3"
)

func (db *DBClient) CreateUser(ctx context.Context, email string) (*ent.User, error) {
	return db.entclient.User.
		Create().
		SetEmail(email).
		Save(ctx)
}

func (db *DBClient) GetUser(ctx context.Context, email string) (*ent.User, error) {
	return db.entclient.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)
}

func (db *DBClient) GetOrCreateUser(ctx context.Context, email string) (*ent.User, error) {
	u, err := db.GetUser(ctx, email)
	if ent.IsNotFound(err) {
		return db.CreateUser(ctx, email)
	}
	return u, err
}

func (db *DBClient) GenerateVerificationCode(ctx context.Context, email string) (string, error) {
	u, err := db.GetOrCreateUser(ctx, email)
	if err != nil {
		return "", err
	}

	code := uuid.New().String()

	_, err = db.entclient.User.
		UpdateOne(u).
		SetVerifyCode(code).
		SetVerifyExpiry(time.Now().Add(time.Hour)).
		Save(ctx)
	if err != nil {
		return "", err
	}

	return code, nil
}

func (db *DBClient) UpdateUserStatus(ctx context.Context, u *ent.User, newStatus user.Status) error {
	_, err := u.Update().SetStatus(newStatus).Save(ctx)
	return err
}
