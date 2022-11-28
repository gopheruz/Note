package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/nurmuhammaddeveloper/Note/config"
	_"github.com/lib/pq"
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

	_, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		log.Fatalf("filed to connect database: %v", err)
	}
	fmt.Println("Database connected")

}
