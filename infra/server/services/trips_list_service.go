package services

import (
	"errors"
	"fmt"
	"tripmemories/server/entities"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (*service) ListTrips(user_pk string) ([]entities.Trip, error) {
	// Returning if user_pk is empty
	if user_pk == "" {
		err := errors.New("User_pk is empty")
		return nil, err
	}

	// Getting trips with this user from Database
	output, err := trips_repository.ListUserTrips(user_pk)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to list trip")
	}

	// Mapping result to a user
	var trips []entities.Trip = []entities.Trip{}

	for _, i := range output.Items {
		trip := entities.Trip{}

		err = dynamodbattribute.UnmarshalMap(i, &trip)
		if err != nil {
			panic("Error mashalling")
		}

		trips = append(trips, trip)
	}

	return trips, nil
}
