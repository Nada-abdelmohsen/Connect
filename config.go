package config

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func GetDB() (db *sql.DB, err error) {
	cfg := mysql.Config{
		User:                 "project",
		Passwd:               "elephant?222",
		Net:                  "tcp",
		Addr:                 "192.168.43.218:3306",
		DBName:               "university",
		AllowNativePasswords: true,
	}
	// Get a database handle.

	db, err = sql.Open("mysql", cfg.FormatDSN())
	return
}
