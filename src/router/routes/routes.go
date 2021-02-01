package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type route struct {
	uri     string
	method  string
	handler func(http.ResponseWriter, *http.Request)
}

// Configure sets up all routes inside the router
func Configure(router *mux.Router) *mux.Router {

	for _, route := range flightsRoutes {
		router.HandleFunc(route.uri, route.handler).Methods(route.method)
	}

	return router
}
