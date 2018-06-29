package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", GetTUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":12345", router))
}

func GetTUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])
	var user User = User{Id: id, Name: "王一"}
	json.NewEncoder(w).Encode(user)
}

type User struct {
	Id   int
	Name string
}
