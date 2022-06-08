package controllers

import (
	"encoding/json"
	"fmt"
	_ "image/png"
	"net/http"
	"strings"
	"tripmemories/server/entities"
	"tripmemories/server/services"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

type TripsController interface {
	CreateTrip(response http.ResponseWriter, request *http.Request)
	FindTrip(response http.ResponseWriter, request *http.Request)
	ListTrips(response http.ResponseWriter, request *http.Request)
	UploadTripImages(response http.ResponseWriter, request *http.Request)
}

var (
	tripsService services.TripsService
	s3session    *s3.S3
)

func init() {
	s3session = s3.New(session.Must(session.NewSessionWithOptions(session.Options{})))
}

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
		fmt.Println(err)
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

	images, err := tripsService.FindTripImages(user_id, trip_sk)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	var trip_images entities.DecodedImages

	err = json.Unmarshal([]byte(images), &trip_images)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	trip.Images = trip_images.Images

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

func (*controller) UploadTripImages(res http.ResponseWriter, req *http.Request) {
	// Set response to JSON format
	res.Header().Set("Content-Type", "application/json")

	err := req.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "request too long"}`))
		return
	}

	formdata := req.MultipartForm

	user := req.Header.Get("Authorization")
	trip := formdata.Value["trip"][0]

	if user == "" {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Unauthorized request"}`))
		return
	}

	if trip == "" {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Missing trip sk"}`))
		return
	}

	var trip_images []entities.TripImage = []entities.TripImage{}

	for _, files := range formdata.File {
		for i := range files {
			file, err := files[i].Open()
			if err != nil {
				fmt.Println(err)
				res.WriteHeader(http.StatusInternalServerError)
				res.Write([]byte(`{"error": "Error with file"}`))

				return
			}

			defer file.Close()

			filename := files[i].Filename
			filename_split := strings.Split(filename, ".")
			extension := filename_split[len(filename_split)-1]

			bucket_item_name := strings.Split(user, "#")[1] + "/" + strings.Split(trip, "#")[1] + "/" + uuid.NewString() + "." + extension

			trip_images = append(trip_images, entities.TripImage{File: file, Filename: bucket_item_name})

		}
	}

	output, err := tripsService.UploadTripImages(user, trip_images)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))

		return
	}

	err = tripsService.TripAddBanner(user, trip, output)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))

		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(output))
}
