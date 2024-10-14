package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PostgresConfig struct {
	HOST     string
	PORT     string
	USER     string
	Password string
	DBName   string
	SSLMode  string
}

func (cfg *PostgresConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.HOST, cfg.PORT, cfg.USER, cfg.Password, cfg.DBName, cfg.SSLMode)
}

func NewPostgresStorage(cfg *PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DSN())

	if err != nil {
		log.Fatal(err)
	}

	return db, err
}
