package handlers

import (
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	// "context"
	"github.com/gorilla/sessions"

	"data"
)
// to store session
var store = sessions.NewCookieStore([]byte("rpscookey"))

type Handler struct {
	l *log.Logger
}

func NewHandler(l *log.Logger) *Handler {
	return &Handler{l}
}

func (h *Handler) auth (w http.ResponseWriter, r *http.Request) (*sessions.Session, bool) {
	s, err := store.Get(r, "rpssesskey")
	if err != nil {
		h.l.Println("[ERROR] get session", err)
		http.Error(w, fmt.Sprintf("Error getting sess: %s", err), http.StatusBadRequest)
		return nil, false
	}

	if s.IsNew {
        s.Options.MaxAge = 600
        s.Options.HttpOnly = false
        s.Options.Secure = false
        log.Println("Create New Session (cookie)")
		return s, false
    }

    log.Println("Use Old Session (old cookie)")
    s.Options.MaxAge = 600
    err = s.Save(r,w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, false
	}
	return s, true
}

func (h *Handler) returnUser (s *sessions.Session, w http.ResponseWriter) {
	type user struct {
		Firstname string `json:"firstname"`
		Lastname string `json:"lastname"`
		Email string `json:"email"`
		CreatedOn string `json:"createdOn"`
	}

	info := user{
			s.Values["fn"].(string),
			s.Values["ln"].(string),
			s.Values["e"].(string),
			s.Values["co"].(string),
		}

	json.NewEncoder(w).Encode(info)
}

func (h *Handler) AuthUser (w http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle GET Auth")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200");

	_, a := h.auth(w, r)
	if a {
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error (w, fmt.Sprintf("Not authorized"), http.StatusBadRequest)
	}
}

func (h *Handler) GetUser (w http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle GET User")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200");

	s, a := h.auth (w, r)
	if a {
		h.returnUser(s, w)
		return
	}

		http.Error(w, "You are not authorized", http.StatusBadRequest)
}

func (h *Handler) RegisterUser (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200");
	h.l.Println("Handle POST Register")
	// reg := r.Context().Value(KeyRegister{}).(data.Register)
	re := data.Register{}

	err := re.FromJSON(r.Body)
	if err != nil {
		h.l.Println("[ERROR] deserializing register", err)
		http.Error(w, "Error reading body req", http.StatusBadRequest)
		return
	}

	err = re.Validate()
	if err != nil {
		h.l.Println("[ERROR] validating body req", err)
		http.Error (w, fmt.Sprintf("Error validating body req: %s", err), http.StatusBadRequest)
		return
	}

	err = re.PostRegister()
	if err != nil {
		h.l.Println("[ERROR] INSERT to DB", err)
		http.Error (w, fmt.Sprintf("Error INSERT DB %s", err), http.StatusBadRequest)
		return	
	}
	h.l.Println("User Registered!")
}

func (h *Handler) LogoutUser (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200");
	h.l.Println("Handle GET Logout")

	s, a := h.auth(w, r)
	if a {
		s.Options.MaxAge = -1
		err := s.Save(r,w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	http.Error(w, "You are not logged in yet", http.StatusBadRequest)
}
func (h *Handler) LoginUser (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200");
	h.l.Println("Handle POST Login")
	// reg := r.Context().Value(KeyRegister{}).(data.Login)
	s, a := h.auth(w, r)

	if a {
		s.Options.MaxAge = -1
		err := s.Save(r,w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Error(w, "You are logged in", http.StatusBadRequest)
		return
	}

	lo := data.Login{}

	err := lo.FromJSON(r.Body)
	if err != nil {
		h.l.Println("[ERROR] deserializing register", err)
		http.Error(w, "Error reading body req", http.StatusBadRequest)
		return
	}

	err = lo.Validate()
	if err != nil {
		h.l.Println("[ERROR] validating body req", err)
		http.Error (w, fmt.Sprintf("Error validating body req: %s", err), http.StatusBadRequest)
		return
	}

	user, err := lo.PostLogin()
	if err != nil {
		h.l.Println("[ERROR] SELECT to DB", err)
		http.Error (w, fmt.Sprintf("Error SELECT DB %s", err), http.StatusBadRequest)
		return	
	}

	s.Values["fn"] = user.Firstname
	s.Values["ln"] = user.Lastname
	s.Values["e"] = user.Email
	s.Values["co"] = user.Createdon

	err = s.Save(r,w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.returnUser(s, w)	
}

/* Middleware
type KeyRegister struct{}

func (h *Handler) MiddlewareValidateRegister(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), KeyRegister{}, reg)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w,r)
	})
}
*/