package main

import (
	"database/sql"
	"github.com/fmo/hexagonal-architecture/internal/app"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	httpHandler := app.SetupApplication(db)

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", httpHandler))
}
