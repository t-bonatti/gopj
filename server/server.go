package server

import (
	"net/http"
	"t-bonatti/gopj/controller"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	c := controller.NewCompanyController()

	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r.GET("/company/:cnpj", c.Get())

	return r
}
