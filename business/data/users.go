package data

import "go-db-segregation/business/model"

type UserRepository interface {
	GetUsers() (*[]model.User, error)
}
