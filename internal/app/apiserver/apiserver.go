package apiserver

import (
	"database/sql"
	"github.com/asirkov/http-rest-api/internal/app/store/sqlstore"
	"net/http"
)

func Start(config *Config) error {

	db, err := newDb(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	server := newServer(store)

	return http.ListenAndServe(config.BindAddr, server)
}

func newDb(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
