package main

import (
	"fmt"
	"os"

	"t-bonatti/gopj/internal/pkg/database"
	"t-bonatti/gopj/internal/server"

	"github.com/apex/log"
)

func main() {
	database.StartDB("host=localhost port=5432 user=postgres dbname=gopj sslmode=disable password=")
	defer func() {
		if err := database.CloseConn(); err != nil {
			log.WithError(err).Error("failed to close database connections")
		}
	}()

	fmt.Println("Server: ", os.Getpid())

	server := server.New()
	server.Run(":8013")
}
