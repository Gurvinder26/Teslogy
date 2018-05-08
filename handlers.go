package main

import (
	"encoding/json"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	userList := Users{User{ID: 1, Firstname: "John", Lastname: "me"}}

	json.NewEncoder(w).Encode(userList)
}
