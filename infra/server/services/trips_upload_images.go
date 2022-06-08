package services

import (
	"errors"
	"fmt"
	"strings"
	"tripmemories/server/entities"
)

func (*service) UploadTripImages(user_pk string, trip_images []entities.TripImage) (string, error) {
	var trip_images_only []string

	for _, trip_image := range trip_images {
		_, err := trips_repository.UploadTripImage(trip_image.File, trip_image.Filename)

		trip_images_only = append(trip_images_only, trip_image.Filename)

		if err != nil {
			fmt.Println(err)
			return "", errors.New("Failed to create trip")
		}

	}

	response := fmt.Sprintf(`{"images": ["%s"]}`, strings.Join(trip_images_only, "\", \""))
	return response, nil
}
