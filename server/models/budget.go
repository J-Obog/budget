package models

type BudgetType uint

const (
	BudgetType_INCOME  BudgetType = 0
	BudgetType_EXPENSE BudgetType = 1
)

type Budget struct {
	Id         string
	AccountId  string
	CategoryId *string
	Name       string
	Type       BudgetType
	Month      int
	Year       int
	Projected  float64
	Actual     float64
	CreatedAt  int
	UpdatedAt  int
}
