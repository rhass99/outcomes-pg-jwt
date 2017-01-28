package main

import (
	"github.com/gorilla/mux"
	"github.com/rhass99/outcomes-pg-jwt/api"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/signup", api.SignupUsersp)
	r.HandleFunc("/login", api.LoginUsersp)
	r.HandleFunc("/profile", api.ProfileUserspGet)
	http.ListenAndServe(":8080", r)
}
