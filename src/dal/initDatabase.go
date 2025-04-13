package dal

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
	"github.com/joho/godotenv"
)

var (
	db   *sql.DB
	once sync.Once
)

// InitDB initializes the database connection (Singleton)
func InitDB() *sql.DB {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env File")
		}
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",dbUser,dbPassword,dbHost,dbPort,dbName)
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal("Failed to open DB:", err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatal("Failed to ping DB:", err)
		}
		log.Println("Database connected successfully.")
	})
	return db
}
