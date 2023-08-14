package data

type Transaction struct {
	Id         string     `json:"id"`
	AccountId  string     `json:"accountId"`
	CategoryId *string    `json:"categoryId"`
	Note       *string    `json:"note"`
	Type       BudgetType `json:"type"`
	Amount     float64    `json:"amount"`
	Month      int        `json:"month"`
	Day        int        `json:"day"`
	Year       int        `json:"year"`
	CreatedAt  int64      `json:"createdAt"`
	UpdatedAt  int64      `json:"updatedAt"`
}

type TransactionFilter struct {
	StartDate *Date
	EndDate   *Date
	MinAmount *float64
	MaxAmount *float64
}

type TransactionUpdate struct {
	CategoryId *string
	Note       *string
	Type       BudgetType
	Amount     float64
	Month      int
	Day        int
	Year       int
}
