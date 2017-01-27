//psql -U rami -f setup.sql -d outcomes

package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "rami"
	password = "ramihassanein"
	dbname   = "outcomes"
)

type Usersp struct {
	Id           int
	Pid          string `schema:"pid"`
	Firstname    string `schema:"firstname"`
	Lastname     string `schema:"lastname"`
	Filenumber   string `schema:"filenumber"`
	Email        string `schema:"email"`
	Password     string `schema:"password"`
	Password2    string `schema:"password2"`
	PasswordHash string
}

func DBConnect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (a *Usersp) CreateUsersp(db *sql.DB) error {
	var qemail string
	err := db.QueryRow("SELECT email from usersp where email = $1", a.Email).Scan(&qemail)
	switch {
	case err == sql.ErrNoRows:
		statement := "insert into usersp (pid, firstname, lastname, email, filenumber, passwordhash) values ($1, $2, $3, $4, $5, $6) returning id;"
		stmt, err := db.Prepare(statement)
		if err != nil {
			return err
		}
		defer stmt.Close()
		err = stmt.QueryRow(a.Pid, a.Firstname, a.Lastname, a.Email, a.Filenumber, a.PasswordHash).Scan(&a.Id)
		if err != nil {
			return err
		}
		return nil
	case err == nil:
		fmt.Println("Patient Exists")
	default:
		return err
	}
	return nil
}

func GetUserspByID(pid string, db *sql.DB) (Usersp, error) {
	a := Usersp{}
	err := db.QueryRow("SELECT pid, firstname, lastname, email, filenumber FROM usersp WHERE pid = $1", pid).Scan(&a.Pid, &a.Firstname, &a.Lastname, &a.Email, &a.Filenumber)
	return a, err
}

func GetUserspAuth(email string, db *sql.DB) (Usersp, error) {
	a := Usersp{}
	err := db.QueryRow("SELECT email, firstname, lastname, filenumber, passwordhash FROM usersp WHERE email = $1", email).Scan(&a.Email, &a.Firstname, &a.Lastname, &a.Filenumber, &a.PasswordHash)
	return a, err
}
