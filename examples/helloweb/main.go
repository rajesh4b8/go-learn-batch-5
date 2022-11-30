package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func helloHandler(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("<h1>Hello, web!!!</h1>"))
}
