package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testConnection *Queries

func Test_Main(m *testing.M) {

	conn, err := sql.Open("postgresql", "postgres://root:123456@localhost:5432/mybank?sslmode=disable")
	if err != nil {
		log.Fatal("ERROR Connection db :", err.Error())
	}
	testConnection = New(conn)

	os.Exit(m.Run())
}
