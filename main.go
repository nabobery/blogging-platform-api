package main

import (
	"log"
	"net/http"
	"os"

	"blogging-platform-api/controllers"
	"blogging-platform-api/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables from .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Get the connection string from environment variable (e.g., DATABASE_URL)
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// Open a connection to digital Postgres database through GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// Run auto migration to create/update the underlying database table schema for BlogPost

	if err := db.AutoMigrate(&models.BlogPost{}); err != nil {
		log.Fatal("Error occurred in auto-migration:", err)
	}

	// Initialize the router
	router := mux.NewRouter()

	// Initialize blog post controller, injecting the DB connection
	bpController := controllers.NewBlogPostController(db)

	// Setup RESTful endpoints
	router.HandleFunc("/posts", bpController.CreateBlogPost).Methods("POST")
	router.HandleFunc("/posts", bpController.GetBlogPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", bpController.GetBlogPost).Methods("GET")
	router.HandleFunc("/posts/{id}", bpController.UpdateBlogPost).Methods("PUT")
	router.HandleFunc("/posts/{id}", bpController.DeleteBlogPost).Methods("DELETE")

	// Start the HTTP server
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
