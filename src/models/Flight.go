package models

import (
	"errors"
	"travel-routes/src/exceptions"
	"travel-routes/src/utils"
)

// PathFile Rest API
const PathFile = "./file/"

// Flight that will be stored in the database/file
type Flight struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Price int    `json:"price"`
}

// Graph is a rappresentation of how the points in our graph are connected
type Graph struct {
	Map map[string]map[string]int
}

// FlightRoutes connections routes
type FlightRoutes struct {
	Repository FlightRepository
	Route      Route
	Graphs     Graph
}

// BestRoute result search best route
type BestRoute struct {
	Price int    `json:"price"`
	Route string `json:"route"`
}

// FlightRoute get flights routes
func FlightRoute(flightRepository FlightRepository) *FlightRoutes {
	var flights = Graph{make(map[string]map[string]int)}

	flightRoutes := FlightRoutes{Repository: flightRepository, Graphs: flights, Route: Route{&flights}}
	flightRoutes.getFlights()

	return &flightRoutes
}

func (c *FlightRoutes) getFlights() {
	flightsList, _ := c.Repository.FlightLoad()

	for _, flight := range flightsList {
		c.Graphs.AddFlight(&flight)
	}
}

// AddFlight Graph aiports points
func (g *Graph) AddFlight(flight *Flight) {
	vertex, check := g.Map[flight.From]
	if !check {
		vertex = map[string]int{}
		g.Map[flight.From] = vertex
	}

	vertex[flight.To] = flight.Price
}

// CreateFlight repository flights
func (c *FlightRoutes) CreateFlight(flight *Flight) error {
	flight.From, flight.To = utils.StringTreatment(flight.From, flight.To)

	if err := flight.Validate(); err != nil {
		return err
	}

	if c.Repository.CheckFlight(flight) {
		return errors.New("Informed flight already exists")
	}

	if err := c.Repository.CreateFlight(flight); err != nil {
		return err
	}

	//c.Graphs.AddFlight(flight)

	return nil
}

// FlightLoad return flights respository
func (c *FlightRoutes) FlightLoad() ([]Flight, error) {
	return c.Repository.FlightLoad()
}

// SearchBestRoute return search best routes
func (c *FlightRoutes) SearchBestRoute(flight Flight) (string, int, error) {
	return c.Route.SearchBestRoute(flight)
}

// Validate data when creating/searching a route flight
func (f Flight) Validate() error {

	if err := exceptions.CheckInputAirport(f.From); err != nil {
		return err
	}

	if err := exceptions.CheckInputAirport(f.To); err != nil {
		return err
	}

	if err := exceptions.CheckFlightPrice(f.Price); err != nil {
		return err
	}

	return nil
}
