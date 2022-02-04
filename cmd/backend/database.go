// database functions
package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	*sql.DB
}

var DB *sql.DB

func Init(connectionString string) *sql.DB {
	// open a new connection
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	// check if the connection is alive
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	DB = db
	return DB
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *sql.DB {
	return DB
}
