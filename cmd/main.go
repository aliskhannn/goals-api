package main

import (
	"context"
	"fmt"
	"github.com/aliskhannn/goals-api/internal/handler"
	"github.com/aliskhannn/goals-api/internal/middlewares"
	"github.com/aliskhannn/goals-api/internal/repository"
	"github.com/aliskhannn/goals-api/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}
	defer dbpool.Close()

	repo := repository.NewRepository(dbpool)
	goalService := service.NewGoalService(repo)
	userService := service.NewUserService(repo)

	h := handler.NewHandler(goalService, userService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	authMiddleware := middlewares.NewAuthMiddleware(os.Getenv("JWT_SECRET_KEY"))

	r.Post("/register", h.HandleRegister)
	r.Post("/login", h.HandleLogin)

	r.With(authMiddleware.ValidateToken).Post("/goals", h.HandleCreateGoal)
	r.With(authMiddleware.ValidateToken).Get("/goals", h.HandleGetAllGoals)
	r.With(authMiddleware.ValidateToken).Get("/goals/{id}", h.HandleGetGoalByID)
	r.With(authMiddleware.ValidateToken).Put("/goals/{id}", h.HandleUpdateGoal)
	r.With(authMiddleware.ValidateToken).Delete("/goals/{id}", h.HandleDeleteGoal)

	port := ":8080"
	logger.Println("Server running on port", port)
	logger.Fatalf("Error starting server: %v", http.ListenAndServe(port, r))
}
