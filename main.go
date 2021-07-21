package main

import (
	"log"
	api2 "workshop2/api"
)

func init() {

}

func main() {
	server := api2.New()

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
