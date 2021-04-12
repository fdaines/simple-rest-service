package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/fdaines/simple-rest-service/common/log"
	"github.com/fdaines/simple-rest-service/infrastructure/dto"
	"github.com/fdaines/simple-rest-service/infrastructure/repositories"
	"github.com/fdaines/simple-rest-service/usecase"
	"net/http"
)

type PromotionsHandler struct {
	useCase *usecase.PromotionsUseCase
}


func CreatePromotionsHandler() {
	promotionsRepository := repositories.NewPromotionsRepository()
	promotionsUseCase := usecase.NewPromotionsUseCase(promotionsRepository)
	handler := &PromotionsHandler{promotionsUseCase}
	http.HandleFunc("/", handler.handlerFunction)
}

func (p *PromotionsHandler) handlerFunction(w http.ResponseWriter, r *http.Request) {
	response := &dto.PromotionsResponse{
		Promotions: p.useCase.ResolveCustomPromotions(),
	}
	resp, err := json.Marshal(response)
	if err != nil {
		log.Error("Error: %s", err.Error())
		return
	}
	fmt.Fprintf(w, string(resp))
}
