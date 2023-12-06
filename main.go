package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lenimbugua/vc/api"
	db "github.com/lenimbugua/vc/db/sqlc"
	"github.com/lenimbugua/vc/logger"
	"github.com/lenimbugua/vc/util"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config", err)
		return
	}

	logFile, err := util.LogFile(&config)
	if err != nil {
		log.Fatal("Failed to create or open log file", err)
		return
	}

	logger := logger.NewLogger(&config, logFile)

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database ", err)
	}

	// runDBMigration(config.MigrationURL, config.DBSource)

	store := db.NewSQLStore(conn)

	var httpClient = &http.Client{}

	server, err := api.NewServer(&store, logger, &config, httpClient)
	if err != nil {
		log.Fatal("cannot create server", err)
	}
	var serverAddress string
	if config.Env == "PRODUCTION" {
		port := os.Getenv("PORT")
		serverAddress = ":" + port
	} else {
		serverAddress = config.HTTPServerAddress
	}
	fmt.Println(serverAddress)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server ", err)
	}
}
