package main

import (
	"log"
	"workshop2/api"
)

func init() {

}

func main() {
	server := api.New()

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
