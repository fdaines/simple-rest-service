package infrastructure

import (
	"github.com/fdaines/simple-rest-service/model"
	"net/http"
	"net/http/httptest"
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

func TestPromotionsHandler(t *testing.T) {
	t.Run("should return expected promotions", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/promotions", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := SetupMyHandlers()

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check the response body is what we expect.
		expectedResponse := "{\"Promotions\":[{\"Title\":\"Promotion 1\",\"Description\":\"Brief description of promotion 1\",\"StartDate\":\"2021-01-28\",\"EndDate\":\"2021-01-15\",\"IsActive\":false},{\"Title\":\"Promotion 2\",\"Description\":\"Brief description of promotion 2\",\"StartDate\":\"2021-03-28\",\"EndDate\":\"2021-05-28\",\"IsActive\":true},{\"Title\":\"Promotion 3\",\"Description\":\"Brief description of promotion 3\",\"StartDate\":\"2021-03-28\",\"EndDate\":\"2021-05-28\",\"IsActive\":true},{\"Title\":\"Promotion 4\",\"Description\":\"Brief description of promotion 4\",\"StartDate\":\"2021-03-28\",\"EndDate\":\"2021-05-28\",\"IsActive\":true}]}"
		if rr.Body.String() != expectedResponse {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expectedResponse)
		}
	})
}
