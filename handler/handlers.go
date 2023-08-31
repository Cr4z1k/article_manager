package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

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

func HTMLHandler(filename string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		html, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, string(html))
	}
}

func UploadFile() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error uploading file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		path := filepath.Join(".", "articles")

		fileName, err := generateFileName(header.Filename, path)
		if err != nil {
			panic(err)
		}

		out, err := os.Create(filepath.Join(path, fileName))
		if err != nil {
			http.Error(w, "Error creating file", http.StatusInternalServerError)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			http.Error(w, "Error copying file", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("File uploaded"))
	}
}

func AddArticleHandler(repo repository.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func generateFileName(fileName string, folderPath string) (string, error) {
	extension := filepath.Ext(fileName)
	nameWithoutExtention := fileName[:len(fileName)-len(extension)]

	for i := 1; ; i++ {
		uniqueName := nameWithoutExtention + "_" + strconv.Itoa(i) + extension
		filePath := filepath.Join(folderPath, uniqueName)
		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			return uniqueName, nil
		} else {
			return "", fmt.Errorf("failed to generate unique filename")
		}
	}
}
