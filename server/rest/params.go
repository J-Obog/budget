package rest

type Params map[string]any

func (p Params) TaskId() string {
	return p["taskId"].(string)
}

func (p Params) BudgetId() string {
	return p["budgetId"].(string)
}

func (p Params) CategoryId() string {
	return p["categoryId"].(string)
}

func (p Params) TransactionId() string {
	return p["transactionId"].(string)
}
