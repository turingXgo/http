package main

import (
	"log"
	"net/http"

	"github.com/turingXgo/http/app"
	"github.com/turingXgo/http/db"
	"github.com/turingXgo/http/postgres"
)

func main() {
	conn, err := db.NewConnection("demo", "postgres", "data")
	if err != nil {
		log.Fatal(err)
	}
	repo := postgres.New(conn)
	application := app.NewApplication(repo)
	http.HandleFunc("/user", application.GetUserByName)
	http.HandleFunc("/", application.CreateUser)
	http.ListenAndServe(":8081", nil)
}
