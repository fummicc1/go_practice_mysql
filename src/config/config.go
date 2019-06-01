package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "fummicc1"
	dbPass := "19991004"
	dbName := "localhost"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return db, err
}
