package main

import (
	"log"

	"github.com/rafael-ferreira3/poc-api/internal/api"
	"github.com/rafael-ferreira3/poc-api/internal/database"
	"github.com/rafael-ferreira3/poc-api/internal/helper"
)

func main() {
	defer database.Close()
	server := api.NewAPIServer(helper.GetAddress())
	log.Fatal(server.Run())
}
