package services

import (
	"tripmemories/server/entities"
	"tripmemories/server/repositories"
)

type UserService interface {
	CreateUser(user *entities.User) (*entities.User, error)
	FindUserByPk(user_pk string) (*entities.User, error)
}

var (
	user_repository repositories.UsersRepository
)

func NewUserService(repository repositories.UsersRepository) UserService {
	user_repository = repository
	return &service{}
}
