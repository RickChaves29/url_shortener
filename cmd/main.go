package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"regexp"

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

type Response struct {
	Message string `json:"message"`
}

func init() {
	conn, err := data.ConnectionDB()
	if err != nil {
		log.Printf("ERROR DATABASE: %v\n", err.Error())
	}
	_, err = conn.Exec(`
		CREATE TABLE IF NOT EXISTS url (
		id SERIAL PRIMARY KEY,
		origin_url TEXT NOT NULL,
	 	hash_url VARCHAR(6) UNIQUE NOT NULL  
   		)
	`)
	if err != nil {
		log.Printf("ERROR DATABASE: no successfully create url table, %v", err.Error())
	}
	log.Println("DATABASE: successfully create url table")
}

func main() {
	conn, err := data.ConnectionDB()
	if err != nil {
		log.Printf("ERROR DATABASE: %v\n", err.Error())
	}
	repository := data.NewRepository(conn)
	usecase := domain.NewUsecase(repository)
	http.HandleFunc("/api/code/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" {

			hashUrl := path.Base(r.URL.Path)
			result, err := usecase.GetOriginUrlFromRedirect(hashUrl)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("ERROR: %v\n", err.Error())
			}
			http.Redirect(w, r, result, http.StatusSeeOther)
			log.Println("GET/: redirect to original URL")
		}
		if r.Method == "POST" {
			var urlInput *UrlInput
			json.NewDecoder(r.Body).Decode(&urlInput)
			result, err := regexp.MatchString(`(?:https?:\/\/)`, urlInput.OriginUrl)
			if err != nil {
				log.Printf("ERROR REGEX: %v", err.Error())
			}
			if !result {
				w.WriteHeader(http.StatusBadRequest)
				message := Response{Message: "url format is incorrect should have http:// or https://"}
				json.NewEncoder(w).Encode(&message)
				log.Printf("WARN: url format is incorrect")
			} else {
				hashUrl, err := usecase.CreateNewUrl(urlInput.OriginUrl)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					log.Println(err.Error())
				}
				urlOutput := UrlOutput{HashUrl: hashUrl}
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(&urlOutput)
				log.Println("POST/: successfully added new URL")

			}
		}
	})
	println("Server runnig on http://localhost:3030/api/code")
	http.ListenAndServe(":3030", nil)
}
