package usecase

import (
	"github.com/fdaines/simple-rest-service/model"
	"testing"
)

type RepositoryStub struct {
	errorResponse error
}

func (tgr *RepositoryStub) GetPromotions() ([]*model.Promotion, error) {
	promotions := []*model.Promotion{
		&model.Promotion{
			Title: "Promotion 1",
			Description: "Brief description of promotion 1",
			StartDate: "2021-01-28",
			EndDate: "2021-01-15",
			IsActive: false,
		},
		&model.Promotion{
			Title: "Promotion 2",
			Description: "Brief description of promotion 2",
			StartDate: "2021-03-28",
			EndDate: "2021-05-28",
			IsActive: true,
		},
		&model.Promotion{
			Title: "Promotion 3",
			Description: "Brief description of promotion 3",
			StartDate: "2021-03-28",
			EndDate: "2021-05-28",
			IsActive: true,
		},
		&model.Promotion{
			Title: "Promotion 4",
			Description: "Brief description of promotion 4",
			StartDate: "2021-03-28",
			EndDate: "2021-05-28",
			IsActive: true,
		},
	}

	return promotions, nil
}


func TestPromotionsUseCase(t *testing.T) {
	t.Run("should return expected promotions", func(t *testing.T) {
		repositoryStub := &RepositoryStub{}
		promotionsUseCase := NewPromotionsUseCase(repositoryStub)

		promotions := promotionsUseCase.ResolveCustomPromotions()
		if len(promotions) != 4 {
			t.Errorf("got %d promotions, expects %d", len(promotions), 4)
		}
	})
}
