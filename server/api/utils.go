package api

import (
	"encoding/json"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/manager"
)

func FromMap[T any](m map[string]any) (T, error) {
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

func FromJSON[T any](body []byte) (T, error) {
	var d T
	err := json.Unmarshal(body, d)

	if err != nil {
		return d, err
	}

	return d, nil
}

func ToJSON(serializable any) ([]byte, error) {
	return json.Marshal(serializable)
}

func checkCategoryExists(categoryId *string, accountId string, catManager *manager.CategoryManager) *data.RestResponse {
	if categoryId == nil {
		return nil
	}

	cat, err := catManager.Get(*categoryId)
	if err != nil {
		return buildServerError(err)
	}
	if cat == nil || cat.AccountId != accountId {
		return buildBadRequestError()
	}

	return nil
}
