package rest

type JsonBody []byte

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
	CategoryId  *string `json:"categoryId"`
	Description *string `json:"description"`
	Amount      float64 `json:"amount"`
	Month       int     `json:"month"`
	Day         int     `json:"day"`
	Year        int     `json:"year"`
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

func (b JsonBody) Map() map[string]any {
	return map[string]any{}
}

func (b JsonBody) AccountUpdateBody() (AccountUpdateBody, error) {
	return AccountUpdateBody{}, nil
}

func (b JsonBody) CategoryCreateBody() (CategoryCreateBody, error) {
	return CategoryCreateBody{}, nil
}

func (b JsonBody) CategoryUpdateBody() (CategoryUpdateBody, error) {
	return CategoryUpdateBody{}, nil
}

func (b JsonBody) TransactionCreateBody() (TransactionCreateBody, error) {
	return TransactionCreateBody{}, nil
}

func (b JsonBody) TransactionUpdateBody() (TransactionUpdateBody, error) {
	return TransactionUpdateBody{}, nil
}

func (b JsonBody) BudgetCreateBody() (BudgetCreateBody, error) {
	return BudgetCreateBody{}, nil
}

func (b JsonBody) BudgetUpdateBody() (BudgetUpdateBody, error) {
	return BudgetUpdateBody{}, nil
}
