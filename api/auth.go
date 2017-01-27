package api

import (
	"database/sql"
	//"fmt"
	"github.com/gorilla/schema"
	_ "github.com/lib/pq"
	"github.com/rhass99/outcomes/storage"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func EncryptPass(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	return string(hashedBytes), nil
}
func CompareEncryptPass(password string, hashedpass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpass), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func AuthUsersp(b *storage.Usersp, db *sql.DB) string {
	a, err := storage.GetUserspAuth(b.Email, db)
	switch {
	case err == sql.ErrNoRows:
		return "userdoesntexist"
	case err != nil:
		return "servererror"
	case err == nil:
		if !CompareEncryptPass(b.Password, a.PasswordHash) {
			return "wrongpassword"
		}
		return "userauthenticated"
	}
	return "servererror"
}

func ProcessUserspForm(r *http.Request) *storage.Usersp {
	// Define a new Applicant
	var a storage.Usersp
	// Use Gorilla schema to get from data
	err := r.ParseForm()
	if err != nil {
		log.Println("Error with form")
	}
	// Decode form data and place it in the object
	decoder := schema.NewDecoder()
	err = decoder.Decode(&a, r.PostForm)
	return &a
}

func StoreUsersp(a *storage.Usersp, db *sql.DB) error {
	a.Pid = RandId(40)
	var err error
	a.PasswordHash, err = EncryptPass(a.Password)
	if err != nil || a.PasswordHash == "" {
		//How to handle this error
		log.Println("Cannot encrypt password")
		return err
	}
	err = a.CreateUsersp(db)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
