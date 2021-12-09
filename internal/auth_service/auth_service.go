package authservice

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
	kvstore "github.com/orted-org/vyoza/pkg/kv_store"
)

// session age in seconds
const SessionAge = 24 * 60 * 60

var ErrInternalServerError = errors.New("internal server error")
var ErrUnauthorized = errors.New("unauthorized")

type LoginArgs struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Session struct {
	ID      string
	Expires time.Time
}

type AuthService struct {
	kv kvstore.KVStore
}

func New(kv kvstore.KVStore) *AuthService {
	return &AuthService{
		kv: kv,
	}
}

func (as *AuthService) PerformLogin(ctx context.Context, arg LoginArgs) (Session, error) {
	var session Session

	// getting the username and password for admin from the env
	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")

	// setting the admin's username to default 'admin' if nothing provided in env
	if username == "" {
		username = "admin"
	}

	// setting the admin's password to default 'admin' if nothing provided in env
	if password == "" {
		password = "admin"
	}

	// checking if the username and password match
	if !(username == arg.Username && password == arg.Password) {
		return Session{}, ErrUnauthorized
	}

	// creating a session
	session = Session{
		ID:      uuid.New().String(),
		Expires: time.Now().UTC().Add(SessionAge * time.Second),
	}

	// session to store to store in kv store
	mData, mErr := json.Marshal(session)
	if mErr != nil {
		return session, ErrInternalServerError
	}

	// storing the session in kv store
	err := as.kv.Set(session.ID, string(mData))
	if err != nil {
		return session, ErrInternalServerError
	}
	return session, nil
}

func (as *AuthService) IfLogin(sessionId string) (Session, error) {
	if sessionId == "" {
		return Session{}, ErrUnauthorized
	}

	var session Session
	str, err := as.kv.Get(sessionId)

	if err != nil {
		return session, ErrUnauthorized
	}

	err = json.Unmarshal([]byte(str), &session)
	if err != nil {
		return session, ErrInternalServerError
	}

	if time.Now().UTC().Sub(session.Expires) > 0 {
		//session expired, deleting from kv store
		as.kv.Delete(sessionId)
		return session, ErrUnauthorized
	}
	return session, nil
}

func (as *AuthService) PerformLogout(sessionId string) error {
	return as.kv.Delete(sessionId)
}
