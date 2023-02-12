package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path"

	"github.com/RickChaves29/url_shortener/internal/data"
	"github.com/RickChaves29/url_shortener/internal/domain"
	_ "github.com/lib/pq"
)

type UrlInput struct {
	OriginUrl string `json:"originUrl"`
}

type UrlOutput struct {
	HashUrl string `json:"hashUrl"`
}

func main() {
	conn, err := data.ConnectionDB()
	if err != nil {
		log.Fatalf("database: %v", err.Error())
	}
	repository := data.NewRepository(conn)
	usecase := domain.NewUsecase(repository)
	http.HandleFunc("/api/code/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" {

			hashUrl := path.Base(r.URL.Path)
			println("path: ", hashUrl)
			result, err := usecase.GetOriginUrlFromRedirect(hashUrl)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Fatalf(err.Error())
			}
			http.Redirect(w, r, result, http.StatusSeeOther)
		}
		if r.Method == "POST" {
			var urlInput *UrlInput

			json.NewDecoder(r.Body).Decode(&urlInput)
			hashUrl, err := usecase.CreateNewUrl(urlInput.OriginUrl)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Fatal(err.Error())
			}
			urlOutput := UrlOutput{HashUrl: hashUrl}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(&urlOutput)
		}
	})
	println("Server runnig on http://localhost:3030/api/code")
	http.ListenAndServe(":3030", nil)
}
