package main

import (
	"github.com/fdaines/simple-rest-service/common/log"
	"github.com/fdaines/simple-rest-service/infrastructure"
	"net/http"
)

func main() {
	log.Info("Initializing Handlers...")
	serverMux := infrastructure.SetupMyHandlers()

	log.Info("Starting Server...")
	err := http.ListenAndServe(":10000", serverMux)
	if err != nil {
		log.Fatal(err.Error())
	}
}
