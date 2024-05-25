package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"sample_banking/db/utils"
	"testing"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("can't connect to db!")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
