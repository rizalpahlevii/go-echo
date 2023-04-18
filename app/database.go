package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-echo/helper"
	"time"
)

func NewDB() *sql.DB {
	connectionString := "root@tcp(localhost:3306)/go_echo"
	db, err := sql.Open("mysql", connectionString)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
