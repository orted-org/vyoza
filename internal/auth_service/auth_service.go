package authservice

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	kvstore "github.com/orted-org/vyoza/pkg/kv_store"
	httperror "github.com/orted-org/vyoza/pkg/http_error"
)
type Admin struct {
	Name string;
	Email string;
	Password string
}

var admins []Admin = []Admin{
	{Name: "admin1", Password: "pass1", Email: "admin1@gmail.com"},
}
//session Age in seconds
const SessionAge = 24*60*60

type AuthService struct {
	inMemKVStore kvstore.KVStore
}

func New(kv kvstore.KVStore) *AuthService {
	return &AuthService{
		inMemKVStore: kv,
	}
}

type LoginArgs struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


type SessionData struct {
	AdminData Admin
	Expires time.Time
}

type Session struct {
	Id string
	Data SessionData
}

func (authService *AuthService) PerformLogin(sessionId string, arg LoginArgs) (Session, error) {
	var session Session
	var err error
	if sessionId=="" {
		session, err = performNormalLogin(authService, arg)
		if err!=nil {
			return session, err
		}
		return session, nil
	}

	session, err = getSession(authService, sessionId)
	if err != nil {
		session, err = performNormalLogin(authService, arg)
		if err!=nil {
			return session, err
		}
		return session, nil
	}
	return session, nil
}

func performNormalLogin(authService *AuthService,arg LoginArgs) (Session, error){
	var session Session;
	admin, err := verifyCredentials(arg)
	if err!=nil {
		return session, err
	}

	session = Session{
		Id: uuid.New().String(), 
		Data: SessionData {
			AdminData: admin, Expires: time.Now().UTC().Add(SessionAge * time.Second),
		},
	}
	
	//Set the session inMemKVStore
	mData, mErr := json.Marshal(session.Data)
	if mErr!=nil {
		return session, &httperror.CError{Status: 500, Message: "internal server error"}
	}
	err = authService.inMemKVStore.Set(session.Id, string(mData))
	if err!=nil {
		return session, &httperror.CError{Status: 500, Message: "internal server error"}
	}

	return session, nil
}

func (authService *AuthService) PerformCheckAllowance(sessionId string) (Session, error) {
	session, err := getSession(authService, sessionId)
	if err!=nil {
		return session, err
	}
	return session, nil
}

func (authService *AuthService) PerformLogout(sessionId string) {
	authService.inMemKVStore.Delete(sessionId)
}

func verifyCredentials(cred LoginArgs) (Admin, error){
	var admin Admin;
	for i := 0; i < len(admins); i++ {
		if cred.Email == admins[i].Email&& cred.Password==admins[i].Password {
			admin = admins[i]
			return admin, nil
		}
	}

	return admin, &httperror.CError{Status: 401, Message: "Incorrect username or Password"}
}

func getSession(authService *AuthService, sessionId string) (Session, error) {
	var data SessionData;
	var session Session
	str, err := authService.inMemKVStore.Get(sessionId)

	if err!=nil {
		return session, &httperror.CError{Status: 401, Message: "Unauthorized"}
	}

	err = json.Unmarshal([]byte(str), &data)
	if err!=nil {
		return session, &httperror.CError{Status: 500, Message: "Internal Server Error"}
	}
	session = Session{Id: sessionId, Data: data}

	
	if time.Now().UTC().Sub(session.Data.Expires) > 0{
		//session expired, deleting from inMem
		authService.inMemKVStore.Delete(sessionId)
		return session, &httperror.CError{Status: 500, Message: "Unauthorized"}
	}

	return session, nil
}
