package main

import (
	"log"

	"github.com/SicParv1sMagna/mdhh_backend/internal/api"
)

// @title MDHH_back
// @version 1.0
// @description bank branch searcher project for hackaton More.Tech VTB

// @license.name AS IS (NO WARRANTY)

// @host 127.0.0.1
// @schemes http
// @BasePath /
func main() {
	application, err := api.New()
	if err != nil {
		log.Fatal(err)
	}

	application.StartServer()
}
