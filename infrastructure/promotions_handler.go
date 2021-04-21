package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/fdaines/simple-rest-service/common/log"
	"github.com/fdaines/simple-rest-service/infrastructure/dto"
	"github.com/fdaines/simple-rest-service/usecase"
	"net/http"
)

type PromotionsHandler struct {
	useCase *usecase.PromotionsUseCase
}

func createPromotionsHandler(useCase *usecase.PromotionsUseCase, mux *http.ServeMux) {
	handler := &PromotionsHandler{useCase }
	mux.HandleFunc("/promotions", handler.handlerFunction)
}

func (p *PromotionsHandler) handlerFunction(w http.ResponseWriter, r *http.Request) {
	promotions := p.useCase.ResolveCustomPromotions()

	response := &dto.PromotionsResponse{
		Promotions: promotions,
	}
	resp, err := json.Marshal(response)
//	resp, err := p.mapper.MapToDto(promotions)
	if err != nil {
		log.Error("Error: %s", err.Error())
		return
	}
	fmt.Fprintf(w, string(resp))
}

