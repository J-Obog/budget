package resource

import (
	"encoding/json"
	"net/http"

	"github.com/J-Obog/paidoff/db"
	"github.com/J-Obog/paidoff/queue"
)

type AccountResource struct {
	accountStore db.AccountStore
	accountQueue queue.Queue
}

func NewAccountResource(accountStore db.AccountStore, accountQeueue queue.Queue) *AccountResource {
	return &AccountResource{
		accountStore: accountStore,
		accountQueue: accountQeueue,
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
	var accountUpdateRequest AccountUpdateRequest

	err := json.Unmarshal(req.Body, &accountUpdateRequest)

	if err != nil {
		//return 500
	}

	accountId := mustGetAccountId(req)
	account, err := this.accountStore.Get(accountId)

	if err != nil {
		//return 500
	}

	if account == nil {
		//return 404
	}

	if accountUpdateRequest.Password != nil {

	}

	if accountUpdateRequest.Email != nil {

	}

	if accountUpdateRequest.NotificationsEnabled != nil {
		account.NotificationsEnabled = *accountUpdateRequest.NotificationsEnabled
	}

	err = this.accountStore.Update(*account)

	if err != nil {
		//return 500
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

func (this *AccountResource) CreateAccount(req Request) *Response {
	var accountCreateRequest AccountCreateRequest

	err := json.Unmarshal(req.Body, &accountCreateRequest)

	if err != nil {
		//return 500
	}

	//do validations

	account, err := this.accountStore.GetByEmail(accountCreateRequest.Email)

	if err != nil {
		//return 500
	}

	if account != nil {
		//return 400 email already taken
	}

	timeNow := int64(123) //dummy current time

	newAccount := db.Account{
		Id:                   "generated-uuid",
		Email:                accountCreateRequest.Email,
		Password:             accountCreateRequest.Password, //make sure this gets encrypted
		NotificationsEnabled: false,
		CreatedAt:            timeNow,
		UpdatedAt:            timeNow,
	}

	err = this.accountStore.Insert(newAccount)

	if err != nil {
		//return 500
	}

	accountResponse := this.toAccountResponse(newAccount)
	responseBody, err := json.Marshal(accountResponse)

	if err != nil {
		//return 500
	}

	return &Response{
		Body:   responseBody,
		Status: http.StatusOK,
	}
}

func (this *AccountResource) DeleteAccount(req Request) *Response {
	accountId := mustGetAccountId(req)
	account, err := this.accountStore.Get(accountId)

	if err != nil {
		//return 500
	}

	if account == nil {
		//return 404
	}

	err = this.accountStore.Delete(accountId)

	if err != nil {
		//return 500
	}

	accountDeletionMessage := &queue.AccountDeletionMessage{
		AccountId: accountId,
	}

	err = this.accountQueue.Push(accountDeletionMessage)

	if err != nil {
		//log error
	}

	return &Response{} //populate response
}

func (this *AccountResource) toAccountResponse(account db.Account) *AccountResponse {
	return nil
}
