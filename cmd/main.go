package main

import (
	"log"

	"github.com/emirrcaglar/go-restapi/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal("error running api server: ", err)
	}
}
