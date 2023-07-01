package api

import (
	"encoding/json"
	"net/http"

	"github.com/J-Obog/paidoff/data"
)

func isErrorResponse(status int) bool {
	return (status == http.StatusForbidden ||
		status == http.StatusNotFound ||
		status == http.StatusInternalServerError)
}

func FromMap[T interface{}](m map[string]interface{}) (T, error) {
	var d T
	b, err := json.Marshal(m)

	if err != nil {
		return d, err
	}

	err = json.Unmarshal(b, d)

	if err != nil {
		return d, err
	}

	return d, nil
}

func FromJSON[T interface{}](body []byte) (T, error) {
	var d T
	err := json.Unmarshal(body, d)

	if err != nil {
		return d, err
	}

	return d, nil
}

func ToJSON(serializable interface{}) ([]byte, error) {
	return json.Marshal(serializable)
}

func buildUpdatedAccount(timestamp int64, req data.AccountUpdateRequest, ref *data.Account) {
	ref.UpdatedAt = timestamp
	ref.Name = req.Name
}

func buildUpdatedCategory(timestamp int64, req data.CategoryUpdateRequest, ref *data.Category) {
	ref.Color = req.Color
	ref.Name = req.Name
	ref.UpdatedAt = timestamp
}

func buildUpdatedTransaction(timestamp int64, req data.TransactionUpdateRequest, ref *data.Transaction) {
	ref.CategoryId = req.CategoryId
	ref.Description = req.Description
	ref.Amount = req.Amount
	ref.Month = req.Month
	ref.Day = req.Day
	ref.Year = req.Year
	ref.UpdatedAt = timestamp
}

/*
CategoryId *string `json:"categoryId"`
	Name       string  `json:"name"`
	Month      int     `json:"month"`
	Year       int     `json:"year"`
	Projected  float64 `json:"projected"`
	Actual     float64 `json:"actual"`
*/

func buildUpdatedBudget(timestamp int64, req data.BudgetUpdateRequest, ref *data.Budget) {
	ref.CategoryId = req.CategoryId
	ref.Name = req.Name
	ref.Month = req.Month
	ref.Year = req.Year
	ref.Projected = req.Projected
	ref.Actual = req.Actual
	ref.UpdatedAt = timestamp
}

func makeNewBudget(id string, accountId string, timestamp int64, req data.BudgetCreateRequest) data.Budget {
	return data.Budget{
		Id:         id,
		AccountId:  accountId,
		CategoryId: req.CategoryId,
		Name:       req.Name,
		Month:      req.Month,
		Year:       req.Year,
		Projected:  req.Projected,
		Actual:     req.Actual,
		CreatedAt:  timestamp,
		UpdatedAt:  timestamp,
	}
}

func makeNewTransaction(id string, accountId string, timestamp int64, req data.TransactionCreateRequest) data.Transaction {
	return data.Transaction{
		Id:          id,
		AccountId:   accountId,
		CategoryId:  req.CategoryId,
		Description: req.Description,
		Amount:      req.Amount,
		Month:       req.Month,
		Day:         req.Day,
		Year:        req.Year,
		CreatedAt:   timestamp,
		UpdatedAt:   timestamp,
	}
}

func makeNewCategory(id string, accountId string, timestamp int64, req data.CategoryCreateRequest) data.Category {
	return data.Category{
		Id:        id,
		AccountId: accountId,
		Name:      req.Name,
		Color:     req.Color,
		UpdatedAt: timestamp,
		CreatedAt: timestamp,
	}
}
