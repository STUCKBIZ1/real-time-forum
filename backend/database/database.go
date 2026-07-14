package database

import (
	"database/sql"
	"fmt"
	"os"

	"real-time-forum/backend/config"
)

var DB *sql.DB

func InitDb() error {
	var err error
	DB, err = sql.Open(config.DB.Driver, config.DB.DSN)
	if err != nil {
		return err
	}

	schema, err := os.ReadFile("database/schema.sql")
	if err != nil {
		return err
	}
	_, err = DB.Exec(string(schema))
	if err != nil {
		DB.Close()
		return err
	}
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}
