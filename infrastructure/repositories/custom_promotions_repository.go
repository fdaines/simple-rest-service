package repositories

import (
	"github.com/fdaines/simple-rest-service/model"
)

type customPromotionsRepository struct {
//	session *bigquery.BigQuerySession
}

func NewPromotionsRepository() *customPromotionsRepository {
	return &customPromotionsRepository{
//		session,
	}
}

func (c *customPromotionsRepository) GetPromotions() ([]*model.Promotion, error) {
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
