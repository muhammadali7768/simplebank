package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/muhammadali7768/simplebank/util"
)

var testQueries *Queries

var testPool *pgxpool.Pool

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	ctx := context.Background()

	testPool, err = pgxpool.New(ctx, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	testQueries = New(testPool)
	defer testPool.Close()

	os.Exit(m.Run())
}
