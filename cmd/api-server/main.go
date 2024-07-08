package main

import (
	"fmt"
	"os"

	"github.com/rafael-ferreira3/poc-api/internal/api"
	"github.com/rafael-ferreira3/poc-api/internal/database"
	"github.com/rafael-ferreira3/poc-api/internal/helper"
)

func main() {
	defer database.Close()
	server := api.NewAPIServer(helper.GetAddress())
	if err := server.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
