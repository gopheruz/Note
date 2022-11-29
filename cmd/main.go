package main

import (
	"Note/storage"
	"Note/storage/repo"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nurmuhammaddeveloper/Note/config"
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
	data := storage.New(db)

	_, err = data.User().Create(&repo.User{
		FirstName:   "Nurmuhammad",
		LastName:    "Hasanov",
		Email:       "nuemuhammad@gmail.com",
		PhoneNumber: "+998915550648",
		ImageUrl:    "salom",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("User created")
	res, err := data.User().Get(1)
	if err!= nil {
        panic(err)
    }
	fmt.Println(res)
	result, err := data.User().GetAll(&repo.GetallUsersParams{
		Limit: 4,
		Page: 1,
		Search: "Nurmuhammad",
		SortBy: "id",
	})
	if err!= nil {
        panic(err)
    }
	for _, val := range result.Users {
		fmt.Println(val)
	}
	fmt.Println(result.Count)

	a, err := data.User().Update(&repo.User{
		ID: 2,
		FirstName:   "New data",
        LastName:    "New lastname",
        Email:       "new email@gmail.com",
        PhoneNumber: "+new phone number",
	})
	if err!= nil {
        panic(err)
    }
	fmt.Println(a)


	err = data.User().Delete(2)
	if err!= nil {
        panic(err)
    }
	fmt.Println("User deleted")
}
