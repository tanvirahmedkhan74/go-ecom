package main

import (
	"database/sql"
	"log"

	"github.com/tanvirahmedkhan74/go-ecom/cmd/api"
	"github.com/tanvirahmedkhan74/go-ecom/config"
	"github.com/tanvirahmedkhan74/go-ecom/db"
)

func main() {
	db, err := db.NewPostgresStorage(&db.PostgresConfig{
		HOST:     config.Envs.PublicHost,
		PORT:     config.Envs.Port,
		USER:     config.Envs.DBUser,
		Password: config.Envs.DBPassword,
		DBName:   config.Envs.DBName,
		SSLMode:  config.Envs.SSLMode,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB connected successfully")
}
