package common

import (
	"database/sql"
	//driver of database
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbName         = ""
	dataSourceName = ""
)

//Connectdatabase : func for connect db
func Connectdatabase() *sql.DB {
	db, err := sql.Open(dbName, dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
