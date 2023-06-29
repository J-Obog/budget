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

func buildServerError(res *data.RestResponse, err error) {
	res.Status = http.StatusInternalServerError
}

func buildNotFoundError(res *data.RestResponse) {
	res.Status = http.StatusNotFound
}

func buildForbiddenError(res *data.RestResponse) {
	res.Status = http.StatusForbidden
}

func buildOKResponse(res *data.RestResponse, d interface{}) {
	res.Status = http.StatusOK
	res.Data = d
}

func getAccount(req *data.RestRequest) data.Account {
	return req.Meta["curr_account"].(data.Account)
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
