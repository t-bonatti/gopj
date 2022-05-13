package company

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyHandler interface {
	Get() gin.HandlerFunc
}

type companyHandlerImpl struct {
	repo CompanyRepository
}

// New creates a new repository
func NewCompanyHandler(r *gin.Engine) {
	repo := NewCompanyRepository()
	handler := &companyHandlerImpl{repo: repo}

	r.GET("/company/:cnpj", handler.Get())
}

// Get a company by cnpj
func (handler *companyHandlerImpl) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		cnpj := c.Param("cnpj")

		company, err := handler.repo.Get(cnpj)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}

		c.JSON(http.StatusOK, company)
	}
}
