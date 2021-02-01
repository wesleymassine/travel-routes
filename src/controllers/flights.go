package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"travel-routes/src/models"
	"travel-routes/src/repositories"
	"travel-routes/src/responses"
)

// CreateFlightsRoute handles the flight creation endpoint
func CreateFlightsRoute(w http.ResponseWriter, r *http.Request) {

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please specify the file name: input-routes.csv")
		os.Exit(0)
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	var flight models.Flight
	if err = json.Unmarshal(requestBody, &flight); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	repository, _ := repositories.FlightRepository(models.PathFile + args[0])
	flightRoutes := models.FlightRoute(repository)

	if err = flightRoutes.CreateFlight(&flight); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, flight)
}

// GetFlightsRoute handles the flight get
func GetFlightsRoute(w http.ResponseWriter, r *http.Request) {

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please specify the file name: input-routes.csv")
		os.Exit(0)
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	var flight models.Flight
	if err = json.Unmarshal(requestBody, &flight); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	repository, _ := repositories.FlightRepository(models.PathFile + args[0])
	flightRoutes := models.FlightRoute(repository)

	route, price, err := flightRoutes.SearchBestRoute(flight)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	bestRoute := models.BestRoute{
		Price: price,
		Route: route,
	}

	responses.JSON(w, http.StatusOK, bestRoute)
}
