package models

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)


type EarlyRepaymentDTO struct {
	Amount   float32               
	ClientId openapi_types.UUID `gorm:"column:id"` // установить имя столбца `id`
}

type FullRepaymentDTO struct {          
	ClientId openapi_types.UUID `gorm:"column:id"` // установить имя столбца `id`
}