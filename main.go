package main

import (
	"context"
	"dobledcloud.com/consumers/handlers"
	"dobledcloud.com/consumers/middleware"
	"dobledcloud.com/consumers/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatalf("Error creating server %v\n", err)
	}
	s.Start(BindRoutes)

}

func BindRoutes(s server.Server, r *mux.Router) {

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	println(c)
	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.CheckAuthMiddleware(s))

	r.HandleFunc("/health", handlers.HealthHandler(s)).Methods(http.MethodGet)
	api.HandleFunc("/contents", handlers.PublishesHandler(s)).Methods(http.MethodGet)

	//r.HandleFunc("/contents", handlers.PublishesHandler(s)).Methods(http.MethodGet)

}
