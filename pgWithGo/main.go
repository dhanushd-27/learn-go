package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db  *gorm.DB
	err error
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  bool
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type Book struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

func main() {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=gorm port=5432 sslmode=disable"
	repo := &Repository{}
	repo.db, repo.err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if repo.err != nil {
			panic("failed to connect database")
	}
	fmt.Println("Successfully connected to PostgreSQL")

	repoErr := repo.db.AutoMigrate(&Book{})
	if repoErr != nil {
		log.Fatal("Error migrating database")
	}

	repo.db.Create(&Book{Author: "John Doe", Title: "Book 1"})

	if repo.err != nil {
		log.Fatal("Error connecting to database")
	}

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	log.Println("Server is running on port 8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
