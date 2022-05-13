package company

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyController interface {
	Get() gin.HandlerFunc
}

type companyControllerImpl struct {
	r CompanyRepository
}

// New creates a new repository
func NewCompanyController() CompanyController {
	r := NewCompanyRepository()
	return companyControllerImpl{r: r}
}

// Get a company by cnpj
func (ctrl companyControllerImpl) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		cnpj := c.Param("cnpj")

		company, err := ctrl.r.Get(cnpj)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}

		c.JSON(http.StatusOK, company)
	}
}
