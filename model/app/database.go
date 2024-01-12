package app

import (
	"database/sql"
	"github.com/Agungsusilo2/Hackfest/model/helper"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:agung@tcp(localhost:3306)/belajar_golang?parseTime=true")

	helper.PanicErr(err)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
