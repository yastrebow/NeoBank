package repository

import (
	"gorm.io/gorm"

	"msa-bank-client/models"
)

var _ ClientRepository = &clientRepositoryImpl{}

type clientRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) ClientRepository {
	return &clientRepositoryImpl{db}
}

func (c *clientRepositoryImpl) Save(client models.Client) error {
	result := c.DB.Create(&client)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *clientRepositoryImpl) Delete(clientId string) error {
	var client models.Client
	if result := c.DB.First(&client, "id = ?", clientId); result.Error != nil {
		return result.Error
	}
	c.DB.Delete(&client)
	return nil
}

func (c *clientRepositoryImpl) Get(clientId string) (models.Client, error) {
	var client models.Client

	if result := c.DB.First(&client, "id = ?", clientId); result.Error != nil {
		return models.Client{}, result.Error
	}

	return client, nil
}

func (c *clientRepositoryImpl) GetAll() ([]*models.Client, error) {
	var clients []*models.Client
	if result := c.DB.Find(&clients); result.Error != nil {
		return nil, result.Error
	}
	return clients, nil
}

func (c *clientRepositoryImpl) Update(clientId string, updatedClient models.Client) error {
	var client models.Client
	if result := c.DB.First(&client, "id = ?", clientId); result.Error != nil {
		return result.Error
	}
	client.ID = updatedClient.ID
	client.FirstName = updatedClient.FirstName
	client.LastName = updatedClient.LastName
	client.BirthDate = updatedClient.BirthDate
	c.DB.Save(&client)
	return nil
}

func (c *clientRepositoryImpl) GetByPassport(passportNumber string) ([]*models.Client, error) {
	var clients []*models.Client
	if result := c.DB.Find(&clients, "passport_number = ?", passportNumber); result.Error != nil {
		return nil, result.Error
	}
	return clients, nil
}
