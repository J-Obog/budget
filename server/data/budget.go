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
	Period     int64   `json:"period"`
	Projected  float64 `json:"projected"`
	CreatedAt  int64   `json:"createdAt"`
	UpdatedAt  int64   `json:"updatedAt"`
}

type BudgetUpdate struct {
	CategoryId string
	Period     int64
	Projected  float64
}

type BudgetFilter struct {
	CategoryId string
	Period     int64
}

type BudgetMaterialized struct {
	Actual float64 `json:"actual"`
}
