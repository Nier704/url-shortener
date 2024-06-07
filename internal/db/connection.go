package db

import (
	"database/sql"
	"fmt"

	"github.com/Nier704/url-shortener/internal/config"
	_ "github.com/lib/pq"
)

func NewDBConnection() (*sql.DB, error) {
	cfg := config.NewPostgresConfig()

	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Name)

	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS urls (
		id UUID PRIMARY KEY NOT NULL,
		url TEXT UNIQUE NOT NULL,
		views INT NOT NULL DEFAULT 0
	)
	`)
	if err != nil {
		return nil, err
	}

	var constraintExists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.constraint_column_usage WHERE table_name = 'urls' AND constraint_name = 'unique_url')").Scan(&constraintExists)
	if err != nil {
		return nil, err
	}

	if !constraintExists {
		_, err = db.Exec(`
    ALTER TABLE urls ADD CONSTRAINT unique_url UNIQUE (url);
    `)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
