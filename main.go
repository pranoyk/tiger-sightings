package main

import (
	"net/http"

	"github.com/pranoyk/tiger-sightings/db"
	"github.com/pranoyk/tiger-sightings/router"
)

func main() {
	conn := db.Init()
	defer conn.Close()
	router := router.Init(conn)
	http.ListenAndServe(":8080", router)
}
