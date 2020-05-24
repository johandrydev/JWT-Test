package main

import (
	c "github/takeodev/JWT-test/controllers"
	t "github/takeodev/JWT-test/token"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", c.Login).Methods("POST")
	r.HandleFunc("/test-token", t.ValidateToken(c.TestToken)).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Init on port :8000")
	log.Fatal(srv.ListenAndServe())
}
