package repository

import (
	"database/sql"
)

type Repository interface {
	GetUsers() *sql.Rows
	SignUp(name string, login string, password_hash string, is_author bool) (bool, string)
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
