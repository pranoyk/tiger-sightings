package main

import (
	"net/http"

	"github.com/pranoyk/tiger-sightings/router"
)

func main() {
	router := router.Init()
	http.ListenAndServe(":8080", router)
}
