package main

import (
	"log"

	"github.com/SicParv1sMagna/mdhh_backend/internal/api"
)

func main() {
	application, err := api.New()
	if err != nil {
		log.Fatal(err)
	}

	application.StartServer()
}
