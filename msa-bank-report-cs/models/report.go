package models

type Report struct {
	Client Client
	Credit []Credit
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

