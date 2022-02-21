package main

import (
	"github.com/apex/log"
	"t-bonatti/gopj/server"
	"t-bonatti/gopj/datastore/database"
)

func main() {
	database.StartDB("host=localhost port=5432 user=postgres dbname=gopj sslmode=disable password=")
	defer func() {
		if err := database.CloseConn(); err != nil {
			log.WithError(err).Error("failed to close database connections")
		}
	}()

	server := server.New()
	server.Run(":8013")
}
