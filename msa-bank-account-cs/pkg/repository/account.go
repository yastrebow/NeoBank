package repository

import (
	"log"
	"msa-bank-account-cs/models"
	"gorm.io/gorm"
	
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

type Account struct {
	db *gorm.DB
}

func NewAccount(db *gorm.DB) *Account {
	return &Account{
		db: db,
	}
}

// Get account
func (b *Account) GetAccount(clientId  openapi_types.UUID) (*models.Account, error) {
	var account  models.Account
	if err := b.db.Table("msa_bank_account_cs_schema.account").First(&account, "id = ?", clientId).Error; err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &account, nil
}

// Get all account
func (b *Account) GetAccounts() (*[]models.Account, error) {
	var accounts  []models.Account
	if err := b.db.Table("msa_bank_account_cs_schema.account").Find(&accounts).Error; err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &accounts, nil
}

// Add account
func (b *Account) AddAccount(e *models.Account) (*models.Account, error) {
	var account  models.Account = *e
	if err := b.db.Table("msa_bank_account_cs_schema.account").Create(&account).First(&account).Error; err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &account, nil
}

// Update account
func (b *Account) UpdateAccount(e *models.Account) (*models.Account, error) {
	var account  models.Account
	if result := b.db.First(&account, "id = ?", e.Id); result.Error != nil {
		log.Fatal(result.Error)
	}
	account.AccountNumber = e.AccountNumber
	account.Amount = e.Amount
	account.ClientId = e.ClientId
	account.StartDate = e.StartDate
	account.EndDate = e.EndDate
	b.db.Save(&account)

	return &account, nil
}

// Delete account
func (b *Account) DeleteAccount(id  openapi_types.UUID) (*models.Account, error) {
	var account  models.Account
	if result := b.db.First(&account, "id = ?", id); result.Error != nil {
		log.Fatal(result.Error)
	}
	// Delete that account
	b.db.Delete(&account)

	return &account, nil
}

// Update account balance
func (b *Account) UpdateAccountBalance(id openapi_types.UUID, e *models.ChangeAccountBalance) (*models.Account, error) {
	var account  models.Account
	if result := b.db.First(&account, "id = ?", id); result.Error != nil {
		log.Fatal(result.Error)
	}

	account.Amount = e.RequestedAmount

	b.db.Save(&account)

	return &account, nil
}
