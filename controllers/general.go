package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// User struct for login test
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Message struct for user message
type Message struct {
	Message string `json:"message"`
}

func ErrResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	errPrint := json.NewEncoder(w).Encode(struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
	if errPrint != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
