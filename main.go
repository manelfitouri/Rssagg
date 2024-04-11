package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5" //import packages to spin up the server
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	"github.com/manelfitouri/Rssagg/internal/database"

	_ "github.com/lib/pq" // import the PostgreSQL driver package, to allow the sql package to interact with PostgreSQL databases
)

// apiConfig struct holds configuration options for the API
type apiConfig struct {
	DB *database.Queries // DB is a pointer to the database queries object
}

// main function is the entry point of the application
func main() {
	godotenv.Load(".env") // Load environment variables from the .env file

	port := os.Getenv("PORT") // Retrieve the port from the environment variables
	if port == "" {
		log.Fatal("PORT environment variable is not set") // Log a fatal error if the PORT environment variable is not set
	}

	dbURL := os.Getenv("DATABASE_URL") // Retrieve the database URL from the environment variables
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set") // Log a fatal error if the DATABASE_URL environment variable is not set
	}

	db, err := sql.Open("postgres", dbURL) // Open a connection to the PostgreSQL database using the database URL
	if err != nil {
		log.Fatal(err) // Log a fatal error if there's an error opening the database connection
	}
	dbQueries := database.New(db) // Create a new database queries object

	apiCfg := apiConfig{ // Initialize the API configuration with the database queries object
		DB: dbQueries,
	}

	router := chi.NewRouter() // Create a new chi router

	router.Use(cors.Handler(cors.Options{ // Use CORS middleware to allow cross-origin requests
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter() // Create a new chi router for version 1 of the API

	v1Router.Post("/users", apiCfg.handlerUsersCreate) // Define a route to handle user creation requests
	v1Router.Get("/users", apiCfg.handlerUsersGet)     // Define a route to handle user retrieval requests

	v1Router.Get("/healthz", handlerReadiness) // Define a route to check the health of the server
	v1Router.Get("/err", handlerErr)           // Define a route to handle errors

	router.Mount("/v1", v1Router) // Mount the version 1 router under the "/v1" path prefix

	srv := &http.Server{ // Create a new HTTP server
		Addr:    ":" + port, // Set the server address to the specified port
		Handler: router,     // Set the router as the server's request handler
	}

	log.Printf("Serving on port: %s\n", port) // Log a message indicating that the server is serving on the specified port
	log.Fatal(srv.ListenAndServe())           // Start the HTTP server and log any errors that occur
}
