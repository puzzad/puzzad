package puzzad

import (
	"context"
	"testing"

	"github.com/greboid/puzzad/ent"
	"github.com/greboid/puzzad/ent/user"
	"github.com/stretchr/testify/assert"
)

type FakeUserDatabase struct {
	email string
	user  *ent.User
	err   error
}

func (f *FakeUserDatabase) GetUser(_ context.Context, email string) (*ent.User, error) {
	f.email = email
	return f.user, f.err
}

const testHash = "argon2id$65536$1$4$aJuLQpIeRmFotQzV90bnIMe7e49RMTm5Ve8g89xhUrQ$DiASQQu2g3NneKskie+z9B2lUFSvD7s5ZN/ggPGR8arPPXpV+ZJPr1HZm15Z8LPgGsIpqBcNG8hj3nACLkkZdA"

func TestUserManager_Authenticate_returnsErrorIfUserDoesNotExist(t *testing.T) {
	db := &FakeUserDatabase{
		user: nil,
		err:  &ent.NotFoundError{},
	}

	manager := &UserManager{db: db}

	u, err := manager.Authenticate(context.Background(), "em@il", "password")
	assert.Equal(t, "em@il", db.email, "should try to retrieve the correct user")
	assert.Nil(t, u, "should return nil user")
	assert.Equal(t, err, ErrBadUsernameOrPassword, "should return correct error")
}

func TestUserManager_Authenticate_returnsErrorIfPasswordHashIsInvalid(t *testing.T) {
	db := &FakeUserDatabase{
		user: &ent.User{Passhash: ""},
		err:  nil,
	}

	manager := &UserManager{db: db}

	u, err := manager.Authenticate(context.Background(), "em@il", "password")
	assert.Equal(t, "em@il", db.email, "should try to retrieve the correct user")
	assert.Nil(t, u, "should return nil user")
	assert.Equal(t, err, ErrBadUsernameOrPassword, "should return correct error")
}

func TestUserManager_Authenticate_returnsErrorIfPasswordHashDoesNotMatch(t *testing.T) {
	db := &FakeUserDatabase{
		user: &ent.User{Passhash: testHash},
		err:  nil,
	}

	manager := &UserManager{db: db}

	u, err := manager.Authenticate(context.Background(), "em@il", "password")
	assert.Equal(t, "em@il", db.email, "should try to retrieve the correct user")
	assert.Nil(t, u, "should return nil user")
	assert.Equal(t, err, ErrBadUsernameOrPassword, "should return correct error")
}

func TestUserManager_Authenticate_returnsErrorIfAccountUnverified(t *testing.T) {
	db := &FakeUserDatabase{
		user: &ent.User{Passhash: testHash, Status: user.StatusUnverified},
		err:  nil,
	}

	manager := &UserManager{db: db}

	u, err := manager.Authenticate(context.Background(), "em@il", "test")
	assert.Equal(t, "em@il", db.email, "should try to retrieve the correct user")
	assert.Nil(t, u, "should return nil user")
	assert.Equal(t, err, ErrAccountUnverified, "should return correct error")
}

func TestUserManager_Authenticate_returnsErrorIfAccountDisabled(t *testing.T) {
	db := &FakeUserDatabase{
		user: &ent.User{Passhash: testHash, Status: user.StatusDisabled},
		err:  nil,
	}

	manager := &UserManager{db: db}

	u, err := manager.Authenticate(context.Background(), "em@il", "test")
	assert.Equal(t, "em@il", db.email, "should try to retrieve the correct user")
	assert.Nil(t, u, "should return nil user")
	assert.Equal(t, err, ErrAccountDisabled, "should return correct error")
}

func TestUserManager_Authenticate_returnsErrorIfAccountStatusUnknown(t *testing.T) {
	db := &FakeUserDatabase{
		user: &ent.User{Passhash: testHash, Status: "wtf?"},
		err:  nil,
	}

	manager := &UserManager{db: db}

	u, err := manager.Authenticate(context.Background(), "em@il", "test")
	assert.Equal(t, "em@il", db.email, "should try to retrieve the correct user")
	assert.Nil(t, u, "should return nil user")
	assert.Error(t, err, "should return an error")
}

func TestUserManager_Authenticate_returnsUserOnSuccess(t *testing.T) {
	db := &FakeUserDatabase{
		user: &ent.User{Passhash: testHash, Status: user.StatusVerified},
		err:  nil,
	}

	manager := &UserManager{db: db}

	u, err := manager.Authenticate(context.Background(), "em@il", "test")
	assert.Equal(t, "em@il", db.email, "should try to retrieve the correct user")
	assert.Equal(t, db.user, u, "should return user")
	assert.NoError(t, err, "should not return an error")
}
