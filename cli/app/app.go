package app

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"travel-routes/src/models"
	"travel-routes/src/repositories"
	"travel-routes/src/utils"

	"github.com/urfave/cli"
)

// PathFile cli search best route
const PathFile = "./file/"

//Cli interface
func Cli() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line application"
	app.Usage = "Search search for the best flight routes"
	readInputCli()
	return app
}

func readInputCli() {
	var route string

	for {
		fmt.Print("Please enter the route: ")
		fmt.Scanf("%s \n", &route)
		from, to, err := splitRoute(route, "-")
		if err == nil {
			route, err = searchBestRoute(from, to)
			if err == nil {
				fmt.Printf("Best route:  %s\n", route)
			}
		}
	}
}

func searchBestRoute(from string, to string) (string, error) {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please specify the file name: input-routes.csv")
		os.Exit(0)
	}

	input := &models.Flight{
		From: from,
		To:   to,
	}

	repository, _ := repositories.FlightRepository(PathFile + args[0])
	flightRoutes := models.FlightRoute(repository)

	routes, price, err := flightRoutes.SearchBestRoute(*input)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return fmt.Sprintf("%s > %d", routes, price), nil
}

func splitRoute(input string, separated string) (string, string, error) {
	s := strings.Split(input, separated)

	if len(s) == 1 {
		println("\nInvalid Format! Enter the correct format: GRU-CDG\n")
		return "", "", errors.New("Invalid Format")
	}
	from, to := utils.StringTreatment(s[0], s[1])

	return from, to, nil
}
