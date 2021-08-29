package logdb

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"matsu.dev/matsuri-contest-log/resources"
)

var ()

var db *sql.DB

func Init(database string) error {
	db_, err := sql.Open("sqlite3", database)
	if err != nil {
		return err
	}
	db = db_

	if _, err := db.Exec(resources.DatabaseStructureSQL); err != nil {
		return err
	}

	return nil
}
