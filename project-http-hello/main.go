package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", Greet)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Greet(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	fmt.Fprint(writer, "Hello, ", name)
}
