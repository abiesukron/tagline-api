package app

import (
	"database/sql"
	"tagline_api/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	serverDB := helper.ReadFileEnvToString(".env", "DB_SERVER")
	userDB := helper.ReadFileEnvToString(".env", "DB_USER")
	passDB := helper.ReadFileEnvToString(".env", "DB_PASS")
	namaDB := helper.ReadFileEnvToString(".env", "DB_NAME")
	db, err := sql.Open("mysql", userDB+":"+passDB+"@tcp("+serverDB+":3306)/"+namaDB)
	helper.PanicErrorHandle(err)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(60 * time.Minute)
	return db
}
