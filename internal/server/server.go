package server

import (
	"net/http"
	"t-bonatti/gopj/internal/company"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	c := company.NewCompanyController()

	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r.GET("/company/:cnpj", c.Get())

	return r
}
