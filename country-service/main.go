package main

import (
	"fmt"
	"fulka-api/database"
	"fulka-api/routes"
	"fulka-api/util"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Starting application...")

	dbUsername := util.GetConfig("DB_USERNAME")
	dbPassword := util.GetConfig("DB_PASSWORD")
	dbHost := util.GetConfig("DB_HOST")
	dbPort := util.GetConfig("DB_PORT")
	dbName := util.GetConfig("DB_NAME")

	fmt.Println("DB_USERNAME:", dbUsername)
	fmt.Println("DB_PASSWORD:", dbPassword)
	fmt.Println("DB_HOST:", dbHost)
	fmt.Println("DB_PORT:", dbPort)
	fmt.Println("DB_NAME:", dbName)

	configDB := database.ConfigDB{
		DB_USERNAME: dbUsername,
		DB_PASSWORD: dbPassword,
		DB_HOST:     dbHost,
		DB_PORT:     dbPort,
		DB_NAME:     dbName,
	}

	fmt.Println("Initializing database...")
	db := configDB.InitDB()
	if db == nil {
		log.Fatal("Failed to initialize database")
	}
	defer db.Close()
	fmt.Println("Database initialized successfully")

	e := echo.New()

	routes.RegisterRoutes(e, db)

	fmt.Println("Starting server on port 1323...")
	log.Fatal(e.Start(":1323"))
}
