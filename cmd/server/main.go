package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"weight-tracker/pkg/api"
	"weight-tracker/pkg/app"

	"github.com/gin-contrib/cors"
	"weight-tracker/pkg/repository"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "This is the startup error: %s\\n", err)
		os.Exit(1)
	}
}

func run() error {
	connectionString := "postgres://postgres:laredo@192.168.49.2:30201/weight-tracker?sslmode=disable"
	db, err := setupDatabase(connectionString)

	if err != nil {
		return err
	}
	//Create storage dependency
	storage := repository.NewStorage(db)

	//Create router dependency
	router := gin.Default()
	router.Use(cors.Default())

	//Create user service
	userService := api.NewUserService(storage)
	//Create Weight service
	weightService := api.NewWeightService(storage)

	server := app.NewServer(router, userService, weightService)

	err = server.Run()

	if err != nil {
		return err
	}
	return nil
}

func setupDatabase(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, err
	}
	err = db.Ping()

	if err != nil {
		return nil, err
	}
	return db, nil
}
