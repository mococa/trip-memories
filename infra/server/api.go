package api

import (
	"fmt"
	"net/http"
	"tripmemories/server/controllers"
	"tripmemories/server/repositories"
	"tripmemories/server/services"

	"github.com/gorilla/mux"
)

var (
	/* -------------- Repositories -------------- */
	usersRepository repositories.UsersRepository = repositories.NewUsersRepository()
	tripsRepository repositories.TripRepository  = repositories.NewTripRepository()

	/* -------------- Services -------------- */
	tripsService services.TripsService = services.NewTripService(tripsRepository)
	userService  services.UserService  = services.NewUserService(usersRepository)

	/* -------------- Controllers -------------- */
	tripsController controllers.TripsController = controllers.NewTripsController(tripsService)
	usersController controllers.UsersController = controllers.NewUserController(userService)
)

func SetupRoutes(router *mux.Router) {
	/* -------------- Index -------------- */
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hello world")
	})

	/* -------------- Trips -------------- */
	router.HandleFunc("/trips", tripsController.FindTrip).Methods("GET")
	router.HandleFunc("/trips/list", tripsController.ListTrips).Methods("GET")
	router.HandleFunc("/trips", tripsController.CreateTrip).Methods("POST")
	router.HandleFunc("/trips/upload-images", tripsController.UploadTripImages).Methods("POST")

	/* -------------- Authentication -------------- */
	router.HandleFunc("/auth", usersController.FindUserByPk).Methods("GET")
	router.HandleFunc("/auth", usersController.CreateUser).Methods("POST")
}
