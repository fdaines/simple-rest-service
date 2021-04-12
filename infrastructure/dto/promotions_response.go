package dto

import "github.com/fdaines/simple-rest-service/model"

type PromotionsResponse struct {
	Promotions []*model.Promotion `json"promotions"`
}
