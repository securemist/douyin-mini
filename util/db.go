package util

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	conf "github.com/securemist/douyin-mini/config"
	"log"
)

var url string

func GetDbConnection() *sqlx.DB {

	db, err := sqlx.Open("mysql", conf.Db_url)
	if err != nil {
		log.Fatal("database connection failed")
	}
	return db
}
