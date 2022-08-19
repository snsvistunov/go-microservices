package data

import "github.com/snsvistunov/go-microservices/go-user-service/models"

var Users = map[string]models.User{
	"Uncle Bob": {
		Username:    "Uncle Bob",
		DateOfBirth: "05-12-1952",
		Age:         69,
		Email:       "unclebob@clean.architecture",
		PhoneNumber: "760-474-1420",
	},
	"Gopher": {
		Username:    "Gopher",
		DateOfBirth: "10-11-2009",
		Age:         13,
		Email:       "go@golang.google.com",
		PhoneNumber: "304-786-2880",
	},
	"Deve": {
		Username:    "Deve",
		DateOfBirth: "29-04-2003",
		Age:         19,
		Email:       "deve@example.com",
		PhoneNumber: "304-786-2880",
	},
	"Steve": {
		Username:    "Steve",
		DateOfBirth: "29-04-2000",
		Age:         22,
		Email:       "steve@example.com",
		PhoneNumber: "304-786-2880",
	},
	"Molly": {
		Username:    "Molly",
		DateOfBirth: "29-04-1992",
		Age:         30,
		Email:       "molly@example.com",
		PhoneNumber: "304-786-2880",
	},
}
