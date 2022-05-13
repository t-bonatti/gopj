package company

import (
	"t-bonatti/gopj/internal/pkg/database"

	"gorm.io/gorm"
)

type CompanyRepository interface {
	Get(cnpj string) (c Company, err error)
}

type companyRepositoryImpl struct {
	db *gorm.DB
}

// New creates a new repository
func NewCompanyRepository() CompanyRepository {
	db := database.GetDatabase()
	return &companyRepositoryImpl{db: db}
}

func (r *companyRepositoryImpl) Get(cnpj string) (c Company, err error) {
	err = r.db.Where("cnpj = ?", cnpj).Find(&c).Error
	return
}
