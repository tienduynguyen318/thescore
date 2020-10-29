package main

import (
	"fmt"
	"net/http"
	"os"
	"thescore/internal/datareader"
	"thescore/internal/domain"
	"thescore/internal/httprest"
	"thescore/pkg/logger"
)

func main() {
	logger, err := buildLogger()
	if err != nil {
		fmt.Printf("Failed to create standard logger")
		return
	}
	service := domain.New(domain.Config{
		Logger: logger,
	})
	server, err := httprest.New(httprest.Config{
		Service: service,
		Logger:  logger,
	})
	if err != nil {
		logger.Errorf("Failed to initialize HTTP REST server: %v", err)
		return
	}
	serverPort := ":8080"
	if os.Getenv("PORT") != "" {
		serverPort = fmt.Sprintf(":%s", os.Getenv("PORT"))
	}

	if err = http.ListenAndServe(serverPort, server); err != nil {
		logger.Errorf(err.Error())
	}
}

func init() {
	logger, err := buildLogger()
	if err != nil {
		fmt.Printf("Failed to create standard logger")
		return
	}
	service := domain.New(domain.Config{
		Logger: logger,
	})
	dataReader := datareader.NewDataReader(datareader.Config{
		Logger: logger,
	})
	data, err := dataReader.ReadFile("rushing.json")
	if err != nil {
		logger.Errorf("Unable to read data file", err)
		return
	}
	players, err := dataReader.ParseFile(data)
	if err != nil {
		logger.Errorf("Unable to parse data file", err)
		return
	}
	domainPlayers, err := dataReader.ToDomainPlayers(players)
	if err != nil {
		logger.Errorf("Unable to convert data to player", err)
		return
	}
	err = service.StorePlayers(domainPlayers)
	if err != nil {
		logger.Errorf("Unable to store players data", err)
		return
	}
}

func buildLogger() (logger.Logger, error) {
	config := logger.Config{
		Writer:     os.Stdout,
		AppName:    "TheScore",
		AppVersion: "v1.0",
		Hostname:   "localhost",
	}
	return logger.NewZapLogger(config)
}
