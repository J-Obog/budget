package data

type BudgetType string

const (
	BudgetType_Income  BudgetType = "Income"
	BudgetType_Expense BudgetType = "Expense"
)

type Budget struct {
	Id         string  `json:"id"`
	AccountId  string  `json:"accountId"`
	CategoryId string  `json:"categoryId"`
	Month      int     `json:"month"`
	Year       int     `json:"year"`
	Projected  float64 `json:"projected"`
	CreatedAt  int64   `json:"createdAt"`
	UpdatedAt  int64   `json:"updatedAt"`
}

type BudgetUpdate struct {
	CategoryId string
	Projected  float64
}

type BudgetFilter struct {
	Month int
	Year  int
}

type BudgetMaterialized struct {
	Actual float64 `json:"actual"`
}
