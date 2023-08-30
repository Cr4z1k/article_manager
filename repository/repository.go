package repository

import (
	"database/sql"
)

type Repository interface {
	GetUsers() *sql.Rows
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
