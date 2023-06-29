package api

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/store"
)

type RestAPI struct {
	store store.Store
	clock clock.Clock
}

func (api *RestAPI) makeNewCategory(accountId string, req data.CategoryCreateRequest) data.Category {
	now := api.clock.Now()

	return data.Category{
		Id:        "", //Update
		AccountId: accountId,
		Name:      req.Name,
		Color:     req.Color,
		UpdatedAt: now,
		CreatedAt: now,
	}
}

func getAccount(req *data.RestRequest) data.Account {
	return req.Meta["curr_account"].(data.Account)
}

func (api *RestAPI) getCategory(req *data.RestRequest, res *data.RestResponse) data.Category {
	id := req.UrlParams["category_id"].(string)
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

	account.Name = updateReq.Name
	account.UpdatedAt = api.clock.Now()

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

func (api *RestAPI) GetCategory(req *data.RestRequest, res *data.RestResponse) {
	category := api.getCategory(req, res)
	if isErrorResponse(res.Status) {
		return
	}

	buildOKResponse(res, category)
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

	category.Color = updateReq.Color
	category.Name = updateReq.Name
	category.UpdatedAt = api.clock.Now()

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
