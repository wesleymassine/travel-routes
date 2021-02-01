package tests_test

import (
	"testing"
	"travel-routes/src/models"
	"travel-routes/src/repositories"
)

func TestRoute(t *testing.T) {
	repository, _ := repositories.FlightRepository("./input-routes-test.csv")
	service := models.FlightRoute(repository)

	t.Run("SearchBestRouteCheckInputAirport", func(t *testing.T) {
		flightBestRoute := &models.Graph{
			Map: make(map[string]map[string]int),
		}

		route := &models.Route{
			Graphs: flightBestRoute,
		}

		flight := models.Flight{
			From: "BRC",
			To:   "",
		}
		_, _, err := route.SearchBestRoute(flight)

		if err.Error() != checkInputAirport.Error() {
			t.Errorf("Expected error: %s - but got error: %v", checkInputAirport, err)
		}
	})

	t.Run("SearchBestRouteCheckAirportCode", func(t *testing.T) {
		flightBestRoute := &models.Graph{
			Map: make(map[string]map[string]int),
		}

		route := &models.Route{
			Graphs: flightBestRoute,
		}

		flight := models.Flight{
			From: "AAA",
			To:   "CDG",
		}
		_, _, err := route.SearchBestRoute(flight)

		if err.Error() != checkAirportCode.Error() {
			t.Errorf("Expected error: %s - but got error: %v", checkAirportCode, err)
		}
	})

	t.Run("SearchBestRouteFlightNotFound", func(t *testing.T) {
		flightBestRoute := &models.Graph{
			Map: make(map[string]map[string]int),
		}

		flight := &models.Flight{
			From: "CDG",
			To:   "BRC",
		}

		flightBestRoute.AddFlight(flight)

		_, _, err := service.SearchBestRoute(*flight)
		if err.Error() != flightNotFound.Error() {
			t.Errorf("Expected error: %s - but got error: %v", flightNotFound, err)
		}
	})

	t.Run("SearchBestRoute", func(t *testing.T) {
		flightBestRoute := &models.Graph{
			Map: make(map[string]map[string]int),
		}

		flight := &models.Flight{
			From: "GRU",
			To:   "ORL",
		}

		flightBestRoute.AddFlight(flight)

		route, price, err := service.SearchBestRoute(*flight)
		if err == nil {

			if route != "GRU-BRC-SCL-ORL" {
				t.Errorf("Expected error: %s - but got error: %v", "GRU-BRC-SCL-ORL", route)
			}

			if price != 35 {
				t.Errorf("Expected error: %d - but got error: %v", 35, price)

			}

		} else {
			t.Errorf("Expected error: %s - but got error: %v", "GRU-BRC-SCL-ORL > 35", err)
		}

	})
}
