package services

import (
	"crypto/rand"
	"math"
	"math/big"
	"msa-bank-account-cs/models"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

type accountRepository interface {
	GetAccount(clientId openapi_types.UUID) (*models.Account, error)
	GetAccounts() (*[]models.Account, error)
	AddAccount(e *models.Account) (*models.Account, error)
	UpdateAccount(e *models.Account) (*models.Account, error)
	DeleteAccount(id openapi_types.UUID) (*models.Account, error)
	UpdateAccountBalance(id openapi_types.UUID, e *models.ChangeAccountBalance) (*models.Account, error)
}

type Account struct {
	accountRepo accountRepository
}

func AccountService(accountRepo accountRepository) Account {
	return Account{
		accountRepo: accountRepo,
	}
}

// Get account from the store
// (GET /account/{id})
func (b *Account) GetAccountService(clientId openapi_types.UUID) (*models.Account, error) {
	account, err := b.accountRepo.GetAccount(clientId)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func generateRandomNumber(numberOfDigits int) (int, error) {
	maxLimit := int64(int(math.Pow10(numberOfDigits)) - 1)
	lowLimit := int(math.Pow10(numberOfDigits - 1))

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(maxLimit))
	if err != nil {
		return 0, err
	}
	randomNumberInt := int(randomNumber.Int64())

	if randomNumberInt <= lowLimit {
		randomNumberInt += lowLimit
	}

	if randomNumberInt > int(maxLimit) {
		randomNumberInt = int(maxLimit)
	}
	return randomNumberInt, nil
}

// Add a new account to the store
// (POST /account)
func (b *Account) AddAccountService(e *models.Account) (*models.Account, error) {
	if e.AccountNumber == nil {
		return nil, nil
	}

	randNum, err := generateRandomNumber(7)
	if err != nil {
		return nil, err
	}

	accNum := *e.AccountNumber

	accNum = accNum + string(rune(randNum))

	*e.AccountNumber = accNum
	account, err := b.accountRepo.AddAccount(e)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// Update an existing account
// (PUT /account)
func (b *Account) UpdateAccountService(e *models.Account) (*models.Account, error) {
	account, err := b.accountRepo.UpdateAccount(e)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// Delete account from the store
// (DELETE /account/{id})
func (b *Account) DeleteAccountService(id openapi_types.UUID) (*models.Account, error) {
	account, err := b.accountRepo.DeleteAccount(id)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// Get all accounts from the store
// (GET /accounts)
func (b *Account) GetAccountsService() (*[]models.Account, error) {
	accounts, err := b.accountRepo.GetAccounts()
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// update balance for account from the store
// (POST /account/{id}/update-balance)
func (b *Account) UpdateAccountBalanceService(id openapi_types.UUID, e *models.ChangeAccountBalance) (*models.Account, error) {
	account, err := b.accountRepo.UpdateAccountBalance(id, e)
	if err != nil {
		return nil, err
	}

	return account, nil
}
