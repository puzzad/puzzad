package database

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/greboid/puzzad/ent/user"

	"github.com/greboid/puzzad/ent"
	_ "github.com/mattn/go-sqlite3"
)

func (db *DBClient) CreateUser(ctx context.Context, email, hash string) (*ent.User, error) {
	return db.entclient.User.
		Create().
		SetEmail(email).
		SetPasshash(hash).
		Save(ctx)
}

func (db *DBClient) GetUser(ctx context.Context, email string) (*ent.User, error) {
	return db.entclient.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)
}

func (db *DBClient) GenerateVerificationCode(ctx context.Context, email string) (string, error) {
	u, err := db.GetUser(ctx, email)
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

func (db *DBClient) GeneratePasswordResetCode(ctx context.Context, u *ent.User) (string, error) {
	code := uuid.New().String()

	_, err := db.entclient.User.
		UpdateOne(u).
		SetResetCode(code).
		SetResetExpiry(time.Now().Add(time.Hour)).
		Save(ctx)
	if err != nil {
		return "", err
	}

	return code, nil
}

func (db *DBClient) InvalidatePasswordResetCode(ctx context.Context, u *ent.User) error {
	_, err := db.entclient.User.
		UpdateOne(u).
		SetResetCode("").
		Save(ctx)
	return err
}

func (db *DBClient) SetPassword(ctx context.Context, u *ent.User, hash string) error {
	_, err := db.entclient.User.
		UpdateOne(u).
		SetPasshash(hash).
		Save(ctx)
	return err
}

func (db *DBClient) UpdateUserStatus(ctx context.Context, u *ent.User, newStatus user.Status) error {
	_, err := u.Update().SetStatus(newStatus).Save(ctx)
	return err
}

func (db *DBClient) HasAdmins(ctx context.Context) (bool, error) {
	return db.entclient.User.Query().Where(user.Admin(true)).Exist(ctx)
}

func (db *DBClient) SetAdmin(ctx context.Context, u *ent.User, admin bool) error {
	_, err := db.entclient.User.UpdateOne(u).
		SetAdmin(admin).
		Save(ctx)
	return err
}

func (db *DBClient) SetStatus(ctx context.Context, u *ent.User, status user.Status) error {
	_, err := db.entclient.User.UpdateOne(u).
		SetStatus(status).
		Save(ctx)
	return err
}
