package services

import (
	"errors"
	"fmt"
	"tripmemories/server/entities"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (*service) FindUserByPk(user_pk string) (*entities.User, error) {
	// Returning if user_pk is empty
	if user_pk == "" {
		err := errors.New("User_pk is empty")
		return nil, err
	}

	// Getting user from Database
	output, err := user_repository.FindUserByPk(user_pk)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to find user")
	}

	// Checking whether it was found or not
	if len(output.Items) == 0 {
		err = errors.New("User not found")
		return nil, err
	}

	// Mapping result to a user
	user := entities.User{}
	err = dynamodbattribute.UnmarshalMap(output.Items[0], &user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
