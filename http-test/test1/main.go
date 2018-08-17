package main

import "fmt"

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header()
	fmt.Fprintln(w, "hello world")
}

func Getlist(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println(r.RequestURI)
	fmt.Println(query)

	time.Sleep(time.Duration(2)*time.Second)

	fmt.Fprintln(w, "12,3,3,,4,5")
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/getlist", Getlist)

	log.Fatal(http.ListenAndServe("127.0.0.1:8081", router))
}
