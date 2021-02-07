package main

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"context"
	"path/filepath"
	"os"
	"os/signal"

	"handlers"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle SPA")
    // get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
        // if we failed to get the absolute path respond with a 400 bad request
        // and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
    // prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)
    // check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
        // if we got an error (that wasn't that the file doesn't exist) stating the
        // file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    // otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	l := log.New(os.Stdout, "rps-register-api", log.LstdFlags)

	h := handlers.NewHandler(l)

	sm := mux.NewRouter()


	postRouter := sm.PathPrefix("/api").Methods(http.MethodPost, "OPTIONS").Subrouter()
	postRouter.HandleFunc("/register", h.RegisterUser)
	postRouter.HandleFunc("/login", h.LoginUser)
	// postRouter.Use(h.MiddlewareValidateRegister)

	getRouter := sm.PathPrefix("/api").Methods(http.MethodGet, "OPTIONS").Subrouter()
	getRouter.HandleFunc("/auth", h.AuthUser)
	getRouter.HandleFunc("/logout", h.LogoutUser)
	getRouter.HandleFunc("/account", h.GetUser)

	spa := spaHandler{staticPath: "../views", indexPath: "index.html"}
	sm.PathPrefix("/").Handler(spa)

	s := http.Server {
		Addr : ":8081",
		Handler : sm,
		IdleTimeout : 120*time.Second,
		ReadTimeout : 1*time.Second,
		WriteTimeout : 1* time.Second,
	}

	go func () {
		err := s.ListenAndServe();
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	log.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}