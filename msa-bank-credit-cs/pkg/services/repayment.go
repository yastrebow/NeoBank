package services

import (
	"msa-bank-credit-cs/models"
)

type repaymentRepository interface {
	AddEarlyRepayment(e *models.EarlyRepayment) (*models.EarlyRepayment, error)
}

type EarlyRepayment struct {
	repaymentRepo repaymentRepository
}

func NewRepayment(repaymentRepo repaymentRepository) *EarlyRepayment {
	return &EarlyRepayment{
		repaymentRepo: repaymentRepo,
	}
}

func (b *EarlyRepayment) PostEarlyRepayment(credit *models.EarlyRepayment) (*models.EarlyRepayment, error) {
	credit, err := b.repaymentRepo.AddEarlyRepayment(credit)
	if err != nil {
		return nil, err
	}

	return credit, nil
}

