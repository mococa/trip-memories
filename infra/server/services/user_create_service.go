package services

import (
	"errors"
	"fmt"
	"time"
	"tripmemories/server/entities"
)

func (*service) CreateUser(user *entities.User) (*entities.User, error) {
	// Making sure user is a user
	if user == nil {
		err := errors.New("User is empty")
		return nil, err
	}

	// Validate fields
	if user.GoogleId == "" || user.Name == "" || user.Email == "" {
		err := errors.New("Missing required fields")
		return nil, err
	}

	// Add PK, SK and timestamp to user
	user.PartitionKey = "user#" + user.GoogleId
	user.SortKey = "root"
	user.CreatedAt = time.Now().UnixMilli()

	fmt.Println(user)
	// Create user in DB
	err := user_repository.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to create user")
	}

	return user, nil
}
