package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go!")
}

func main() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Listening on port 9003")
	log.Fatal(http.ListenAndServe(":9003", nil))
}
