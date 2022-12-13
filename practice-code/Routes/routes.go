package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func postRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post Request\n")
}
func putRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "put Request\n")
}
func deleteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "deleteRequest\n")
}
func getRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get Request \n")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getRequest).Methods("GET")
	r.HandleFunc("/", postRequest).Methods("POST")
	r.HandleFunc("/", putRequest).Methods("PUT")
	r.HandleFunc("/", deleteRequest).Methods("DELETE")
	r.Handle("/", r)
	fmt.Println("Server Started !")

	log.Fatal(http.ListenAndServe(":9003", r))
}
