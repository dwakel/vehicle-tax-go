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
	config2 "vehicle-tax/config"
	"vehicle-tax/middleware"
	"vehicle-tax/repository"
	"vehicle-tax/services"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("./config.env")
	// GET ENV VARIABLES //
	dbConfig := repository.Config{
		ConnectionString: os.Getenv("DB_CONNECTIONSTRING"),
		Enabled:          true,
		Port:             os.Getenv("DB_PORT"),
		Database:         os.Getenv("DB_NAME"),
	}
	port := os.Getenv("API_PORT")
	token := os.Getenv("TOKEN")
	dKey := os.Getenv("DECRYPTION_KEY")
	downstreamUrl := os.Getenv("DOWNSTREAM_URL")
	downstreamAPIKey := os.Getenv("DOWNSTREAM_API_KEY")
	//INITIALIZE HANDLERS FOR DEPENDENCY INJECTION //
	//Order of initialization matters
	logger := log.New(os.Stdout, "Template Service: ", log.LstdFlags)
	//Setup database and repository
	db := repository.NewMongoDB(&dbConfig, logger)
	repo, _ := db.ConnectMongoDB()
	config := config2.LoadConfig(&config2.Config{
		Token:            &token,
		DecryptionKey:    &dKey,
		DownstreamUrl:    &downstreamUrl,
		DownstreamAPIKey: &downstreamAPIKey,
	})
	dr := repository.NewDataRepo(&repository.DataRepository{Logger: logger, Repo: repo})
	sd := services.NewSampleDownstreamService(&services.SampleDownstreamService{Logger: logger, Config: &config})
	li := api.NewTemplate(logger, &dr, &config, &sd)
	hb := api.NewHeartbeat(logger)

	//Middleware
	mw := middleware.GzipMiddleware{}

	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()

	//Use this router for Gzip compressed request
	getRouterGzip := sm.Methods(http.MethodGet).Subrouter()
	getRouterGzip.Use(mw.GzipMiddleware)

	// HANDLE ROUTES
	getRouter.HandleFunc("/{heartbeat|healthcheck}", hb.Heartbeat)
	getRouter.HandleFunc("/", li.TemplateMethod)
	//Use this route handle for gzip request
	//getRouterGzip.HandleFunc("/", li.TemplateMethod)

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
