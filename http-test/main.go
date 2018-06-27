package main

import "fmt"

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header()
	fmt.Fprintln(w, "hello world")
}

func Getlist(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println(query)
	fmt.Fprintln(w, "12,3,3,,4,5")
}

func main() {

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/getlist", Getlist)

	http.ListenAndServe("127.0.0.1:12345", nil)
}
