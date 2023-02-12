package data

import (
	"database/sql"

	"github.com/RickChaves29/url_shortener/internal/domain"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}

func (conn repository) Create(originUrl, hashUrl string) error {
	insert, err := conn.db.Prepare("INSERT INTO url (origin_url, hash_url) VALUES ($1,$2)")
	if err != nil {
		return err
	}
	_, err = insert.Exec(originUrl, hashUrl)
	if err != nil {
		return err
	}
	defer insert.Close()
	return nil
}

func (conn repository) FindByHashUrl(hashUrl string) (domain.UrlEntity, error) {
	var url domain.UrlEntity

	err := conn.db.QueryRow("SELECT * FROM url WHERE hash_url = $1", hashUrl).Scan(&url.ID, &url.OriginUrl, &url.HashUrl)
	if err != nil {
		return url, err
	}
	return url, nil

}
