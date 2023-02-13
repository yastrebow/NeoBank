package services

import (
	"msa-bank-credit-cs/models"
)

type creditRepository interface {
	AddCredit(e *models.Credit) (*models.Credit, error)
}

type Credit struct {
	creditRepo creditRepository
}

func NewCredit(creditRepo creditRepository) *Credit {
	return &Credit{
		creditRepo: creditRepo,
	}
}

func (b *Credit) PostCredit(credit *models.Credit) (*models.Credit, error) {

	credit, err := b.creditRepo.AddCredit(credit)
	if err != nil {
		return nil, err
	}

	return credit, nil
}
