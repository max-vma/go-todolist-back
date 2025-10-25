package main

import (
	"database/sql"
	"go-todolist-back/internal/config"
	deliveryHttp "go-todolist-back/internal/delivery/http"
	"go-todolist-back/internal/services"
	"go-todolist-back/internal/storage/postgres"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	cfg := config.MustLoad()

	db, err := sql.Open("pgx", cfg.DatabaseURL)
	if err != nil {
		log.Fatal("failed to open DB:", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping DB:", err)
	}

	userRepo := postgres.NewUserRepository(db)

	userService := services.NewUserService(userRepo)

	userHandler := deliveryHttp.NewUserHandler(userService)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/register", userHandler.Register)

	log.Println("server started on :8080")
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))
}
