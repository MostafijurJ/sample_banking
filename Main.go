package main

import (
	"database/sql"
	"log"
	"sample_banking/api"
	db "sample_banking/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:password@localhost:5432/sample_bank?sslmode=disable"
	serverAddress = "0.0.0.0:9191"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("can't connect to db!")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("can't start server", err)
	}

}
