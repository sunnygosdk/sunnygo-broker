package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8001"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting Broker Service on Port %s \n", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
