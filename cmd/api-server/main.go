package main

import (
	"log"

	"github.com/rafael-ferreira3/poc-api/internal/api"
	"github.com/rafael-ferreira3/poc-api/internal/database"
	"github.com/rafael-ferreira3/poc-api/internal/util"
)

func main() {
	defer database.Close()
	server := api.NewAPIServer(util.GetAddress())
	log.Fatal(server.Run())
}
