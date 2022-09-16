package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//API routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world from sludgebrain")
	})
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	port := ":8008"
	fmt.Println("Server is running on port" + port)

	log.Fatal(http.ListenAndServe(port, nil))
}
