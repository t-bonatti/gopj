package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {

	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	return r
}
