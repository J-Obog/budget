package models

type Budget struct {
	Id         string
	AccountId  string
	CategoryId *string
	Name       string
	//Type: BudgetType
	Month     int
	Year      int
	Projected float64
	Actual    float64
	CreatedAt int
	UpdatedAt int
}
