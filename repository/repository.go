package repository

import (
	"database/sql"

	"github.com/lib/pq"
)

type Repository interface {
	GetUsers() *sql.Rows
	SignUp(name string, login string, password_hash string, is_author bool) (bool, string)
	AddArticle(name string, authors pq.Int64Array, themes pq.StringArray, link string, file_path string) bool
	GetHash(login string) string
	CloseConnection()
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
