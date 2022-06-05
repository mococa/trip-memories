package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tripmemories/server/entities"
	"tripmemories/server/services"
)

type UsersController interface {
	FindUserByPk(response http.ResponseWriter, request *http.Request)
	CreateUser(response http.ResponseWriter, request *http.Request)
}

var (
	userService services.UserService
)

func NewUserController(service services.UserService) UsersController {
	userService = service
	return &controller{}
}

func (*controller) CreateUser(res http.ResponseWriter, req *http.Request) {
	// Set response to JSON format
	res.Header().Set("Content-Type", "application/json")

	// Get request body and map it to a user
	var user entities.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{"error": "Error unmarshalling"}`))
		return
	}

	// Create user
	created_user, err := userService.CreateUser(&user)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	// Transform created user into JSON
	response, err := json.Marshal(created_user)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling"}`))
		return
	}

	// Send user as JSON
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(response))

}

func (*controller) FindUserByPk(res http.ResponseWriter, req *http.Request) {
	// Set response to JSON format
	res.Header().Set("Content-Type", "application/json")

	// Getting user_pk query param
	user_pk := req.URL.Query().Get("user_pk")

	// Find user
	user, err := userService.FindUserByPk(user_pk)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	// Transform found user into JSON
	response, err := json.Marshal(user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling"}`))
		return
	}

	// Send user as JSON
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(response))

}
