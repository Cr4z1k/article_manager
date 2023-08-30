package repository

import (
	"database/sql"
)

func (r *repository) GetUsers() *sql.Rows {
	rows, err := r.db.Query("select name from users")
	if err != nil {
		panic(err)
	}

	//log.Default().Println(check)

	return rows
}

func (r *repository) CloseConnection() {
	r.db.Close()
}
