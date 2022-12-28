package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/src/controllers"
	"database/sql"
	_ "github.com/lib/pq"
)

func indexPage(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, hannu")
	fmt.Println(req)
}

func main() {
	connectDatabase()
	http.HandleFunc("/hi", indexPage)
	// http.Handle("/", http.FileServer(http.Dir("./static"))) serve static files
	http.HandleFunc("/api/tasks/", controllers.TaskHandler)
	http.HandleFunc("/api/ping/", controllers.Ping)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3001"
	}
	fmt.Println("Server started on port", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func connectDatabase() {
	// either put these in dockerfile or .env
	const (
		host	="db"
		port	= 5432
		user	= "admin"
		password= "admin"
		dbname	= "admin"
	)

	connStr := fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable",
	host,port,user,password,dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to db: " + dbname)

}
