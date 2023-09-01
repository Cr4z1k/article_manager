package main

import (
	"net/http"

	"github.com/Cr4z1k/article_manager/db"
	"github.com/Cr4z1k/article_manager/handler"
	"github.com/Cr4z1k/article_manager/repository"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	repo := repository.NewRepository(db.GetConnection())

	defer repo.CloseConnection()

	r.HandleFunc("/", handler.HTMLHandler("upload.html"))
	r.HandleFunc("/users/get", handler.GetUsersHandler(repo)).Methods("GET")
	r.HandleFunc("/signup", handler.SignUpHandler(repo)).Methods("POST")
	r.HandleFunc("/login", handler.LogInHandler(repo)).Methods("POST")
	r.HandleFunc("/articles/upload", handler.AddArticleHandler(repo)).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
