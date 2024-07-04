package main

import (
	"log"

	"github.com/rafael-ferreira3/poc-api/internal/api"
	"github.com/rafael-ferreira3/poc-api/internal/util"
)

func main() {
	server := api.NewAPIServer(util.GetAddress())
	log.Fatal(server.Run())
}
