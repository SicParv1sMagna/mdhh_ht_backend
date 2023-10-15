package main

import (
	"log"

	"github.com/SicParv1sMagna/mdhh_backend/internal/api"
)

// @title RIpPeakBack
// @version 1.0
// @description rip course project about alpinists and their expeditions

// @contact.name Alex Chinaev
// @contact.url https://vk.com/l.chinaev
// @contact.email ax.chinaev@yandex.ru

// @license.name AS IS (NO WARRANTY)

// @host 127.0.0.1E
// @schemes Zhttp
// @BasePath /
func main() {
	application, err := api.New()
	if err != nil {
		log.Fatal(err)
	}

	application.StartServer()
}
