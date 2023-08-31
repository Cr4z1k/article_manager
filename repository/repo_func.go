package repository

import (
	"database/sql"
	"encoding/json"
)

type registrationResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (r *repository) GetUsers() *sql.Rows {
	rows, err := r.db.Query("select name from users")
	if err != nil {
		panic(err)
	}

	return rows
}

func (r *repository) SignUp(name string, login string, password_hash string, is_author bool) (bool, string) {
	var resultJSON string
	var registrationResult registrationResult

	err := r.db.QueryRow("select sign_up($1, $2, $3, $4)", name, login, password_hash, is_author).Scan(&resultJSON)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(resultJSON), &registrationResult)
	if err != nil {
		panic(err)
	}

	return registrationResult.Success, registrationResult.Message
}

func (r *repository) GetHash(login string) string {
	var passHash string

	err := r.db.QueryRow("select get_hash($1)", login).Scan(&passHash)
	if err != nil {
		panic(err)
	}

	return passHash
}

func (r *repository) AddArticle(name string, authors []int, themes []string, link string, file_path string) bool {
	var success bool

	err := r.db.QueryRow("select add_article($1, $2, $3, $4, $5)", name, authors, themes, link, file_path).Scan(&success)
	if err != nil {
		panic(err)
	}

	return success
}

func (r *repository) CloseConnection() {
	r.db.Close()
}
