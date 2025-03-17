package main

import (
	"context"
	"fmt"
	"github.com/aliskhanx/goals-api/internal/handler"
	"github.com/aliskhanx/goals-api/internal/repository"
	"github.com/aliskhanx/goals-api/internal/service"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer dbpool.Close()

	repo := repository.NewRepository(dbpool)
	goalService := service.NewGoalService(repo)

	h := handler.NewHandler(goalService)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Goals API is running!")
	})
	mux.HandleFunc("/goals", h.HandleCreateGoal)

	port := ":8080"
	log.Println("Server running on port", port)
	log.Fatalf("Error starting server: %v", http.ListenAndServe(port, mux))
}
