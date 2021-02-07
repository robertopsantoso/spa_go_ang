package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
    "time"
    "context"
    "os"
    "os/signal"
    "handlers"

    "github.com/gorilla/mux"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  // parse arguments, you have to call this by yourself
    fmt.Println(r.Form)  // print form information in server side
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello rps!") // send data to client side
}

func main() {
    // handler from package handlers
    l := log.New(os.Stdout, "rps-api ", log.LstdFlags)
    //hh := handlers.NewHello(l)
    uh := handlers.NewUsers(l)

    // custom ServeMux
    /*
    sm := http.NewServeMux()
    sm.Handle("/", hh)
    sm.Handle("/user", uh)
    */
    sm := mux.NewRouter()

    getRouter := sm.Methods(http.MethodGet).Subrouter()
    getRouter.HandleFunc("/users", uh.GetUsers)
    getRouter.HandleFunc("/users/{id:[0-9]+}", uh.GetUserById)

    // custom && configured Server
    s := http.Server{
        Addr: ":8080",
        Handler: sm,
        IdleTimeout: 120*time.Second,
        ReadTimeout: 1*time.Second,
        WriteTimeout: 1*time.Second,
    }

    // implement gracefull shutdown!
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
    s.Shutdown(tc);
}