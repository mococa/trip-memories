package services

import (
	"tripmemories/server/entities"
	"tripmemories/server/repositories"
)

type TripsService interface {
	CreateTrip(user_pk string, trip *entities.Trip) (*entities.Trip, error)
	FindTrip(user_pk string, trip_sk string) (*entities.Trip, error)
	ListTrips(user_pk string) ([]entities.Trip, error)
	UploadTripImages(user_pk string, trip_images []entities.TripImage) (string, error)
	TripAddBanner(user_pk string, trip_sk string, encodedImages string) error
	FindTripImages(user_pk string, trip_sk string) (string, error)
}

var (
	trips_repository repositories.TripRepository
)

func NewTripService(repository repositories.TripRepository) TripsService {
	trips_repository = repository
	return &service{}
}
