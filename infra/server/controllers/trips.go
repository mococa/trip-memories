package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tripmemories/server/entities"
	"tripmemories/server/services"
)

type TripsController interface {
	CreateTrip(response http.ResponseWriter, request *http.Request)
	FindTrip(response http.ResponseWriter, request *http.Request)
	ListTrips(response http.ResponseWriter, request *http.Request)
}

var (
	tripsService services.TripsService
)

func NewTripsController(service services.TripsService) TripsController {
	tripsService = service

	return &controller{}
}

func (*controller) CreateTrip(res http.ResponseWriter, req *http.Request) {
	// Set response to JSON format
	res.Header().Set("Content-Type", "application/json")

	// Getting user ID
	user_id := req.Header.Get("Authorization")

	// Get request body and map it to a trip
	var trip entities.Trip
	err := json.NewDecoder(req.Body).Decode(&trip)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{"error": "Error unmarshalling"}`))
		return
	}

	// Create trip
	created_trip, err := tripsService.CreateTrip(user_id, &trip)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	// Transform created trip into JSON
	response, err := json.Marshal(created_trip)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling"}`))
		return
	}

	// Send user as JSON
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(response))

}

func (*controller) FindTrip(res http.ResponseWriter, req *http.Request) {
	// Set response to JSON format
	res.Header().Set("Content-Type", "application/json")

	// Getting user ID
	user_id := req.Header.Get("Authorization")

	// Getting trip_sk query params
	trip_sk := req.URL.Query().Get("trip_sk")

	// Find trip
	trip, err := tripsService.FindTrip(user_id, trip_sk)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	// Transform found trip into JSON
	response, err := json.Marshal(trip)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling"}`))
		return
	}

	// Send user as JSON
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(response))

}

func (*controller) ListTrips(res http.ResponseWriter, req *http.Request) {
	// Set response to JSON format
	res.Header().Set("Content-Type", "application/json")

	// Getting user ID
	user_id := req.Header.Get("Authorization")

	// Find trips
	trips, err := tripsService.ListTrips(user_id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	// Transform found trips into JSON
	response, err := json.Marshal(trips)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling"}`))
		return
	}

	// Send user as JSON
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(response))

}
