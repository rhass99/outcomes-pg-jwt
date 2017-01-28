package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rhass99/outcomes-pg-jwt/api"
	_ "net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/signup", api.SignupUsersp)
	r.HandleFunc("/login", api.LoginUsersp)
	r.HandleFunc("/profile", api.ProfileUserspGet)
	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":8080")
}