package data

import (
	"encoding/json"
	"io"
	"fmt"
	"log"

	"orm"
)

type User struct {
	ID 			int 	`json:"id"`
	LastName 	string 	`json:"lastName" validate:"required"`
	FirstName	string 	`json:"firstName" validate:"required"`
	Email 		string 	`json:"email" validate:"required"`
	CreatedOn	string 	`json:"-"`
}

type Users []*User

func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func GetUsers () Users {
	rows := orm.GetUsers()
	defer rows.Close()
	
	var users Users
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.LastName, &user.FirstName, &user.Email, &user.CreatedOn)
		if err != nil {
			log.Println(err)
		}

		users = append(users, &user)
	}

	return users
}

var ErrorUserNotFound = fmt.Errorf("User not found!")

func GetUserById (id int) (Users, error) {
	row := orm.GetUserById(id)

	var users Users
	var user User
	err := row.Scan(&user.ID, &user.LastName, &user.FirstName, &user.Email, &user.CreatedOn)
	switch {
		case len(user.LastName) == 0: {
			log.Printf("No user with ID %d", id)
			return nil, ErrorUserNotFound
		}
		case err != nil: {
			log.Fatal(err)
			return nil, err
		}
		default:
			users = append(users, &user)
	}

	return users, nil
}