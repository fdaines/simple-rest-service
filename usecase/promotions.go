package usecase

import (
	"github.com/fdaines/simple-rest-service/model"
	"github.com/fdaines/simple-rest-service/repository"
)

type PromotionsUseCase struct {
	repository repository.PromotionsRepository
}

func NewPromotionsUseCase(promotionsRepository repository.PromotionsRepository) *PromotionsUseCase {
	return &PromotionsUseCase{
		promotionsRepository,
	}
}

func (p *PromotionsUseCase) ResolveCustomPromotions() []*model.Promotion {
	promotions, _ := p.repository.GetPromotions()
	return promotions
}
