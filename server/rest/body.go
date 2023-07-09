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

func (b JsonBody) AccountUpdateBody() AccountUpdateBody {
	return AccountUpdateBody{}
}

func (b JsonBody) CategoryCreateBody() CategoryCreateBody {
	return CategoryCreateBody{}
}

func (b JsonBody) CategoryUpdateBody() CategoryUpdateBody {
	return CategoryUpdateBody{}
}

func (b JsonBody) TransactionCreateBody() TransactionCreateBody {
	return TransactionCreateBody{}
}

func (b JsonBody) TransactionUpdateBody() TransactionUpdateBody {
	return TransactionUpdateBody{}
}

func (b JsonBody) BudgetCreateBody() BudgetCreateBody {
	return BudgetCreateBody{}
}

func (b JsonBody) BudgetUpdateBody() BudgetUpdateBody {
	return BudgetUpdateBody{}
}
