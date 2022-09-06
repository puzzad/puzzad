package puzzad

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/greboid/puzzad/ent"
	"github.com/greboid/puzzad/ent/user"
	"github.com/rs/zerolog/log"
)

var (
	ErrBadUsernameOrPassword = errors.New("invalid username or password")
	ErrAccountDisabled       = errors.New("account is disabled")
	ErrAccountUnverified     = errors.New("account is unverified")
)

type UserDatabase interface {
	GetUser(ctx context.Context, email string) (*ent.User, error)
	GeneratePasswordResetCode(ctx context.Context, u *ent.User) (string, error)
	InvalidatePasswordResetCode(ctx context.Context, u *ent.User) error
	SetPassword(ctx context.Context, u *ent.User, hash string) error
}

type UserMailer interface {
	SendPasswordResetLink(ctx context.Context, email string, code string) error
}

type UserManager struct {
	db UserDatabase
	m  UserMailer
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
