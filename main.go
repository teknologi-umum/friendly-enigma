package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"refrigerator/handlers"
	"refrigerator/packages/migration"

	"github.com/go-chi/chi/v5"
	// _ (underscore) in Go means the variable is ommited
	// but whatever thing on the right hand side of the
	// operation remains valid.
	//
	// On this case, "go-sqlite3" package is not used
	// as gosqlite3.Something, but it would still be
	// imported. The reason for this is because this package
	// is a driver for the connection of SQLite3 database
	// to Go's database/sql standard library.
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// This line below opens an SQL connection to the SQLite database
	// called house.db on the current working directory.
	db, err := sql.Open("sqlite3", "./house.db")
	if err != nil {
		log.Fatal(err)
	}
	// Defer in Go means "Yeah, I'll execute this function later".
	// Usually this will be executed at the end of the function after
	// the "return" keyword.
	defer db.Close()

	err = migration.Migrate(db, context.Background())
	if err != nil {
		log.Fatal(err)
	}

	h := handlers.Deps{
		DB: db,
	}

	r := chi.NewRouter()
	r.Use(h.HasAccess)
	r.Get("/", h.GetFood)
	r.Post("/", h.AddFood)
	r.Delete("/", h.DeleteFood)
	r.Patch("/", h.UpdateFood)

	http.ListenAndServe(":3000", r)
}
