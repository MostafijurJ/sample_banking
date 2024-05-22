package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"testing"
)

const (
	dbSource = "postgresql://root:password@localhost:5432/sample_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := pgxpool.New(context.Background(), dbSource)

	if err != nil {
		log.Fatal("can't connect to db!")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
