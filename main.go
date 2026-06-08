package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"database/sql"
	"net/http"
	"os/signal"

	"som/internal/config"
	"som/internal/database"
	"som/internal/routes"
	"som/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title somAPI
// @version 1.0
// @license Copytright
// @BasePath /api/v1
func main() {
	initalizeEnvironment("SOM_ENVIRONMENT")

	db := connectDatabase()
	defer db.Close()

	router := gin.Default()
	routes.RegisterRoutes(router)

	// Graceful shutdown
	cfg := config.Load()
	srv := &http.Server{
		Addr:    cfg.ServerAddr,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err.Error())
		}
	}()
	log.Printf("listening at %s", srv.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}

func initalizeEnvironment(environmentKey string) {
	environment := os.Getenv(environmentKey)
	environment = strings.ToUpper(environment)

	var err error
	switch environment {
	case "PRODUCTION":
		err = godotenv.Load(".env")
	case "DEVELOPMENT":
		err = godotenv.Load(".env.dev")
	default:
		log.Fatalln("unable to load environment")
	}

	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	log.Printf("running in %s environment\n", environment)
}

func connectDatabase() *sql.DB {
	cfg := config.Load()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln("Failed to connect to database: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalln("MariaDB connection failed: ", err)
	}
	log.Println("--- Connected to database successfully ---")

	queries := database.New(db)
	services.SetDatabase(db, queries)

	return db
}
