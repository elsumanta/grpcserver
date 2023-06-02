package model

type Register struct {
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Address     string `json:"address" db:"address"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
}
