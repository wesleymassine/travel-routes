package models

// FlightRepository services interface
type FlightRepository interface {
	CreateFlight(flight *Flight) error
	CheckFlight(flight *Flight) bool
	FlightLoad() ([]Flight, error)
}

// FlightService services interface
type FlightService interface {
	FlightLoad() ([]Flight, error)
	SearchBestRoute(flight Flight) (string, int, error)
	CreateFlight(flight *Flight) error
}
