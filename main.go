package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"vehicle-tax/api"
	v "vehicle-tax/migrator"
	"vehicle-tax/repository"
	service "vehicle-tax/services"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("./.env")
	// GET ENV VARIABLES //
	dbConfig := repository.Config{
		ConnectionString: os.Getenv("DB_CONNECTIONSTRING"),
		Enabled:          true,
		Port:             os.Getenv("DB_PORT"),
		Database:         os.Getenv("DB_NAME"),
	}
	port := os.Getenv("API_PORT")
	//INITIALIZE HANDLERS FOR DEPENDENCY INJECTION //
	//Order of initialization matters
	logger := log.New(os.Stdout, "Import Duty Service: ", log.LstdFlags)
	//Setup database and repository
	db := repository.NewPostgresDB(&dbConfig, logger)
	repo, _ := db.ConnectPostgresDB()

	//Run Migrations here on startup if migrate arg is passed
	runMigrations := os.Args[1]
	fmt.Println("Args:", runMigrations)
	if runMigrations == "migrate" {
		fmt.Println("-----------Running Migrations--------------")
		migrate := v.New(repo, "")
		migErr := migrate.Up()
		if migErr != nil {
			logger.Println("Migration failed")
			logger.Println(migErr)
		}
	}

	taxRepo := repository.NewTaxRepo(logger, repo)
	taxService := service.NewTaxService(logger, taxRepo)

	vc := api.NewVehicleController(logger, &taxService)
	hb := api.NewHeartbeat(logger)

	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	postRouter := sm.Methods(http.MethodPost).Subrouter()

	// HANDLE ROUTES
	getRouter.HandleFunc("/heartbeat", hb.Heartbeat)
	getRouter.HandleFunc("/VehicleCategories", vc.ListCategories)
	getRouter.HandleFunc("/VehicleTypes", vc.ListTypes)
	getRouter.HandleFunc("/TaxInformation", vc.ListTax)

	getRouter.HandleFunc("/VehicleType/{vehicleTypeId}/Duty", vc.CalculateDuty)

	postRouter.HandleFunc("/TaxInformation/SearchSort", vc.ListTaxSearchAndSort)

	//todo: Fetch from configuration file (MAY NOT BE NECESSARY)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      sm,
		IdleTimeout:  2 * time.Second,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	go func() {
		logger.Println(fmt.Sprintf("Starting Server on port: %s", port))
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}
