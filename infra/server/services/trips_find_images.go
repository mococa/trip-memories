package services

import (
	"errors"
	"fmt"
	"strings"
)

func (*service) FindTripImages(user_pk string, trip_sk string) (string, error) {
	var trip_images_only []string

	output, err := trips_repository.FindTripImages(strings.Split(user_pk, "#")[1], strings.Split(trip_sk, "#")[1])
	if err != nil {
		fmt.Println(err)
		return "", errors.New("Failed to list trip images")
	}

	for _, content := range output.Contents {
		trip_images_only = append(trip_images_only, *content.Key)
	}

	response := fmt.Sprintf(`{"images": ["%s"]}`, strings.Join(trip_images_only, "\", \""))
	return response, nil
}
