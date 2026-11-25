package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")

	godotenv.Load(".env")
	portstring := os.Getenv("PORT")
	if portstring == "" {
		log.Fatal("	PORT environment variable not set")
	}
	fmt.Println("Port:", portstring)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthy", handleReadiness)
	v1Router.Get("/err", handleError)

	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portstring,
	}
	log.Println("Starting server on port", portstring)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
