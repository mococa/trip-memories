package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"tripmemories/server/entities"
)

func (*service) TripAddBanner(user_pk string, trip_sk string, encodedImages string) error {
	var decoded_images entities.DecodedImages

	err := json.Unmarshal([]byte(encodedImages), &decoded_images)

	if err != nil {
		return errors.New("error unmarshalling")
	}

	fmt.Println(trip_sk)

	err = trips_repository.AddTripBanner(user_pk, trip_sk, decoded_images.Images[0])
	if err != nil {
		fmt.Println(err)
		return errors.New("error adding banner")
	}

	return nil
}
