package models

type BudgetType uint

const (
	BudgetType_INCOME  BudgetType = 0
	BudgetType_EXPENSE BudgetType = 1
)

type Budget struct {
	Id         string     `json:"id"`
	AccountId  string     `json:"accountId"`
	CategoryId *string    `json:"categoryId"`
	Name       string     `json:"name"`
	Type       BudgetType `json:"type"`
	Month      int        `json:"month"`
	Year       int        `json:"year"`
	Projected  float64    `json:"projected"`
	Actual     float64    `json:"actual"`
	CreatedAt  int        `json:"createdAt"`
	UpdatedAt  int        `json:"updatedAt"`
}
