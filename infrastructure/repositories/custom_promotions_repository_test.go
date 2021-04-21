package repositories

import (
	"testing"
)

func TestCustomPromotionsRepository(t *testing.T) {
	t.Run("should return expected promotions", func(t *testing.T) {
		repository := NewPromotionsRepository()

		promotions, _ := repository.GetPromotions()
		if len(promotions) != 4 {
			t.Errorf("got %d promotions, expects %d", len(promotions), 4)
		}
	})
}
