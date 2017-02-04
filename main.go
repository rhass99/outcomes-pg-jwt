package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/rhass99/outcomes-pg-jwt/api"
	"net/http"
	"os"
)

func loggingHandler(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}

func main() {
	r := mux.NewRouter()
	commonHandlers := alice.New(loggingHandler)
	r.Handle("/signup", commonHandlers.ThenFunc(api.SignupUsersp))
	r.Handle("/login", commonHandlers.ThenFunc(api.LoginUsersp))
	r.Handle("/profile", commonHandlers.ThenFunc(api.ProfileUserspGet))
	http.ListenAndServe(":8080", r)
}
