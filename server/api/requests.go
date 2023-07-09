package api

import "github.com/J-Obog/paidoff/data"

func getAccountCtx(req *data.RestRequest) data.Account {
	return *req.Account
}

func getAccountUpdateBody(req *data.RestRequest) (data.AccountUpdateRequest, error) {
	return FromJSON[data.AccountUpdateRequest](req.Body)
}

func getBugetCreateBody(req *data.RestRequest) (data.BudgetCreateRequest, error) {
	return FromJSON[data.BudgetCreateRequest](req.Body)
}

func getBugetUpdateBody(req *data.RestRequest) (data.BudgetUpdateRequest, error) {
	return FromJSON[data.BudgetUpdateRequest](req.Body)
}

func getBugetGetQuery(req *data.RestRequest) (data.BudgetQuery, error) {
	return FromMap[data.BudgetQuery](req.Query)
}

func getTransactionCreateBody(req *data.RestRequest) (data.TransactionCreateRequest, error) {
	return FromJSON[data.TransactionCreateRequest](req.Body)
}

func getTransactionUpdateBody(req *data.RestRequest) (data.TransactionUpdateRequest, error) {
	return FromJSON[data.TransactionUpdateRequest](req.Body)
}

func getTransactionGetQuery(req *data.RestRequest) (data.TransactionQuery, error) {
	return FromMap[data.TransactionQuery](req.Query)
}

func getCategoryCreateBody(req *data.RestRequest) (data.CategoryCreateRequest, error) {
	return FromJSON[data.CategoryCreateRequest](req.Body)
}

func getCategoryUpdateBody(req *data.RestRequest) (data.CategoryUpdateRequest, error) {
	return FromJSON[data.CategoryUpdateRequest](req.Body)
}
