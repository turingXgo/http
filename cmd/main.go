package main

import (
	"log"
	"net/http"
	"os"

	"github.com/turingXgo/http/app"
	"github.com/turingXgo/http/db"
	"github.com/turingXgo/http/postgres"
)

func main() {
	var (
		dbUser = os.Getenv("DB_USER")
		dbPass = os.Getenv("DB_PASS")
		dbName = os.Getenv("DB_NAME")
	)
	conn, err := db.NewConnection(dbUser, dbPass, dbName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected")
	repo := postgres.New(conn)
	application := app.NewApplication(repo)
	http.HandleFunc("/user", application.GetUserByName)
	http.HandleFunc("/", application.CreateUser)
	http.ListenAndServe(":8081", nil)
}
