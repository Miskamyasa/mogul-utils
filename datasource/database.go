package datasource

import (
	"database/sql"
	"github.com/Miskamyasa/mogul-utils/notify"
	"os"

	_ "github.com/lib/pq"
)

var conn *sql.DB

func InitDB() *sql.DB {
	var err error
	url := os.Getenv("DATABASE_URL")

	conn, err = sql.Open("postgres", url)
	if err != nil {
		notify.Fatal("Failed to connect to the Database!", err)
	}

	err = conn.Ping()
	if err != nil {
		notify.Fatal("Failed to ping the database", err)
	}

	return conn
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return conn
}
