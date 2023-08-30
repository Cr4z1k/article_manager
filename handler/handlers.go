package handler

import (
	"fmt"
	"net/http"

	"github.com/Cr4z1k/http_api/repository"
)

func HelloHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello, world")
	}
}

func GetUsersHandler(repo repository.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []string
		rows := repo.GetUsers()

		for rows.Next() {
			var user string

			err := rows.Scan(&user)
			if err != nil {
				panic(err)
			}

			users = append(users, user)
		}
		defer rows.Close()

		fmt.Println(users)
	}
}
