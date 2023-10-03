package repository

import (
	"gorm.io/gorm"
	"msa-bank-credit-cs/models"
	"time"
)

type Credit struct {
	db *gorm.DB
}

func NewCredit(db *gorm.DB) *Credit {
	return &Credit{
		db: db,
	}
}

// Create new credit
func (b *Credit) AddCredit(e *models.Credit) (*models.Credit, error) {
	credit := &models.CreditDTO{
		Amount:          e.Amount,
		ClientId:        e.ClientId,
		Id:              e.Id,
		Months:          e.Months,
		Rate:            e.Rate,
		TotalAmount:     e.TotalAmount,
		StartCreditDate: time.Now().UTC(),
	}
	if err := b.db.Table("msa_bank_credit_cs_schema.credit").Create(&credit).First(&credit).Error; err != nil {
		return nil, err
	}
	ret := &models.Credit{
		Amount:       credit.Amount,
		ClientId:     credit.ClientId,
		Id:           credit.Id,
		Months:       credit.Months,
		Rate:         credit.Rate,
		TotalAmount:  credit.TotalAmount,
		MonthPayment: e.MonthPayment, //TODO:
	}
	return ret, nil
}
