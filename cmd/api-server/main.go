package main

import (
	"log"

	"github.com/rafael-ferreira3/poc-api/internal/api"
)

func main() {
	server := api.NewAPIServer("", "8081")
	log.Fatal(server.Run())
}
