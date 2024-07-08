package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rafael-ferreira3/poc-api/internal/api"
	"github.com/rafael-ferreira3/poc-api/internal/config"
	"github.com/rafael-ferreira3/poc-api/internal/database"
	"github.com/rafael-ferreira3/poc-api/internal/helper"
)

func init() {
	config.LoadEnv()
	if err := database.Initialize(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
}

func main() {
	defer database.Close()
	server := api.NewAPIServer(helper.GetAddress())
	if err := server.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
