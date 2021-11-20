package authservice

import (
	kvstore "github.com/orted-org/vyoza/pkg/kv_store"
)

type AuthService struct {
	inMemKVStore *kvstore.InMemKVStore
}

func New(kv *kvstore.InMemKVStore) *AuthService {
	return &AuthService{
		inMemKVStore: kv,
	}
}

func (authService *AuthService) PerformLogin(){

}

func (authService *AuthService) PerformLogout(){
	
}

func (authService *AuthService) PerformCheckAllowance(){
	
}

