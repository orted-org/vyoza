package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	authservice "github.com/orted-org/vyoza/internal/auth_service"
)



func (app *App) handleLogin(rw http.ResponseWriter, r *http.Request){
	var incBody authservice.LoginArgs;
	err :=  getBody(r, &incBody)
	fmt.Printf("incBody is %v", incBody)

	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
		return
	}

	err = validateLoginInp(r.Context(), &incBody)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
		return
	}

	session, loginErr := app.authService.PerformLogin(incBody)

	if loginErr!=nil {
		if loginErr.Error()=="401" {
			sendErrorResponse(rw, http.StatusUnauthorized, nil, "incorrect email or password")
			return
		}
		sendErrorResponse(rw, http.StatusInternalServerError, nil, "internal server error")
		return
	}

	//set the cookie 
	http.SetCookie(rw, &http.Cookie{
		Name:    "_LOC_ID",
		Value:   session.SessionId,
		Expires: time.Now().Add(24 * 60 * 60 * time.Second),
	})

	sendResponse(rw, http.StatusOK, session, "Logged in successfully")
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


func validateLoginInp(ctx context.Context, i *authservice.LoginArgs) error {
	return validation.ValidateStructWithContext(ctx, i,

		validation.Field(&i.Email, validation.Required),
		validation.Field(&i.Password, validation.Required),
	)
}