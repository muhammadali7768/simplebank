package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries

const (
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testPool *pgxpool.Pool

func TestMain(m *testing.M) {
	fmt.Println("TestMain has started......")
	//conn,err := sql.Open(dbDriver,dbSource)
	ctx := context.Background()
	var err error
	testPool, err = pgxpool.New(ctx, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	testQueries = New(testPool)
	defer testPool.Close()

	os.Exit(m.Run())
}
