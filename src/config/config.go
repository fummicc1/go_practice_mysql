package config

import "database/sql"

func GetDB(db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "1234"
	dbName := "demo6"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
}
