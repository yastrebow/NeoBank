package models

type Report struct {
	Client Client
	Credits []Credit
	Accounts []Account
}

type Client struct {
	ID string 				`json:"id"`
	FirstName string 		`json:"first_name"`
	LastName string 		`json:"last_name"`
	BirthDate string 		`json:"birth_date"`
}

type Credit struct {
	Amount   string        `json:"amount"`
	Id       string        `json:"id"`
	Months   interface{}   `json:"months"`
	Rate     string        `json:"rate"`
	TotalAmount   string   `json:"total_amount"`
}

type Account struct {
	AccountNumber string              `json:"account_number"`
	Amount        float32             `json:"amount"`
	EndDate       string              `json:"end_date"`
	Id            string              `json:"id"`
	StartDate     string              `json:"start_date"`
}
