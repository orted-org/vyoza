package authservice

import (
	"errors"

	"github.com/google/uuid"
	kvstore "github.com/orted-org/vyoza/pkg/kv_store"
)
type Admin struct {
	Name string;
	Email string;
	Password string
}

var admins []Admin = []Admin{
	Admin{Name: "admin1", Password: "pass1", Email: "admin1@gmail.com"},
} 
type AuthService struct {
	inMemKVStore *kvstore.InMemKVStore
}

func New(kv *kvstore.InMemKVStore) *AuthService {
	return &AuthService{
		inMemKVStore: kv,
	}
}

type LoginArgs struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Session struct {
	SessionId string
	SessionData Admin
}

func (authService *AuthService) PerformLogin(arg LoginArgs) (Session, error) {
	var session Session;
	admin, err := verifyCredentials(arg)
	if err!=nil {
		return session, err
	}

	
	sessionId := uuid.New().String()
	//Set the session inMemKVStore
	err = authService.inMemKVStore.Set(sessionId, admin)
	if err!=nil {
		return session, errors.New("500")
	}

	session = Session{SessionId: sessionId, SessionData: admin}
	return session, nil
}


func verifyCredentials(cred LoginArgs) (Admin, error){
	var admin Admin;
	for i := 0; i < len(admins); i++ {
		if cred.Email == admins[i].Email&& cred.Password==admins[i].Password {
			admin = admins[i]
			return admin, nil
		}
	}

	return admin, errors.New("401")
}

func (authService *AuthService) PerformLogout(){
	
}

func (authService *AuthService) PerformCheckAllowance(){
	
}

