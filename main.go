package main

import (
	"github.com/fdaines/simple-rest-service/common/log"
	"github.com/fdaines/simple-rest-service/infrastructure"
	"net/http"
)

func main() {
	log.Info("Initializing Handlers...")
	infrastructure.CreatePromotionsHandler()

	log.Info("Starting Server...")
	err := http.ListenAndServe(":10000", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
