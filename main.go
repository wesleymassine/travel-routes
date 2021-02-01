package main

import (
	"fmt"
	"log"
	"net/http"
	"travel-routes/src/router"
)

func main() {
	fmt.Println("Application Started listening on port: 5000")
	r := router.Generante()
	log.Fatal(http.ListenAndServe(":5000", r))
}
