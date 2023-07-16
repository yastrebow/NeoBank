package repository

import (
	"msa-bank-client/models"
)

type ClientRepository interface {
	Save(client models.Client) error
	Delete(clientId string) error
	Get(clientId string) (models.Client, error)
	GetByPassport(passportNumber string) ([]*models.Client, error)
	GetAll() ([]*models.Client, error)
	Update(clientId string, client models.Client) error
}
