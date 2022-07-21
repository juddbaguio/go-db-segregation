package users

import (
	"go-db-segregation/business/data"
	"go-db-segregation/business/model"
)

type Service struct {
	repo data.UserRepository
}

func NewCore(user_repo data.UserRepository) *Service {
	return &Service{
		repo: user_repo,
	}
}

func (service *Service) GetUsers() (*[]model.User, error) {
	return service.repo.GetUsers()
}
