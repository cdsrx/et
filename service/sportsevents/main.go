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
	"strconv"
)

const (
	envDBPath = "DBPATH"
	envDBSeed = "DBSEED"
)

func main() {
	// Create service using sportsevents as the default name when it is not set during runtime
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
	logger.Infof("using path for DB: %s\n", dbPath)

	// Check if we should seed the database
	seedDb := false
	envDBSeedValue := os.Getenv(envDBSeed)
	var err error
	if envDBSeedValue != "" {
		seedDb, err = strconv.ParseBool(envDBSeedValue)
		if err != nil {
			logger.Errorf("couldn't parse DBSEED value [%s], turning it off.\n", envDBSeedValue)
		}
	}
	logger.Infof("seeding DB: %t\n", seedDb)

	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		logger.Fatal(err)
	}

	// Path where the service is currently running
	// Useful for checking where the service files are e.g. db file when DBPATH is not set
	path, err := os.Getwd()
	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("service instance path: %s\n", path)
	}

	// Create a new events repository
	eventsRepo, err := eventsrepo.New(database, seedDb)
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
