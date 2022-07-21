package main

import (
	"go-db-segregation/business/core/users"
	"go-db-segregation/business/data"
	"go-db-segregation/database/mongo"
	users_mongo "go-db-segregation/database/mongo/users"
	"go-db-segregation/database/postgresql"
	users_postgresql "go-db-segregation/database/postgresql/users"
	"log"
	"os"
)

var dbEngine = "mongo"

func main() {
	var userRepo data.UserRepository

	if dbEngine == "mongo" {
		mongoClient, err := mongo.NewDB()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		userRepo = users_mongo.New(mongoClient)
	} else {
		db, err := postgresql.NewDB()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		userRepo = users_postgresql.New(db)
	}

	userService := users.NewCore(userRepo)
	log.Println(userService.GetUsers())

}
