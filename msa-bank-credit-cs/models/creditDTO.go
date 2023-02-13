package models

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

type CreditDTO struct {
	Amount   float32               
	ClientId openapi_types.UUID `gorm:"column:client_id" json:"clientId"`
	Id       openapi_types.UUID `gorm:"primary_key;column:id" json:"id"`
	Months   int                
	Rate     float32           
	TotalAmount float32
}