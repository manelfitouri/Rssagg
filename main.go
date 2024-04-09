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

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load() //go loads the envirement varibles
	// os.Getenv  build function (expored package)
	portString := os.Getenv("PORT")
	if portString == "" {
		//log.Fatal will exit the prog with code 1 and a msg
		log.Fatal("PORT is not found in the env")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		//log.Fatal will exit the prog with code 1 and a msg
		log.Fatal("dbuURL is not found in the env")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("can't connect to database:", err)
	}

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	router := chi.NewRouter() //create a new router object
	//allow requests to the server from browser
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	//hook up the httphandler with a specific method and path
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router)

	srv := &http.Server{ //connect the router to http server
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("server stating on port %v", portString)
	err := srv.ListenAndServe() //handling http request or code will block
	if err != nil {
		log.Fatal(err)
	}

}
