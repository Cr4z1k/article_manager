package db

import (
	"database/sql"

	"github.com/Cr4z1k/article_manager/conf"
	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("postgres", conf.GetConnectionString())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
