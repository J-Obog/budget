package data

type RestRequest struct {
	Url         string
	Headers     map[string]any
	Meta        map[string]any
	UrlParams   map[string]any
	QueryParams map[string]any
	Metadata    map[string]any
	Body        []byte
}

type RestResponse struct {
	Data   any
	Status int
}

type AccountUpdateRequest struct {
	Name string `json:"name"`
}

type CategoryCreateRequest struct {
	Name  string `json:"name"`
	Color uint   `json:"color"`
}

type CategoryUpdateRequest struct {
	CategoryCreateRequest
}

type TransactionCreateRequest struct {
	CategoryId  *string `json:"categoryId"`
	Description *string `json:"description"`
	Amount      float64 `json:"amount"`
	Month       int     `json:"month"`
	Day         int     `json:"day"`
	Year        int     `json:"year"`
}

type TransactionUpdateRequest struct {
	TransactionCreateRequest
}

type TransactionQuery struct {
	CreatedBefore *int64     `json:"createdBefore"`
	CreatedAfter  *int64     `json:"createdAfter"`
	AmountGte     *float64   `json:"amountGte"`
	AmountLte     *float64   `json:"amountLte"`
	Categories    []Category `json:"categories"`
}

type BudgetCreateRequest struct {
	CategoryId string  `json:"categoryId"`
	Month      int     `json:"month"`
	Year       int     `json:"year"`
	Projected  float64 `json:"projected"`
}

type BudgetUpdateRequest struct {
	BudgetCreateRequest
}

type BudgetQuery struct {
	Month *int `json:"month"`
	Year  *int `json:"year"`
}
