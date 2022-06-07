package services

import (
	"errors"
	"fmt"
	"tripmemories/server/entities"

	"github.com/google/uuid"
)

func (*service) CreateTrip(user_pk string, trip *entities.Trip) (*entities.Trip, error) {
	// Making sure trip is a trip
	if trip == nil {
		err := errors.New("Trip is empty")
		return nil, err
	}

	if user_pk == "" {
		err := errors.New("Unauthorized request")
		return nil, err
	}

	// Validate fields
	if trip.Name == "" || trip.Description == "" {
		err := errors.New("Missing required fields")
		return nil, err
	}

	// Add PK and SK to user
	trip.PartitionKey = user_pk
	trip.SortKey = "trip#" + uuid.NewString()

	// Create trip in DB
	err := trips_repository.CreateTrip(user_pk, trip)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to create trip")
	}

	return trip, nil
}
