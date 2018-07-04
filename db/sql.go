package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	//If we don't get any errors but somehow still don't get a db connection
	//we exit as well

	if db == nil {
		panic("db nil")
	}

	return db
}

func Migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXIST tasks(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL
	);
	`
	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above

	if err != nil {
		panic(err)
	}
}
