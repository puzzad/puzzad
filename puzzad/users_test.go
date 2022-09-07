package puzzad

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"testing"
	"testing/iotest"
	"time"

	"github.com/greboid/puzzad/ent"
	"github.com/greboid/puzzad/ent/user"
	"github.com/stretchr/testify/assert"
)

type FakeUserDatabase struct {
	email           string
	getUserResponse *ent.User
	getUserError    error

	generateCodeResponse string
	generateCodeUser     *ent.User
	generateCodeError    error

	invalidateCodeUser  *ent.User
	invalidateCodeError error

	setPasswordUser  *ent.User
	setPasswordHash  string
	setPasswordError error

	getAdminsUsers []*ent.User
	getAdminsError error

	setAdminUser  *ent.User
	setAdminError error
}

func (f *FakeUserDatabase) GetUser(_ context.Context, email string) (*ent.User, error) {
	f.email = email
	return f.getUserResponse, f.getUserError
}

func (f *FakeUserDatabase) GetAdmins(_ context.Context) ([]*ent.User, error) {
	return f.getAdminsUsers, f.getAdminsError
}

func (f *FakeUserDatabase) GeneratePasswordResetCode(_ context.Context, u *ent.User) (string, error) {
	f.generateCodeUser = u
	return f.generateCodeResponse, f.generateCodeError
}

func (f *FakeUserDatabase) InvalidatePasswordResetCode(_ context.Context, u *ent.User) error {
	f.invalidateCodeUser = u
	return f.invalidateCodeError
}

func (f *FakeUserDatabase) SetPassword(_ context.Context, u *ent.User, hash string) error {
	f.setPasswordUser = u
	f.setPasswordHash = hash
	return f.setPasswordError
}

func (f *FakeUserDatabase) SetAdmin(_ context.Context, u *ent.User, admin bool) error {
	u.Admin = admin
	return f.setPasswordError
}

type FakeUserMailer struct {
	passwordResetError error
	passwordResetEmail string
	passwordResetCode  string
}

func (f *FakeUserMailer) SendPasswordResetLink(_ context.Context, email string, code string) error {
	f.passwordResetEmail = email
	f.passwordResetCode = code
	return f.passwordResetError
}

const testHash = "argon2id$65536$1$4$aJuLQpIeRmFotQzV90bnIMe7e49RMTm5Ve8g89xhUrQ$DiASQQu2g3NneKskie+z9B2lUFSvD7s5ZN/ggPGR8arPPXpV+ZJPr1HZm15Z8LPgGsIpqBcNG8hj3nACLkkZdA"

func TestUserManager_Authenticate_returnsErrorIfUserDoesNotExist(t *testing.T) {
	db := &FakeUserDatabase{
		getUserResponse: nil,
		getUserError:    &ent.NotFoundError{},
	}

	manager := &UserManager{db: db}

	u, err := manager.Authenticate(context.Background(), "em@il", "password")
	assert.Equal(t, "em@il", db.email, "should try to retrieve the correct user")
	assert.Nil(t, u)
	assert.Equal(t, err, ErrBadUsernameOrPassword)
}

func TestUserManager_Authenticate_returnsErrorIfPasswordHashIsInvalid(t *testing.T) {
	db := &FakeUserDatabase{
		getUserResponse: &ent.User{Passhash: ""},
		getUserError:    nil,
	}

	manager := &UserManager{db: db}

	u, err := manager.Authenticate(context.Background(), "em@il", "password")
	assert.Equal(t, "em@il", db.email, "should try to retrieve the correct user")
	assert.Nil(t, u)
	assert.Equal(t, err, ErrBadUsernameOrPassword)
}

func TestUserManager_Authenticate_returnsErrorIfPasswordHashDoesNotMatch(t *testing.T) {
	db := &FakeUserDatabase{
		getUserResponse: &ent.User{Passhash: testHash},
		getUserError:    nil,
	}

	manager := &UserManager{db: db}

	u, err := manager.Authenticate(context.Background(), "em@il", "password")
	assert.Equal(t, "em@il", db.email, "should try to retrieve the correct user")
	assert.Nil(t, u)
	assert.Equal(t, err, ErrBadUsernameOrPassword)
}

func TestUserManager_Authenticate_returnsErrorIfAccountUnverified(t *testing.T) {
	db := &FakeUserDatabase{
		getUserResponse: &ent.User{Passhash: testHash, Status: user.StatusUnverified},
		getUserError:    nil,
	}

	manager := &UserManager{db: db}

	u, err := manager.Authenticate(context.Background(), "em@il", "test")
	assert.Equal(t, "em@il", db.email, "should try to retrieve the correct user")
	assert.Nil(t, u)
	assert.Equal(t, err, ErrAccountUnverified)
}

func TestUserManager_Authenticate_returnsErrorIfAccountDisabled(t *testing.T) {
	db := &FakeUserDatabase{
		getUserResponse: &ent.User{Passhash: testHash, Status: user.StatusDisabled},
		getUserError:    nil,
	}

	manager := &UserManager{db: db}

	u, err := manager.Authenticate(context.Background(), "em@il", "test")
	assert.Equal(t, "em@il", db.email, "should try to retrieve the correct user")
	assert.Nil(t, u)
	assert.Equal(t, err, ErrAccountDisabled)
}

func TestUserManager_Authenticate_returnsErrorIfAccountStatusUnknown(t *testing.T) {
	db := &FakeUserDatabase{
		getUserResponse: &ent.User{Passhash: testHash, Status: "wtf?"},
		getUserError:    nil,
	}

	manager := &UserManager{db: db}

	u, err := manager.Authenticate(context.Background(), "em@il", "test")
	assert.Equal(t, "em@il", db.email, "should try to retrieve the correct user")
	assert.Nil(t, u)
	assert.Error(t, err)
}

func TestUserManager_Authenticate_returnsUserOnSuccess(t *testing.T) {
	db := &FakeUserDatabase{
		getUserResponse: &ent.User{Passhash: testHash, Status: user.StatusVerified},
		getUserError:    nil,
	}

	manager := &UserManager{db: db}

	u, err := manager.Authenticate(context.Background(), "em@il", "test")
	assert.Equal(t, "em@il", db.email, "should try to retrieve the correct user")
	assert.Equal(t, db.getUserResponse, u)
	assert.NoError(t, err)
}

func TestUserManager_StartPasswordReset_returnsNilIfUserDoesNotExist(t *testing.T) {
	db := &FakeUserDatabase{
		getUserResponse: nil,
		getUserError:    &ent.NotFoundError{},
	}

	manager := &UserManager{db: db}

	err := manager.StartPasswordReset(context.Background(), "em@il")
	assert.NoError(t, err)
}

func TestUserManager_StartPasswordReset_returnsErrorIfCodeNotGenerated(t *testing.T) {
	db := &FakeUserDatabase{
		getUserResponse:   &ent.User{},
		generateCodeError: fmt.Errorf("whelp"),
	}

	manager := &UserManager{db: db}

	err := manager.StartPasswordReset(context.Background(), "em@il")
	assert.Equal(t, db.getUserResponse, db.generateCodeUser, "should attempt to generate code for right user")
	assert.Error(t, err)
}

func TestUserManager_StartPasswordReset_returnsErrorIfMailSendingFails(t *testing.T) {
	db := &FakeUserDatabase{
		getUserResponse:      &ent.User{Email: "user.email@example.com"},
		generateCodeResponse: "some-code-here",
	}

	mailer := &FakeUserMailer{
		passwordResetError: fmt.Errorf("smtp is hard, let's go shopping"),
	}

	manager := &UserManager{db: db, m: mailer}

	err := manager.StartPasswordReset(context.Background(), "em@il")
	assert.Equal(t, db.getUserResponse, db.generateCodeUser, "should attempt to generate code for right user")
	assert.Equal(t, db.generateCodeResponse, mailer.passwordResetCode, "should pass correct code to mailer")
	assert.Equal(t, db.getUserResponse.Email, mailer.passwordResetEmail, "should pass correct e-mail to mailer")
	assert.Error(t, err, "should return an error")
}

func TestUserManager_StartPasswordReset_returnsNilOnSuccess(t *testing.T) {
	db := &FakeUserDatabase{
		getUserResponse:      &ent.User{Email: "user.email@example.com"},
		generateCodeResponse: "some-code-here",
	}

	mailer := &FakeUserMailer{}

	manager := &UserManager{db: db, m: mailer}

	err := manager.StartPasswordReset(context.Background(), "em@il")
	assert.Equal(t, db.getUserResponse, db.generateCodeUser, "should attempt to generate code for right user")
	assert.Equal(t, db.generateCodeResponse, mailer.passwordResetCode, "should pass correct code to mailer")
	assert.Equal(t, db.getUserResponse.Email, mailer.passwordResetEmail, "should pass correct e-mail to mailer")
	assert.NoError(t, err, "should not return an error")
}

func TestUserManager_FinishPasswordReset_returnsFalseIfUserNotFound(t *testing.T) {
	db := &FakeUserDatabase{
		getUserError: &ent.NotFoundError{},
	}

	manager := &UserManager{db: db}

	ok, err := manager.FinishPasswordReset(context.Background(), "em@il", "code", "newpassword")
	assert.False(t, ok)
	assert.NoError(t, err)
	assert.Nil(t, db.invalidateCodeUser, "should not try to invalidate code")
}

func TestUserManager_FinishPasswordReset_returnsFalseIfResetCodeIsEmpty(t *testing.T) {
	db := &FakeUserDatabase{
		getUserResponse: &ent.User{},
	}

	manager := &UserManager{db: db}

	ok, err := manager.FinishPasswordReset(context.Background(), "em@il", "", "newpassword")
	assert.False(t, ok)
	assert.NoError(t, err)
	assert.Nil(t, db.invalidateCodeUser, "should not try to invalidate code")
}

func TestUserManager_FinishPasswordReset_returnsFalseIfResetCodeDoesNotMatch(t *testing.T) {
	db := &FakeUserDatabase{
		getUserResponse: &ent.User{ResetCode: "123"},
	}

	manager := &UserManager{db: db}

	ok, err := manager.FinishPasswordReset(context.Background(), "em@il", "456", "newpassword")
	assert.False(t, ok)
	assert.NoError(t, err)
	assert.Nil(t, db.invalidateCodeUser, "should not try to invalidate code")
}

func TestUserManager_FinishPasswordReset_returnsFalseIfResetCodeHasExpired(t *testing.T) {
	db := &FakeUserDatabase{
		getUserResponse: &ent.User{ResetCode: "123", ResetExpiry: time.Now().Add(-time.Hour)},
	}

	manager := &UserManager{db: db}

	ok, err := manager.FinishPasswordReset(context.Background(), "em@il", "123", "newpassword")
	assert.False(t, ok)
	assert.NoError(t, err, "should not return an error")
	assert.Nil(t, db.invalidateCodeUser, "should not try to invalidate code")
}

func TestUserManager_FinishPasswordReset_returnsErrorIfInvalidationFails(t *testing.T) {
	db := &FakeUserDatabase{
		getUserResponse:     &ent.User{ResetCode: "123", ResetExpiry: time.Now().Add(time.Hour)},
		invalidateCodeError: fmt.Errorf("whelp"),
	}

	manager := &UserManager{db: db}

	ok, err := manager.FinishPasswordReset(context.Background(), "em@il", "123", "newpassword")
	assert.False(t, ok)
	assert.Error(t, err)
	assert.Equal(t, db.getUserResponse, db.invalidateCodeUser, "should try to invalidate code")
}

func TestUserManager_FinishPasswordReset_returnsErrorIfHashingFails(t *testing.T) {
	db := &FakeUserDatabase{
		getUserResponse: &ent.User{ResetCode: "123", ResetExpiry: time.Now().Add(time.Hour)},
	}

	manager := &UserManager{db: db}

	// Bodge rand.Reader into providing an error so the hash generation fails, but make sure we restore it for the
	// tests that run afterwards...
	oldReader := rand.Reader
	defer func() {
		rand.Reader = oldReader
	}()
	rand.Reader = iotest.ErrReader(io.EOF)

	ok, err := manager.FinishPasswordReset(context.Background(), "em@il", "123", "newpassword")
	assert.False(t, ok)
	assert.Error(t, err)
	assert.Equal(t, db.getUserResponse, db.invalidateCodeUser, "should try to invalidate code")
}

func TestUserManager_FinishPasswordReset_returnsErrorIfSettingPasswordFails(t *testing.T) {
	db := &FakeUserDatabase{
		getUserResponse:  &ent.User{ResetCode: "123", ResetExpiry: time.Now().Add(time.Hour)},
		setPasswordError: fmt.Errorf("oh noes"),
	}

	manager := &UserManager{db: db}

	ok, err := manager.FinishPasswordReset(context.Background(), "em@il", "123", "newpassword")
	assert.False(t, ok)
	assert.Error(t, err)
	assert.Equal(t, db.getUserResponse, db.invalidateCodeUser, "should try to invalidate code")
	assert.Equal(t, db.getUserResponse, db.setPasswordUser, "should try setting password")

	hashOk, _ := CheckHash("newpassword", db.setPasswordHash)
	assert.True(t, hashOk, "password hash must match the given password")
}

func TestUserManager_FinishPasswordReset_returnsTrueOnSuccess(t *testing.T) {
	db := &FakeUserDatabase{
		getUserResponse: &ent.User{ResetCode: "123", ResetExpiry: time.Now().Add(time.Hour)},
	}

	manager := &UserManager{db: db}

	ok, err := manager.FinishPasswordReset(context.Background(), "em@il", "123", "newpassword")
	assert.True(t, ok)
	assert.NoError(t, err)
	assert.Equal(t, db.getUserResponse, db.invalidateCodeUser, "should try to invalidate code")
	assert.Equal(t, db.getUserResponse, db.setPasswordUser, "should try setting password")

	hashOk, _ := CheckHash("newpassword", db.setPasswordHash)
	assert.True(t, hashOk, "password hash must match the given password")
}

func TestUserManager_GetAdmins_Error(t *testing.T) {
	db := &FakeUserDatabase{
		getAdminsUsers: []*ent.User{},
		getAdminsError: fmt.Errorf("test error"),
	}
	_, err := db.GetAdmins(context.Background())
	assert.True(t, err != nil)
}

func TestUserManager_GetAdmins_Admins(t *testing.T) {
	db := &FakeUserDatabase{
		getAdminsUsers: []*ent.User{{}},
		getAdminsError: nil,
	}
	admins, err := db.GetAdmins(context.Background())
	assert.True(t, len(admins) == 1)
	assert.True(t, err == nil)
}
