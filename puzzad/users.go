package puzzad

import (
	"context"
	"errors"
	"fmt"

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
}

type UserManager struct {
	db UserDatabase
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
