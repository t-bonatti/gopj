package main

import (
	"fmt"
	"net/http"
	"os"

	"t-bonatti/gopj/internal/company"
	"t-bonatti/gopj/internal/pkg/database"

	"github.com/apex/log"
	"github.com/gin-gonic/gin"
)

func main() {
	database.StartDB("host=localhost port=5432 user=postgres dbname=gopj sslmode=disable password=")
	defer func() {
		if err := database.CloseConn(); err != nil {
			log.WithError(err).Error("failed to close database connections")
		}
	}()

	fmt.Println("Server: ", os.Getpid())

	router := gin.Default()
	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	company.NewCompanyHandler(router)

	router.Run(":8013")
}
