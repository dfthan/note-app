package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexPage(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, hannu")
}

func main() {
	http.HandleFunc("/hi", indexPage)
	// http.Handle("/", http.FileServer(http.Dir("./static"))) serve static files


	port := ":3001"
	fmt.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

