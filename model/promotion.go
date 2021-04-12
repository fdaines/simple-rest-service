package model

type Promotion struct {
	Title       string `json"title"`
	Description string `json"description"`
	StartDate   string `json"start_date"`
	EndDate     string `json"end_date"`
	IsActive    bool   `json"is_active"`
}
