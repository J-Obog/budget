package rest

type PathParams map[string]string

func (p PathParams) GetBudgetId() string {
	return p["budgetId"]
}

func (p PathParams) GetBudgetPeriodMonth() string {
	return p["periodMonth"]
}

func (p PathParams) GetBudgetPeriodYear() string {
	return p["periodYear"]
}

func (p PathParams) GetCategoryId() string {
	return p["categoryId"]
}

func (p PathParams) GetTransactionId() string {
	return p["transactionId"]
}
