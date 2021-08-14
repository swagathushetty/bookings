package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

//middlware func
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

//add CSRF protection to all POST request
func Nosurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "",
		Secure:   app.InProduction, //set true for prod
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

//loads and saves session at every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
