package main

import (
	"log"
	"net/http"
)

func (app *App) handleLogin(rw http.ResponseWriter, r *http.Request){
 	//TODO :-  Perform Login
	rw.Write([]byte("SuccessFully Logged in"))
}

func (app *App) handleLogout(rw http.ResponseWriter, r *http.Request){
	//TODO := Perform Logout
	rw.Write([]byte("SuccessFully Logged Out"))
}

//CheckAllowance MiddleWare
func (app *App) handleCheckAllowance(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		//TODO := CheckAllowance
		log.Println("Checking Allowances")
		next.ServeHTTP(rw, r)
	})
}
