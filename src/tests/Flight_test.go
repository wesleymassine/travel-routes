package tests_test

import (
	"errors"
	"testing"
	"travel-routes/src/exceptions"
	"travel-routes/src/models"
	"travel-routes/src/repositories"
)

// Scenario tests validade
var (
	checkInputAirport  error = errors.New("Airport code must contain only 3 characters")
	checkFlightPrice   error = errors.New("Flight price cannot be 0")
	checkAirportCode   error = errors.New("Airport AAA reported does not exist")
	flightNotFound     error = errors.New("Informed flight not found")
	createFlightExists error = errors.New("Informed flight already exists")
	CreateFlightError  error = errors.New("Error trying to save flights")
)

type flightTestScenario struct {
	From          string `json:"from"`
	To            string `json:"to"`
	Price         int    `json:"price"`
	expectedError error
}

var flightValidateTestScenarios = []flightTestScenario{
	{"GRU", "BR", 10, checkInputAirport},
	{"BRC", "SCL", 0, checkFlightPrice},
	{"GRU", " ", 75, checkInputAirport},
	{"GRU", "SCL", -20, checkFlightPrice},
	{"GRU", "ORL", 56, nil},
	{"ORL", "CDG", 5, nil},
	{"SCL", "ORL", 20, nil},
}

var flightTestScenarios = []flightTestScenario{
	{"GRU", "BRC", 10, nil},
	{"BRC", "SCL", 5, nil},
	{"GRU", "CDG", 75, nil},
	{"GRU", "SCL", 20, nil},
	{"GRU", "ORL", 56, nil},
	{"ORL", "CDG", 5, nil},
	{"SCL", "ORL", 20, nil},
}

func TestFlightValidate(t *testing.T) {
	for _, scenario := range flightValidateTestScenarios {

		flight := models.Flight{
			From:  scenario.From,
			To:    scenario.To,
			Price: scenario.Price,
		}

		if err := flight.Validate(); err == nil {

			if scenario.expectedError != nil {
				t.Errorf("Expected error %s but got error %v", scenario.expectedError.Error(), nil)
			}

		} else if err.Error() != scenario.expectedError.Error() {
			t.Errorf("Expected error %s but got error %v", scenario.expectedError.Error(), err)
		}
	}
}

func TestCheckAirportCode(t *testing.T) {
	flight := models.Flight{
		From:  "BRC",
		To:    "AAA",
		Price: 30,
	}

	err := error(exceptions.CheckAirportCode(flight.To))
	if err.Error() != checkAirportCode.Error() {
		t.Errorf("Expected error: %v but got error: %v", checkAirportCode, err)
	}
}

func TestFlightNotFound(t *testing.T) {
	flight := models.Flight{
		From:  "ABC",
		To:    "CBA",
		Price: 30,
	}

	err := exceptions.FlightNotFound()
	if err.Error() != flightNotFound.Error() {
		t.Errorf("Flight %s-%s expected error: %v but got error: %v", flight.From, flight.To, flightNotFound, err)
	}
}

func TestAddFlight(t *testing.T) {
	flight := models.Flight{
		From:  "BRC",
		To:    "CDG",
		Price: 30,
	}

	flights := &models.Graph{
		Map: make(map[string]map[string]int),
	}

	flights.AddFlight(&flight)

	if flights.Map[flight.From][flight.To] != flight.Price {
		t.Errorf("Flight %s-%s expected error: %v but got error: %v", flight.From, flight.To, flight.Price, flights.Map[flight.From][flight.To])
	}
}

func TestFlight(t *testing.T) {
	repository, _ := repositories.FlightRepository("./input-routes-test.csv")
	service := models.FlightRoute(repository)

	t.Run("FlightLoad", func(t *testing.T) {
		flights, _ := service.FlightLoad()

		for keyC, scenario := range flightTestScenarios {

			flightScenario := models.Flight{
				From:  scenario.From,
				To:    scenario.To,
				Price: scenario.Price,
			}

			for keyR, flight := range flights {
				if keyC == keyR {

					if flight.From != flightScenario.From {
						t.Errorf("Expected error %s - but got error %v", flightScenario.From, flight.From)
					}

					if flight.To != flightScenario.To {
						t.Errorf("Expected error %s but got error %v", flightScenario.To, flight.To)
					}
				}
			}

		}
	})

	t.Run("SearchBestRoute", func(t *testing.T) {
		flightBestRoute := &models.Graph{
			Map: make(map[string]map[string]int),
		}

		flight := &models.Flight{
			From: "GRU",
			To:   "CDG",
		}

		flightBestRoute.AddFlight(flight)

		route, _, err := service.SearchBestRoute(*flight)
		if err == nil {
			if route != "GRU-BRC-CDG" {
				t.Errorf("Expected error: %s - but got error: %v", "GRU-BRC-CDG", route)
			}
		} else {
			t.Errorf("Expected error: %s - but got error: %v", "GRU-BRC-CDG", err)
		}
	})

	t.Run("CreateFlightExist", func(t *testing.T) {

		flight := models.Flight{
			From:  "GRU",
			To:    "CDG",
			Price: 30,
		}

		err := service.CreateFlight(&flight)
		if err.Error() != createFlightExists.Error() {
			t.Errorf("Expected error: %s - but got error: %v", createFlightExists, err)
		}
	})

	t.Run("CreateFlightInvalidCodeAirport", func(t *testing.T) {

		flight := models.Flight{
			From:  "",
			To:    "CDG",
			Price: 30,
		}

		err := service.CreateFlight(&flight)
		if err.Error() != checkInputAirport.Error() {
			t.Errorf("Expected error: %s - but got error: %v", checkInputAirport, err)
		}
	})

	t.Run("CreateFlightInvalidPrice", func(t *testing.T) {
		flight := models.Flight{
			From:  "GRU",
			To:    "CDG",
			Price: 0,
		}

		err := service.CreateFlight(&flight)
		if err.Error() != checkFlightPrice.Error() {
			t.Errorf("Expected error: %s - but got error: %v", checkFlightPrice, err)
		}
	})

	t.Run("CreateFlightError", func(t *testing.T) {
		flight := models.Flight{
			From:  "AAA",
			To:    "AAA",
			Price: 55,
		}

		err := service.CreateFlight(&flight)
		if err.Error() != CreateFlightError.Error() {
			t.Errorf("Expected error: %v - but got error: %v", CreateFlightError, err)
		}
	})

	t.Run("CreateFlightForce", func(t *testing.T) {
		flightBestRoute := &models.Graph{
			Map: make(map[string]map[string]int),
		}

		flight := models.Flight{
			From:  "CGH",
			To:    "SSA",
			Price: 55,
		}

		repository, _ := repositories.FlightRepository("./input-routes-test.csv")
		flightBestRoute.AddFlight(&flight)
		err := repository.CreateFlight(&flight)
		if err != nil {
			t.Errorf("Expected error: %v - but got error: %v", nil, err)
		}
	})
}
