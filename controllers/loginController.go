package controllers

import (
	"encoding/json"
	"fmt"
	"github/takeodev/JWT-test/token"
	"log"
	"net/http"
)

// Login allows access to the system
func Login(w http.ResponseWriter, r *http.Request) {
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		ErrResponse(w, err)
		return
	}

	if user.Username == "admin" && user.Password == "admin" {
		tokens, err := token.GenerateTokenPair(user)
		if err != nil {
			ErrResponse(w, err)
			return
		}
		j, err := json.Marshal(tokens)
		if err != nil {
			log.Fatalf("Error converting token to json: %s", err)
			ErrResponse(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		m := Message{
			Message: "Invalid username or password",
		}
		j, err := json.Marshal(m)
		if err != nil {
			log.Fatalf("Error converting message: %s", err)
			ErrResponse(w, err)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(j)
	}
}

// TestToken verify functions with a validate token middleware
func TestToken(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success Test"))
}
