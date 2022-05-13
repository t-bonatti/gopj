package company

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	ID           uuid.UUID      `json:"id" gorm:"primaryKey;uuid;default:uuid_generate_v4()"`
	Cnpj         string         `json:"cnpj"`
	OfficialName string         `json:"official_name"`
	FantasyName  string         `json:"fantasy_name"`
	CreatedAt    time.Time      `json:"created"`
	UpdatedAt    time.Time      `json:"updated"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted"`
}
