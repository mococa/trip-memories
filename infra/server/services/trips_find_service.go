package services

import (
	"errors"
	"fmt"
	"tripmemories/server/entities"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (*service) FindTrip(user_pk string, trip_sk string) (*entities.Trip, error) {
	// Returning if user_pk is empty
	if user_pk == "" {
		err := errors.New("User_pk is empty")
		return nil, err
	}

	// Returning if trip_sk is empty
	if user_pk == "" {
		err := errors.New("User_pk is empty")
		return nil, err
	}

	// Getting user from Database
	output, err := trips_repository.FindTrip(user_pk, trip_sk)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to find trip")
	}

	// Checking whether it was found or not
	if len(output.Items) == 0 {
		err = errors.New("Trip not found")
		return nil, err
	}

	// Mapping result to a user
	trip := entities.Trip{}
	err = dynamodbattribute.UnmarshalMap(output.Items[0], &trip)

	if err != nil {
		return nil, err
	}

	return &trip, nil
}
