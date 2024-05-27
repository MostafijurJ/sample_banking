package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"sample_banking/api"
	db "sample_banking/db/sqlc"
	"sample_banking/db/utils"
)

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("can't connect to db!")
	}

	queries := db.New(conn)
	server, err := api.NewServer(config, queries)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
