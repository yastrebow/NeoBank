package repository

import (
	"msa-bank-credit-cs/models"
	"gorm.io/gorm"
	log "github.com/sirupsen/logrus"
)

type Repayment struct {
	db *gorm.DB
}

func NewRepayment(db *gorm.DB) *Repayment {
	return &Repayment{
		db: db,
	}
}


// Creating a request early Repayment
func (b *Repayment) AddEarlyRepayment(e *models.EarlyRepayment) (*models.EarlyRepayment, error) {
	// repayment := &models.EarlyRepaymentDTO{
	// 	Amount:   e.Amount,
	// 	ClientId: e.ClientId,
	// }
		if err := b.db.Table("msa_bank_credit_cs_schema.credit").Model(&e).Updates(map[string]interface{}{"amount": e.Amount, "id": e.Id}).First(&e).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return e, nil
}
