package main

import (
	"fmt"
	"fulka-api/database"
	"fulka-api/routes"
	"fulka-api/util"
	"log"
	"net/http"
)

func main() {
	dbUsername := util.GetConfig("DB_USERNAME")
	dbPassword := util.GetConfig("DB_PASSWORD")
	dbHost := util.GetConfig("DB_HOST")
	dbPort := util.GetConfig("DB_PORT")
	dbName := util.GetConfig("DB_NAME")
	secreat := util.GetConfig("JWT_SECRET")
	fmt.Println("secreat", secreat)

	configDB := database.ConfigDB{
		DB_USERNAME: dbUsername,
		DB_PASSWORD: dbPassword,
		DB_HOST:     dbHost,
		DB_PORT:     dbPort,
		DB_NAME:     dbName,
	}

	db := configDB.InitDB()
	if db == nil {
		log.Fatal("Failed to initialize database")
	}
	defer db.Close()

	mux := http.NewServeMux()
	routes.RegisterRoutes(mux, db)

	log.Fatal(http.ListenAndServe(":1323", mux))
}
