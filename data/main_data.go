package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"encoding/json"
	"io"
	"errors"

	"github.com/go-playground/validator"
)

var connStr string = "user=app dbname=app password=learningG0 host=localhost port=7654 sslmode=disable"

type Register struct {
	Id 			int 	`json:id, omitempty`
	Firstname	string 	`json:"firstname" validate:"required"`
	Lastname	string 	`json:"lastname" validate:"required"`
	Email 		string	`json:"email" validate:"required,email"`
	Password 	string	`json:"password" validate:"required,min=8"`
	Createdon 	string 	`json:createdon, omitempty`
}

func (re *Register) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(re)
}

func (re *Register) Validate() error {
	validate := validator.New()

	return validate.Struct(re)
}

func (re *Register) PostRegister() error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStatement := `
		SELECT 1 FROM users WHERE email = $1
	`

	var x int
	row := db.QueryRow(sqlStatement, &re.Email)
	err = row.Scan(&x)
	if x == 1 {
		return errors.New("User exists")
	    
	}

	sqlStatement = `
		INSERT INTO users (id, firstname, lastname, email, password, createdon)
		VALUES ($1, $2, $3, $4, $5, now())`
	_, err = db.Exec(sqlStatement, 1, &re.Firstname, &re.Lastname, &re.Email, &re.Password)
	if err != nil {
	  return err
	}
	return nil
}




type Login struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (lo *Login) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(lo)
}

func (lo *Login) Validate() error {
	validate := validator.New()
	return validate.Struct(lo)
}

func (lo *Login) PostLogin() (*Register, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user Register
	sqlStatement := `
		SELECT * FROM users
		WHERE email = ($1)`
	row := db.QueryRow(sqlStatement, &lo.Email)
	err = row.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Email, &user.Password, &user.Createdon)
	if err != nil && err == sql.ErrNoRows {
		return nil, err
	}
	if lo.Password != user.Password {
		return nil, errors.New("Wrong password")
	} 
	
	return &user, nil
}