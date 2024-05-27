package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"sample_banking/api"
	"sample_banking/db/utils"
	"testing"
	"time"
)

var testQueries *Queries

func newTestServer(t *testing.T, store *Queries) *api.Server {
	config := utils.Config{
		SymmetricKey:  utils.RandomString(32),
		TokenDuration: time.Minute,
	}

	server, err := api.NewServer(config, store)
	if err != nil {
		t.Errorf("failed to create server: %v", err)
	}
	require.NoError(t, err)
	return server
}

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
