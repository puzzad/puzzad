package puzzad

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/google/uuid"
	"github.com/greboid/puzzad/ent"
	"github.com/greboid/puzzad/ent/user"
	"github.com/rs/zerolog/log"
)

var (
	ErrBadUsernameOrPassword = errors.New("invalid username or password")
	ErrAccountDisabled       = errors.New("account is disabled")
	ErrAccountUnverified     = errors.New("account is unverified")
	ErrBadEmailAddress       = errors.New("invalid e-mail address")
)

type UserDatabase interface {
	HasAdmins(ctx context.Context) (bool, error)

	CreateUser(ctx context.Context, email, hash string) (*ent.User, error)
	GetUser(ctx context.Context, email string) (*ent.User, error)
	SetPassword(ctx context.Context, u *ent.User, hash string) error
	SetAdmin(ctx context.Context, u *ent.User, admin bool) error
	SetStatus(ctx context.Context, u *ent.User, status user.Status) error

	VerifyVerificationCode(ctx context.Context, code string) (*ent.User, error)
	InvalidateVerificationCode(ctx context.Context, u *ent.User) error

	GeneratePasswordResetCode(ctx context.Context, u *ent.User) (string, error)
	InvalidatePasswordResetCode(ctx context.Context, u *ent.User) error
}

type UserMailer interface {
	SendPasswordResetLink(ctx context.Context, email string, code string) error
}

type UserManager struct {
	db UserDatabase
	m  UserMailer
}

func NewUserManager(db UserDatabase, mailer UserMailer) *UserManager {
	return &UserManager{
		db: db,
		m:  mailer,
	}
}

// EnsureAdminExists checks if there is at least one admin user, and if not creates a new one with the given
// email and password combination.
func (um *UserManager) EnsureAdminExists(ctx context.Context, defaultEmail, defaultPassword string) error {
	exists, err := um.db.HasAdmins(ctx)
	if err != nil {
		return fmt.Errorf("failed to query admin users: %w", err)
	}

	if exists {
		log.Debug().Msg("There are existing admin users, not creating default.")
		return nil
	}

	if err := validateEmailAddress(defaultEmail); err != nil {
		return fmt.Errorf("no admin user exists and invalid admin e-mail address supplied: %s (%w)", defaultEmail, err)
	}

	// TODO: Use proper password validation logic when we have it
	if len(defaultPassword) == 0 {
		return fmt.Errorf("no admin user exists and no default admin password supplied")
	}

	hash, err := GetHash(defaultPassword)
	if err != nil {
		return fmt.Errorf("failed to create hash: %w", err)
	}

	u, err := um.db.CreateUser(ctx, defaultEmail, hash)
	if err != nil {
		return fmt.Errorf("failed to create admin user: %w", err)
	}

	if err := um.db.SetAdmin(ctx, u, true); err != nil {
		return fmt.Errorf("failed to set admin status: %w", err)
	}

	if err := um.db.SetStatus(ctx, u, user.StatusVerified); err != nil {
		return fmt.Errorf("failed to set admin account to verified: %w", err)
	}

	log.Info().Msgf("Created new admin user: %s", defaultEmail)
	return nil
}

// Authenticate attempts to validate the given email and password combination, returning the corresponding user
// object if successful.
//
// Expected error values for user-facing problems: ErrBadUsernameOrPassword, ErrAccountDisabled, ErrAccountUnverified.
func (um *UserManager) Authenticate(ctx context.Context, email, password string) (*ent.User, error) {
	u, err := um.db.GetUser(ctx, email)
	if err != nil {
		log.Info().Err(err).Msgf("Login failed for user '%s': could not retrieve user", email)
		return nil, ErrBadUsernameOrPassword
	}

	if ok, err := CheckHash(password, u.Passhash); !ok {
		log.Info().Err(err).Msgf("Login failed for user '%s': password check failed", email)
		return nil, ErrBadUsernameOrPassword
	}

	switch u.Status {
	case user.StatusUnverified:
		log.Info().Msgf("Login failed for user '%s': account is unverified", email)
		return nil, ErrAccountUnverified
	case user.StatusDisabled:
		log.Info().Msgf("Login failed for user '%s': account is disabled", email)
		return nil, ErrAccountDisabled
	case user.StatusVerified:
		log.Debug().Msgf("Login successful for user '%s'", email)
		return u, nil
	default:
		log.Error().Msgf("Login failed for user '%s', unknown account status: %s", email, u.Status)
		return nil, fmt.Errorf("unknown account status")
	}
}

// StartPasswordReset attempts to initiate a password reset for the given e-mail address.
//
// It is not possible to tell from the return value whether the given account exists, and user messaging should
// reflect that (e.g. "If you have an account, we have e-mailed you"). Any returned error is an internal error and
// should not be displayed to the user.
func (um *UserManager) StartPasswordReset(ctx context.Context, email string) error {
	u, err := um.db.GetUser(ctx, email)
	if err != nil {
		// No such user - don't return an error to mitigate user enumeration attacks
		log.Info().Err(err).Msgf("Password reset failed for user '%s': could not retrieve user", email)
		return nil
	}

	// TODO: Maybe check when the last reset was and disallow if it's too recent?

	code, err := um.db.GeneratePasswordResetCode(ctx, u)
	if err != nil {
		log.Error().Err(err).Msgf("Password reset failed for user '%s': could not generate code", email)
		return err
	}

	err = um.m.SendPasswordResetLink(ctx, u.Email, code)
	if err != nil {
		log.Error().Err(err).Msgf("Password reset failed for user '%s': could not send message", email)
		return err
	}

	return nil
}

// FinishPasswordReset attempts to complete a password reset previously started using StartPasswordReset. If the user's
// reset code is valid, the password will be set to the one provided.
func (um *UserManager) FinishPasswordReset(ctx context.Context, email, code, password string) (bool, error) {
	u, err := um.db.GetUser(ctx, email)
	if err != nil {
		log.Info().Err(err).Msgf("Password reset failed for user '%s': could not retrieve user", email)
		return false, nil
	}

	if len(u.ResetCode) == 0 || u.ResetCode != code {
		log.Info().Msgf("Password reset failed for user '%s': invalid code", email)
		return false, nil
	}

	if u.ResetExpiry.Before(time.Now()) {
		log.Info().Msgf("Password reset failed for user '%s': code expired", email)
		return false, nil
	}

	// TODO: Check password validity here once we have such rules

	err = um.db.InvalidatePasswordResetCode(ctx, u)
	if err != nil {
		log.Error().Err(err).Msgf("Password reset failed for user '%s': could not invalidate reset code", email)
		return false, err
	}

	hash, err := GetHash(password)
	if err != nil {
		log.Error().Err(err).Msgf("Password reset failed for user '%s': could not create hash", email)
		return false, err
	}

	err = um.db.SetPassword(ctx, u, hash)
	if err != nil {
		log.Error().Err(err).Msgf("Password reset failed for user '%s': could not save password", email)
		return false, err
	}

	return true, nil
}

func (um *UserManager) CreateUser(ctx context.Context, email string) error {
	hash, err := GetHash(uuid.New().String())
	if err != nil {
		return err
	}
	u, err := um.db.CreateUser(ctx, email, hash)
	if err != nil {
		return err
	}
	return um.m.SendPasswordResetLink(ctx, u.Email, u.VerifyCode)
}

func (um *UserManager) CompleteVerification(ctx context.Context, code string) (*ent.User, error) {
	// TODO Should probably split this up to handle expired verify codes better
	u, err := um.db.VerifyVerificationCode(ctx, code)
	if err != nil {
		return nil, err
	}
	err = um.db.SetStatus(ctx, u, user.StatusVerified)
	if err != nil {
		return nil, err
	}
	err = um.db.InvalidateVerificationCode(ctx, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func validateEmailAddress(email string) error {
	parsed, err := mail.ParseAddress(email)

	// We outright can't parse it
	if err != nil {
		return ErrBadEmailAddress
	}

	// The input must have contained some other gunk like a random name before the address.
	if parsed.Address != email {
		return ErrBadEmailAddress
	}

	return nil
}
