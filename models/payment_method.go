package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model

	Id       uuid.UUID `json:"ID" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name     string    `json:"name"`
	IsActive bool      `json:"is_active"`
}

func (PaymentMethod) TableName() string {
	return "mst_payment_methods"
}
