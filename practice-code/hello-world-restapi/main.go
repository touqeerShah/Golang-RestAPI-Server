package main

import (
	"log"
	"fmt"
	"net/http"
)
func helloworld(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "hello world\n");
}
func main()  {
	http.HandleFunc("/",helloworld);
	fmt.Println("Server Started !");

	log.Fatal(http.ListenAndServe(":9003",nil))
}