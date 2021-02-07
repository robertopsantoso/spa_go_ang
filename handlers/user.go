package handlers

import (
	"data" 

	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Users struct {
	l *log.Logger
}

func NewUsers(l *log.Logger) *Users {
	return &Users{l}
}

func (u *Users) GetUsers(w http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle GET Users")
	

	// fetch  User data
	lu := data.GetUsers()


	err := lu.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to Marshal Users", http.StatusInternalServerError)
		return
	}

}

func (u *Users) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	u.l.Println("Handle GET User by ID")

	// fetch  User data
	iu, err := data.GetUserById(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}
	

	err = iu.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to Marshal Users", http.StatusInternalServerError)
		return
	}
}