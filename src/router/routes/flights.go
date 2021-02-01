package routes

import (
	"net/http"
	"travel-routes/src/controllers"
)

var flightsRoutes = []route{
	{
		uri:     "/flights",
		method:  http.MethodPost,
		handler: controllers.CreateFlightsRoute,
	},
	{
		uri:     "/flights",
		method:  http.MethodGet,
		handler: controllers.GetFlightsRoute,
	},
}
