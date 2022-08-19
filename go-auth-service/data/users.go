package data

import "github.com/snsvistunov/go-microservices/go-auth-service/models"

var Users = map[string]models.User{
	"Uncle Bob": {
		Username:   "Uncle Bob",
		AuthStatus: true,
	},
	"Gopher": {
		Username:   "Gopher",
		AuthStatus: true,
	},
	"Deve": {
		Username:   "Deve",
		AuthStatus: false,
	},
	"Steve": {
		Username:   "Steve",
		AuthStatus: true,
	},
	"Molly": {
		Username:   "Molly",
		AuthStatus: false,
	},
}
