package models

type User struct {
	Username    string `json:"username"`
	DateOfBirth string `json:"date_of_birth"`
	Age         uint   `json:"age"`
	Email       string `json:"e-mail"`
	PhoneNumber string `json:"phone_number"`
}
