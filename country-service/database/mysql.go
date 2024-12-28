package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type ConfigDB struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

func (config *ConfigDB) InitDB() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)
	log.Printf("Connecting to database at %s:%s/%s", config.DB_HOST, config.DB_PORT, config.DB_NAME)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	// config.autoMigrate(db)
	log.Println("Successfully connected to the database")

	return db
}

// func (config *ConfigDB) autoMigrate(db *sql.DB) {
// 	// Query untuk membuat tabel-tabel yang diperlukan
// 	query := `
// 	CREATE TABLE IF NOT EXISTS countries (
// 		id CHAR(36) PRIMARY KEY,
// 		zone_id VARCHAR(255),
// 		name VARCHAR(255) NOT NULL,
// 		code VARCHAR(10),
// 		code2 VARCHAR(10),
// 		phone_code VARCHAR(10),
// 		image VARCHAR(255),
// 		is_conflict BOOLEAN,
// 		data_link TEXT,
// 		prohibited_items TEXT,
// 		total_diaspora INT,
// 		total_users INT,
// 		shipment FLOAT,
// 		shipment_average FLOAT,
// 		shipment_value FLOAT,
// 		market_share FLOAT,
// 		is_africa INT,
// 		is_north_america INT,
// 		is_south_america INT,
// 		is_oceania INT,
// 		is_asia INT,
// 		is_europe INT,
// 		created_at DATETIME,
// 		updated_at DATETIME,
// 		deleted_at DATETIME,
// 		FOREIGN KEY (zone_id) REFERENCES zones(id)
// 	);

// 	CREATE TABLE IF NOT EXISTS zones (
// 		id CHAR(36) PRIMARY KEY,
// 		name VARCHAR(255) NOT NULL,
// 		code VARCHAR(10),
// 		created_at DATETIME,
// 		updated_at DATETIME,
// 		deleted_at DATETIME
// 	);

// 	CREATE TABLE IF NOT EXISTS shipping_reviews (
// 		id CHAR(36) PRIMARY KEY,
// 		user_id CHAR(36),
// 		country_id CHAR(36),
// 		combined_order_id CHAR(36),
// 		rating INT,
// 		message TEXT,
// 		is_show INT,
// 		FOREIGN KEY (country_id) REFERENCES countries(id)
// 	);
// 	`

// 	// Menjalankan query untuk membuat tabel
// 	_, err := db.Exec(query)
// 	if err != nil {
// 		log.Fatalf("Error creating tables: %v", err)
// 	} else {
// 		log.Println("Tables created successfully")
// 	}
// }
