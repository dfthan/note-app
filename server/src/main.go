package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/src/controllers"
)

func indexPage(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, hannu")
	fmt.Println(req)
}

func main() {
	http.HandleFunc("/hi", indexPage)
	// http.Handle("/", http.FileServer(http.Dir("./static"))) serve static files
	http.HandleFunc("/api/tasks", controllers.GetTasks)
	http.HandleFunc("/api/ping", controllers.Ping)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3001"
	}
	fmt.Println("Server started on port", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
