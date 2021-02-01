package models

import (
	"strings"
	"travel-routes/src/exceptions"
	"travel-routes/src/utils"
)

// Route struct map key value flights
type Route struct {
	Graphs *Graph
}

type vertex struct {
	Name  string
	Perm  bool
	Price int
	Path  string
}

const maxPrice int = int(^uint(0) >> 1)

// SearchBestRoute Dijkstra algorithm used for find the shortest path between points of a graph.
func (route *Route) SearchBestRoute(flight Flight) (string, int, error) {

	flight.From, flight.To = utils.StringTreatment(flight.From, flight.To)
	flight.Price++

	if err := flight.Validate(); err != nil {
		return "[]", 0, err
	}

	nodes := make([]string, 0)
	for key := range route.Graphs.Map {
		nodes = append(nodes, key)
		for dest := range route.Graphs.Map[key] {
			nodes = append(nodes, dest)
		}
	}

	nodes = route.checkAirportCode(nodes)
	for _, airportCode := range []string{flight.From, flight.To} {
		if !route.checkAirport(nodes, airportCode) {
			return "[InvalidAirportCode]", 0, exceptions.CheckAirportCode(airportCode)
		}
	}

	vertexList := make(map[string]*vertex, 0)
	var current string
	var candidates = true

	for _, node := range nodes {
		var temp = vertex{Name: node, Perm: false, Price: maxPrice, Path: "-"}

		if node == flight.From {
			temp.Price = 0
			current = node
		}

		vertexList[node] = &temp
	}

	for candidates {
		vertexList[current].Perm = true
		var nextElements = route.searchNodes(vertexList[current].Name)

		for _, element := range nextElements {
			if vertexList[element].Price > vertexList[current].Price+route.Graphs.Map[current][element] {
				vertexList[element].Price = vertexList[current].Price + route.Graphs.Map[current][element]
				vertexList[element].Path = current
			}
		}
		current, candidates = route.searchNextCurrent(vertexList)
	}

	if route.flightNotFound(&vertexList, flight.To) {
		return "[FlightNotFound]", 0, exceptions.FlightNotFound()
	}

	return route.connectionPath(&vertexList, flight.From, flight.To), vertexList[flight.To].Price, nil
}

func (route *Route) checkAirport(airports []string, airport string) bool {
	for _, value := range airports {
		if value == airport {
			return true
		}
	}

	return false
}

func (route *Route) flightNotFound(vertexList *map[string]*vertex, target string) bool {
	return (*vertexList)[target].Price == maxPrice
}

func (route *Route) connectionPath(vertexList *map[string]*vertex, origin string, dest string) string {

	var connection []string
	var current = dest

	for current != origin {
		connection = append([]string{current}, connection...)
		current = (*vertexList)[current].Path
	}

	connection = append([]string{origin}, connection...)

	return strings.Join(connection[:], "-")
}

func (route *Route) checkAirportCode(vector []string) []string {
	elements := make(map[string]bool)
	list := []string{}
	for _, entry := range vector {
		if _, value := elements[entry]; !value {
			elements[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

func (route *Route) searchNextCurrent(vertexList map[string]*vertex) (string, bool) {
	var anyCandidate = false
	var currentPrice = maxPrice
	var currentNode = ""

	for _, vertex := range vertexList {
		if !vertex.Perm && vertex.Price < currentPrice {
			currentPrice = vertex.Price
			currentNode = vertex.Name
			anyCandidate = true
		}
	}

	return currentNode, anyCandidate
}

func (route *Route) searchNodes(origin string) []string {
	candidates := make([]string, 0)

	for node := range route.Graphs.Map[origin] {
		candidates = append(candidates, node)
	}

	return candidates
}
