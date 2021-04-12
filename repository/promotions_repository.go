package repository

import (
	"github.com/fdaines/simple-rest-service/model"
)

type PromotionsRepository interface {
	GetPromotions() ([]*model.Promotion, error)
}