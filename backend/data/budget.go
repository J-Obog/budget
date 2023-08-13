package data

type BudgetType string

const (
	BudgetType_Income  BudgetType = "INCOME"
	BudgetType_Expense BudgetType = "EXPENSE"
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
	Id         string
	AccountId  string
	CategoryId string
	Projected  float64
	Timestamp  int64
}

type BudgetFilter struct {
	Month int
	Year  int
}

type BudgetMaterialized struct {
	Budget
	Actual float64 `json:"actual"`
}
