package main

import (
	"context"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	authservice "github.com/orted-org/vyoza/internal/auth_service"
	httperror "github.com/orted-org/vyoza/pkg/http_error"
)



func (app *App) handleLogin(rw http.ResponseWriter, r *http.Request){
	var incBody authservice.LoginArgs;
	err :=  getBody(r, &incBody)

	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
		return
	}

	err = validateLoginInp(r.Context(), &incBody)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
		return
	}
	sessionId, _ := getSessionId(r);



	session, loginErr := app.authService.PerformLogin(sessionId,incBody)

	if loginErr!=nil {
		xs , _ :=  loginErr.(*httperror.CError)
		sendErrorResponse(rw, xs.Status, nil, xs.Message)
		return
	}

	//set the cookie 
	http.SetCookie(rw, &http.Cookie{
		Name:    "_LOC_ID",
		Value:   session.Id,
		Expires: time.Now().UTC().Add(authservice.SessionAge * time.Second),
		HttpOnly: true,
	})

	sendResponse(rw, http.StatusOK, session.Data, "Logged in successfully")
}

func (app *App) handleLogout(rw http.ResponseWriter, r *http.Request){

	session, ok := r.Context().Value("session").(authservice.Session)
	if !ok {
		sendErrorResponse(rw, 500, nil, "internal server error")
	}

	app.authService.PerformLogout(session.Id)

	http.SetCookie(rw, &http.Cookie{
		Name:    "_LOC_ID",
		MaxAge: -1,
		HttpOnly: true,
	})
	sendResponse(rw, 200, nil, "Successfully logged out")
}

//CheckAllowance MiddleWare
func (app *App) handleCheckAllowance(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		
		sessionId, err := getSessionId(r);
		if err!=nil {
			//No cookie found
			sendErrorResponse(rw, http.StatusUnauthorized, nil, "Unauthorized")
			return
		}
		session, CAerr :=  app.authService.PerformCheckAllowance(sessionId)

		if CAerr != nil {
			xs , _ :=  CAerr.(*httperror.CError)
			sendErrorResponse(rw, xs.Status, nil, xs.Message)
			return
		}

		//putting sessionData in request context
		newCtx := context.WithValue(r.Context(), "session", session)

		next.ServeHTTP(rw, r.WithContext(newCtx))
	})
}


func validateLoginInp(ctx context.Context, i *authservice.LoginArgs) error {
	return validation.ValidateStructWithContext(ctx, i,

		validation.Field(&i.Email, validation.Required),
		validation.Field(&i.Password, validation.Required),
	)
}