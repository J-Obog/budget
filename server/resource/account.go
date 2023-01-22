package resource

import (
	"encoding/json"
	"net/http"

	"github.com/J-Obog/paidoff/db"
)

type AccountResource struct {
	accountStore db.AccountStore
}

func NewAccountResource(accountStore db.AccountStore) *AccountResource {
	return &AccountResource{
		accountStore: accountStore,
	}
}

func (this *AccountResource) GetAccount(req Request) *Response {
	accountId := mustGetAccountId(req)
	account, err := this.accountStore.Get(accountId)

	if err != nil {
		//return 500
	}

	if account == nil {
		//return 404
	}

	accountResponse := this.toAccountResponse(*account)
	responseBody, err := json.Marshal(accountResponse)

	if err != nil {
		//return 500
	}

	return &Response{
		Body:   responseBody,
		Status: http.StatusOK,
	}
}

func (this *AccountResource) UpdateAccount(req Request) *Response {
	return nil
}

func (this *AccountResource) CreateAccount(req Request) *Response {
	return nil
}

func (this *AccountResource) DeleteAccount(req Request) *Response {
	return nil
}

func (this *AccountResource) toAccountResponse(account db.Account) *AccountResponse {
	return nil
}
