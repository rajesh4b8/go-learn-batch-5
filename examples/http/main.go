package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal("Something wrong :: ", err.Error())
		os.Exit(1)
	}

	data := make([]byte, 2048)
	n, err := resp.Body.Read(data)
	if err != nil {
		log.Fatal("Something wrong :: ", err.Error())
		os.Exit(1)
	}

	fmt.Println(n, len(data))
	fmt.Println(resp.Header["Content-Type"])
}
