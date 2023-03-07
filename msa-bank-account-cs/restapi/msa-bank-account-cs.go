package restapi

import (
	"encoding/json"
	"msa-bank-account-cs/models"
	"msa-bank-account-cs/pkg/services"
	"net/http"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

type accountRequestApi struct{
	service services.Account
}

func NewAccountApi(service services.Account) ServerInterface {
	return &accountRequestApi{
		service: service,
	}
}

	// (GET /account/{id})
func (c *accountRequestApi) GetAccount(w http.ResponseWriter, r *http.Request, clientId openapi_types.UUID)  {

	found, err := c.service.GetAccountService(clientId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if found == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(found)
}

	// Get all accounts from the store
	// (GET /account)
func (c *accountRequestApi)	GetAccounts(w http.ResponseWriter, r *http.Request){

	found, err := c.service.GetAccountsService()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if found == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Add("content-type", "application/json")
	// w.Write(bytes) // NOTE that we should handle the error here
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(found)
}
	// Add a new account to the store
	// (POST /account)
func (c *accountRequestApi)	AddAccount(w http.ResponseWriter, r *http.Request){
	var account models.Account
	json.NewDecoder(r.Body).Decode(&account)
	found, err := c.service.AddAccountService(&account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if found == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(found)
}
	// Update an existing account
	// (PUT /account)
func (c *accountRequestApi)	UpdateAccount(w http.ResponseWriter, r *http.Request){
	var account models.Account
	json.NewDecoder(r.Body).Decode(&account)
	found, err := c.service.UpdateAccountService(&account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if found == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(found)
}
	// Delete account from the store
	// (DELETE /account/{id})
func (c *accountRequestApi)	DeleteAccount(w http.ResponseWriter, r *http.Request, id openapi_types.UUID){
	found, err := c.service.DeleteAccountService(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if found == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(found)
}
	// update balance for account from the store
	// (POST /account/{id}/update-balance)
func (c *accountRequestApi)	UpdateAccountBalance(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	var balance models.ChangeAccountBalance
	json.NewDecoder(r.Body).Decode(&balance)
	found, err := c.service.UpdateAccountBalanceService(id, &balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if found == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(found)
}