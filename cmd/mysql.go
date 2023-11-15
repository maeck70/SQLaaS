package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// Connect to a MySql instance defined in the config.yaml file
// Use leave blank to connect all, or specify which ones to connect to (as defined in config.yaml)
func ConnectMySql() {

	mysqluri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s",
		settings.Db.DbUser,
		settings.Db.DbPassword,
		settings.Db.DbHost,
		settings.Db.DbPort,
		settings.Db.DbName,
		settings.Db.DbParams,
	)

	log.Printf("Connecting to Database using %s\n", mysqluri)

	var err error
	db, err = sqlx.Connect("mysql", mysqluri)
	FailOnError(err, "Unable to connect to database")

	log.Println("Connected to Database")
}
