package rest

type BudgetQuery struct {
	Month *int `json:"month"`
	Year  *int `json:"year"`
}

type TransactionQuery struct {
	CreatedBefore *int64   `json:"createdBefore"`
	CreatedAfter  *int64   `json:"createdAfter"`
	AmountGte     *float64 `json:"amountGte"`
	AmountLte     *float64 `json:"amountLte"`
}

type AccountUpdateBody struct {
	Name string `json:"name"`
}

type CategoryCreateBody struct {
	Name  string `json:"name"`
	Color uint   `json:"color"`
}

type CategoryUpdateBody struct {
	CategoryCreateBody
}

type TransactionCreateBody struct {
	CategoryId *string `json:"categoryId"`
	Note       *string `json:"note"`
	Amount     float64 `json:"amount"`
	Month      int     `json:"month"`
	Day        int     `json:"day"`
	Year       int     `json:"year"`
}

type TransactionUpdateBody struct {
	TransactionCreateBody
}

type BudgetCreateBody struct {
	CategoryId string  `json:"categoryId"`
	Month      int     `json:"month"`
	Year       int     `json:"year"`
	Projected  float64 `json:"projected"`
}

type BudgetUpdateBody struct {
	BudgetCreateBody
}
