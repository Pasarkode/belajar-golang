package controllers

import (
	"github.com/coopernurse/gorp"
	"database/sql"
	"aah-app/app/models"
	"aahframework.org/log.v0"
)

func InitDB(table string) *gorp.DbMap {
	db, err := sql.Open("sqlite3", "database.sqlite3")
	checkErr(err, "Database Open failed ")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(models.Books{}, table).SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed ")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}