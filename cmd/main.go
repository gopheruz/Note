package main

import (
	"fmt"
	"log"

	"github.com/nurmuhammaddeveloper/Note/api"
	"github.com/nurmuhammaddeveloper/Note/storage"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nurmuhammaddeveloper/Note/config"
	_ "github.com/nurmuhammaddeveloper/Note/api/docs"
)

func main() {
	config := config.Load(".")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.Database,
	)

	db, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		log.Fatalf("filed to connect database: %v", err)
	}
	fmt.Println("Database connected")

	strg := storage.New(db)
	apiServer := api.New(&api.RouterOptions{
		Cfg:     &config,
		Storage: strg,
	})
	err = apiServer.Run(config.HttpPort)
	if err != nil {
		log.Fatalf("filed to start server: %v", err)
	}
	log.Print("Server stoped")
}
