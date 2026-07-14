package database

import (
	"database/sql"
	"os"
	"real-time-forum/backend/config"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() (*sql.DB, error) {
	var err error
	DB, err = sql.Open(config.DB.Driver, config.DB.DSN)
	if err != nil {
		return nil, err
	}
	schema, err := os.ReadFile("backend/database/schema.sql")
	if err != nil {
		DB.Close()
		return nil, err
	}
	_, err = DB.Exec(string(schema))

	if err != nil {
		DB.Close()
		return nil, err
	}
	return DB, nil
}
