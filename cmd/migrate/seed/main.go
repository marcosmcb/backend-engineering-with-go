package main

import (
	"log"

	"github.com/marcosmcb/backend-engineering-with-go/internal/db"
	"github.com/marcosmcb/backend-engineering-with-go/internal/env"
	"github.com/marcosmcb/backend-engineering-with-go/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	store := store.NewStorage(conn)
	db.Seed(store, conn)
}
