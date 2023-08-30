package main

import (
	"net/http"

	"github.com/Cr4z1k/http_api/db"
	"github.com/Cr4z1k/http_api/handler"
	"github.com/Cr4z1k/http_api/repository"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	repo := repository.NewRepository(db.GetConnection())

	defer db.GetConnection().Close()

	r.HandleFunc("/", handler.HelloHandler())
	r.HandleFunc("/users/get", handler.GetUsersHandler(repo)).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
