package infrastructure

import (
	"github.com/fdaines/simple-rest-service/infrastructure/repositories"
	"github.com/fdaines/simple-rest-service/usecase"
	"net/http"
)

func SetupMyHandlers() *http.ServeMux {
	mux := http.NewServeMux()

	promotionsRepository := repositories.NewPromotionsRepository()
	promotionsUseCase := usecase.NewPromotionsUseCase(promotionsRepository)
	CreatePromotionsHandler(promotionsUseCase, mux)

	return mux
}
