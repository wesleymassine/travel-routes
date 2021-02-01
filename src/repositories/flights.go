package repositories

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"travel-routes/src/exceptions"
	"travel-routes/src/models"
)

// FlightsRepository models flight and name args file data
type FlightsRepository struct {
	Flights  []models.Flight
	Filename string
}

// FlightRepository func main repository
func FlightRepository(filename string) (*FlightsRepository, error) {
	repository := FlightsRepository{Filename: filename}

	err := repository.loadDataFile()
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
		return &FlightsRepository{}, errors.New("Fatal error parsing file CSV")
	}

	return &repository, nil
}

func (r *FlightsRepository) loadDataFile() error {

	file, err := os.Open(r.Filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Generating new file %s\n", r.Filename)
			return err
		}
		return err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	for {
		read, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		price, _ := strconv.Atoi(read[2])
		var flight = models.Flight{
			From:  read[0],
			To:    read[1],
			Price: price,
		}

		if err := flight.Validate(); err != nil {
			return err
		}

		r.Flights = append(r.Flights, flight)
	}

	return nil
}

// CreateFlight create news flights
func (r *FlightsRepository) CreateFlight(flight *models.Flight) error {

	r.Flights = append(r.Flights, *flight)

	if err := exceptions.EqualsAirport(flight.From, flight.To); err != nil {
		return errors.New("Error trying to save flights")
	}

	file, err := os.OpenFile(r.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return errors.New("Error trying to save flights")
	}

	if _, err := file.Write([]byte(fmt.Sprintf("\n%s,%s,%d", flight.From, flight.To, flight.Price))); err != nil {
		log.Fatal(err)
		return errors.New("Error trying to save flights")
	}
	defer file.Close()

	return nil
}

//CheckFlight check if fligth exist
func (r *FlightsRepository) CheckFlight(flight *models.Flight) bool {
	for _, n := range r.Flights {
		if n.From == flight.From && n.To == flight.To {
			return true
		}
	}
	return false
}

// FlightLoad return repository flights
func (r *FlightsRepository) FlightLoad() ([]models.Flight, error) {
	return r.Flights, nil
}
