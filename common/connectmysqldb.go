package common

import (
	"database/sql"
	//driver of database
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbName         = "app_fangzhoutest"
	dataSourceName = "douxu:123@tcp(localhost:3306)/sns?charset=utf8"
)

//Connectdatabase : func for connect db
func Connectdatabase() *sql.DB {
	db, err := sql.Open(dbName, dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
