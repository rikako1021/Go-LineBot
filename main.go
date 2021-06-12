package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("http://localhost:8080 で起動中")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, t *http.Request) {
	msg := "Hello World!!!!"
	fmt.Fprintf(w, msg)
}
