package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexPage(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello, hannu")
}

func handler() {
	http.HandleFunc("/", indexPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Server started on port 8080")
	handler()
}

