package main

import (
	"database/sql"
	eventsrepo "github.com/cdsrx/et/service/sportsevents/db"
	"github.com/cdsrx/et/service/sportsevents/handler"
	pb "github.com/cdsrx/et/service/sportsevents/proto"
	_ "github.com/mattn/go-sqlite3"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"os"
)

const (
	envDBPath = "DBPATH"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("sportsevents"),
	)

	// Set up default db location
	// Note: When DBPATH is not set, micro will create this database in the service instance which will be temporary.
	dbPath := "sportsevents.db"

	envDBPathValue := os.Getenv(envDBPath)
	if envDBPathValue != "" {
		dbPath = envDBPathValue
	}
	logger.Infof("Using path for DB: %s", dbPath)

	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		logger.Fatal(err)
	}

	path, err := os.Getwd()
	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("service instance path: %s", path)
	}

	// Create a new events repository
	eventsRepo, err := eventsrepo.New(database)
	if err != nil {
		logger.Fatal(err)
	}

	// Register handler
	err = pb.RegisterSportsEventsHandler(srv.Server(), handler.New(eventsRepo))
	if err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
