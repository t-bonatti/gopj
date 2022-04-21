package datastore

import (
	"t-bonatti/gopj/datastore/database"
	"t-bonatti/gopj/model"

	"gorm.io/gorm"
)

type CompanyRepository interface {
	Get(cnpj string) (c model.Company, err error)
}

type companyRepositoryImpl struct {
	db *gorm.DB
}

// New creates a new repository
func NewCompanyRepository() CompanyRepository {
	db := database.GetDatabase()
	return companyRepositoryImpl{db: db}
}

func (r companyRepositoryImpl) Get(cnpj string) (c model.Company, err error) {
	err = r.db.Where("cnpj = ?", cnpj).Find(&c).Error
	return
}
