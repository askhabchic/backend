package main

import (
	"backend/internal/config"
	"backend/pkg/client/postgresql"
	"backend/pkg/logging"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	logger := logging.GetLogger()

	logger.Info("create config")
	cfg := config.GetConfig()

	logger.Info("create router")
	NewRouter(logger, cfg)

	logger.Info("connect to database")
	//newClient, err := postgresql.NewClient(context.Background(), 5, cfg)

	logger.Info("create tables in DB")
	db, err := postgresql.ConnectDatabase(cfg)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}

	//handler := client.NewHandler(logger)

	if err != nil {
		logger.Errorf("Failed to connect database: %s", err)
	}

}

func NewRouter(logger *logging.Logger, cfg *config.Config) {
	myRouter := mux.NewRouter().StrictSlash(true)
	logger.Info(http.ListenAndServe(":"+cfg.Listen.Port, myRouter))
}
