package database

import (
	"context"
	"flag"
	"fmt"
	"net/smtp"
	"time"

	"github.com/google/uuid"
	"github.com/greboid/puzzad/ent/user"

	"github.com/greboid/puzzad/ent"
	_ "github.com/mattn/go-sqlite3"
)

var (
	SmtpUser     = flag.String("smtp_user", "", "SMTP Username")
	SmtpPassword = flag.String("smtp_password", "", "SMTP Password")
	SmtpServer   = flag.String("smtp_server", "", "SMTP Server")
	SmtpPort     = flag.Int("smtp_port", 25, "SMTP Server port")
	SmtpFrom     = flag.String("smtp_from", "", "SMTP From address")
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

func (db *DBClient) SendValidationEmail(_ context.Context, e *ent.User) error {
	auth := smtp.PlainAuth("", *SmtpUser, *SmtpPassword, *SmtpServer)
	body := fmt.Sprintf("To: %s\r\nSubject: %s\r\nReply-To: %s\r\nFrom: Puzzad <%s>\r\n\r\nPuzzad verify code: %s\r\n", e.Email, "Puzzad Validation", *SmtpFrom, *SmtpFrom, e.VerifyCode)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", *SmtpServer, *SmtpPort), auth, *SmtpFrom, []string{e.Email}, []byte(body))
	if err != nil {
		return err
	}
	return nil
}

func (db *DBClient) GetAdmins(ctx context.Context) ([]*ent.User, error) {
	users, err := db.entclient.User.Query().Where(user.Admin(true)).All(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (db *DBClient) SetAdmin(ctx context.Context, u *ent.User, b bool) error {
	_, err := db.entclient.User.UpdateOne(u).SetAdmin(b).Save(ctx)
	return err
}