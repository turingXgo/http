package main

import (
	"log"
	"net/http"

	"github.com/turingXgo/http/app"
	"github.com/turingXgo/http/config"
	"github.com/turingXgo/http/db"
	"github.com/turingXgo/http/postgres"
)

func main() {
	settings, err := config.ParseYAML("./.config.yml")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := db.NewConnection(settings.DBUser, settings.DBPass, settings.DBName)
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
