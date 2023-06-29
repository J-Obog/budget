package api

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/store"
	"github.com/J-Obog/paidoff/uid"
)

type RestAPI struct {
	store       store.Store
	clock       clock.Clock
	uidProvider uid.UIDProvider
}

func (api *RestAPI) checkCategoryExists(accountId string, categoryId string, res *data.RestResponse) {
	category, err := api.store.GetCategory(categoryId)

	if err != nil {
		buildServerError(res, err)
		return
	}

	if category == nil || category.AccountId != accountId {
		buildForbiddenError(res)
		return
	}
}

func (api *RestAPI) getCategory(req *data.RestRequest, res *data.RestResponse) data.Category {
	id := req.UrlParams["categoryId"].(string)
	category, err := api.store.GetCategory(id)

	if err != nil {
		buildServerError(res, err)
		return data.Category{}
	}

	if category == nil {
		buildNotFoundError(res)
		return data.Category{}
	}

	if category.AccountId != getAccount(req).Id {
		buildForbiddenError(res)
		return data.Category{}
	}

	return *category
}

func (api *RestAPI) getTransaction(req *data.RestRequest, res *data.RestResponse) data.Transaction {
	id := req.UrlParams["transactionId"].(string)
	transaction, err := api.store.GetTransaction(id)

	if err != nil {
		buildServerError(res, err)
		return data.Transaction{}
	}

	if transaction == nil {
		buildNotFoundError(res)
		return data.Transaction{}
	}

	if transaction.AccountId != getAccount(req).Id {
		buildForbiddenError(res)
		return data.Transaction{}
	}

	return *transaction
}

// Accounts

func (api *RestAPI) GetAccount(req *data.RestRequest, res *data.RestResponse) {
	account := getAccount(req)
	buildOKResponse(res, account)
}

func (api *RestAPI) UpdateAccount(req *data.RestRequest, res *data.RestResponse) {
	account := getAccount(req)

	updateReq, err := FromJSON[data.AccountUpdateRequest](req.Body)
	if err != nil {
		buildServerError(res, err)
		return
	}

	now := api.clock.Now()
	buildUpdatedAccount(now, updateReq, &account)

	if err := api.store.UpdateAccount(account); err != nil {
		buildServerError(res, err)
		return
	}

	buildOKResponse(res, nil)
}

func (api *RestAPI) DeleteAccount(req *data.RestRequest, res *data.RestResponse) {
	account := getAccount(req)
	account.IsDeleted = true

	if err := api.store.UpdateAccount(account); err != nil {
		buildServerError(res, err)
		return
	}

	buildOKResponse(res, account)
}

// Categories
func (api *RestAPI) GetCategory(req *data.RestRequest, res *data.RestResponse) {
	category := api.getCategory(req, res)
	if isErrorResponse(res.Status) {
		return
	}

	buildOKResponse(res, category)
}

func (api *RestAPI) GetCategories(req *data.RestRequest, res *data.RestResponse) {
	accountId := getAccount(req).Id
	categories, err := api.store.GetCategories(accountId)

	if err != nil {
		buildServerError(res, err)
		return
	}

	buildOKResponse(res, categories)
}

func (api *RestAPI) CreateCategory(req *data.RestRequest, res *data.RestResponse) {
	createReq, err := FromJSON[data.CategoryCreateRequest](req.Body)
	if err != nil {
		buildServerError(res, err)
		return
	}

	accountId := getAccount(req).Id
	newCategory := api.makeNewCategory(accountId, createReq)

	if err := api.store.InsertCategory(newCategory); err != nil {
		buildServerError(res, err)
		return
	}

	buildOKResponse(res, newCategory)
}

func (api *RestAPI) UpdateCategory(req *data.RestRequest, res *data.RestResponse) {
	category := api.getCategory(req, res)

	if isErrorResponse(res.Status) {
		return
	}

	updateReq, err := FromJSON[data.CategoryUpdateRequest](req.Body)
	if err != nil {
		buildServerError(res, err)
		return
	}

	now := api.clock.Now()
	buildUpdatedCategory(now, updateReq, &category)

	if err := api.store.UpdateCategory(category); err != nil {
		buildServerError(res, err)
		return
	}

	buildOKResponse(res, category)
}

func (api *RestAPI) DeleteCategory(req *data.RestRequest, res *data.RestResponse) {
	category := api.getCategory(req, res)

	if isErrorResponse(res.Status) {
		return
	}

	if err := api.store.DeleteCategory(category.Id); err != nil {
		buildServerError(res, err)
		return
	}

	buildOKResponse(res, category)
}

// Transactions
func (api *RestAPI) GetTransaction(req *data.RestRequest, res *data.RestResponse) {
	transaction := api.getTransaction(req, res)
	if isErrorResponse(res.Status) {
		return
	}

	buildOKResponse(res, transaction)
}

func (api *RestAPI) CreateTransaction(req *data.RestRequest, res *data.RestResponse) {
	createReq, err := FromJSON[data.TransactionCreateRequest](req.Body)
	if err != nil {
		buildServerError(res, err)
		return
	}

	accountId := getAccount(req).Id

	if createReq.CategoryId != nil {
		api.checkCategoryExists(accountId, *createReq.CategoryId, res)
		if isErrorResponse(res.Status) {
			return
		}
	}

	id := api.uidProvider.GetId()
	now := api.clock.Now()
	newTransaction := makeNewTransaction(id, accountId, now, createReq)

	if err := api.store.InsertTransaction(newTransaction); err != nil {
		buildServerError(res, err)
		return
	}

	buildOKResponse(res, newTransaction)
}

func (api *RestAPI) UpdateTransaction(req *data.RestRequest, res *data.RestResponse) {
	transaction := api.getTransaction(req, res)

	if isErrorResponse(res.Status) {
		return
	}

	updateReq, err := FromJSON[data.TransactionUpdateRequest](req.Body)
	if err != nil {
		buildServerError(res, err)
		return
	}

	if updateReq.CategoryId != nil {
		accountId := getAccount(req).Id
		api.checkCategoryExists(accountId, *updateReq.CategoryId, res)
		if isErrorResponse(res.Status) {
			return
		}
	}

	now := api.clock.Now()
	buildUpdatedTransaction(now, updateReq, &transaction)

	if err := api.store.UpdateTransaction(transaction); err != nil {
		buildServerError(res, err)
		return
	}

	buildOKResponse(res, transaction)
}

func (api *RestAPI) DeleteTransaction(req *data.RestRequest, res *data.RestResponse) {
	transaction := api.getTransaction(req, res)

	if isErrorResponse(res.Status) {
		return
	}

	if err := api.store.DeleteCategory(transaction.Id); err != nil {
		buildServerError(res, err)
		return
	}

	buildOKResponse(res, transaction)
}
