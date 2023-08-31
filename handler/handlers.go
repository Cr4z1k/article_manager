package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Cr4z1k/article_manager/repository"
	"golang.org/x/crypto/bcrypt"
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

func SignUpHandler(repo repository.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var info singUpJSON
		err := json.NewDecoder(r.Body).Decode(&info)
		if err != nil {
			panic(err)
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(info.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}

		result, message := repo.SignUp(info.Username, info.Login, string(hashedPassword), info.Is_author)

		if result {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(message))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(message))
		}
	}
}

func LogInHandler(repo repository.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var info logInJSON
		var passHash string

		err := json.NewDecoder(r.Body).Decode(&info)
		if err != nil {
			panic(err)
		}

		passHash = repo.GetHash(info.Login)

		err = bcrypt.CompareHashAndPassword([]byte(passHash), []byte(info.Password))
		if err == nil {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Logged in successfuly"))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Login or password is not correct"))
		}
	}
}
