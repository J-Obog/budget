package api

import (
	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

type TransactionAPI struct {
	transactionManager *manager.TransactionManager
	categoryManager    *manager.CategoryManager
}

func getTransactionId(req *rest.Request) string {
	return ""
}

func (api *TransactionAPI) Get(req *rest.Request) *rest.Response {
	id := getTransactionId(req)

	transaction, err := api.transactionManager.Get(id, req.Account.Id)
	if err != nil {
		return rest.Err(err)
	}

	if transaction == nil {
		return rest.Err(rest.ErrInvalidTransactionId)
	}

	return rest.Ok(transaction)
}

// TODO: implement
func (api *TransactionAPI) GetAll(req *rest.Request) *rest.Response {
	return nil
}

func (api *TransactionAPI) Create(req *rest.Request) *rest.Response {
	accountId := req.Account.Id

	body, err := rest.ParseBody[rest.TransactionCreateBody](req.Body)
	if err != nil {
		return rest.Err(err)
	}

	if err := api.validateCreate(body, accountId); err != nil {
		return rest.Err(err)
	}

	transaction, err := api.transactionManager.Create(accountId, body)
	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(transaction)
}

func (api *TransactionAPI) Update(req *rest.Request) *rest.Response {
	id := getTransactionId(req)
	accountId := req.Account.Id

	body, err := rest.ParseBody[rest.TransactionUpdateBody](req.Body)
	if err != nil {
		return rest.Err(err)
	}

	existing, err := api.transactionManager.Get(id, accountId)
	if err != nil {
		return rest.Err(err)
	}

	if err := api.validateUpdate(existing, body); err != nil {
		return rest.Err(err)
	}

	ok, err := api.transactionManager.Update(existing, body)
	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidTransactionId)
	}

	return rest.Ok(existing)
}

func (api *TransactionAPI) Delete(req *rest.Request) *rest.Response {
	id := getTransactionId(req)
	accountId := req.Account.Id

	ok, err := api.transactionManager.Delete(id, accountId)
	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidTransactionId)
	}

	return rest.Success()
}

// TODO: look into validation for budget type
func (api *TransactionAPI) validateUpdate(existing *data.Transaction, body rest.TransactionUpdateBody) error {
	if existing == nil {
		return rest.ErrInvalidTransactionId
	}

	if err := isDateValid(body.Month, body.Day, body.Year); err != nil {
		return err
	}

	if body.Note != nil {
		if existing.Note == nil || (*existing.Note != *body.Note) {
			if err := api.checkNote(*body.Note); err != nil {
				return err
			}
		}
	}

	if body.CategoryId != nil {
		if existing.CategoryId == nil || (*existing.CategoryId != *body.CategoryId) {
			if err := api.checkCategoryExists(*body.CategoryId, existing.AccountId); err != nil {
				return err
			}
		}
	}

	return nil
}

func (api *TransactionAPI) validateCreate(body rest.TransactionCreateBody, accountId string) error {
	if err := isDateValid(body.Month, body.Day, body.Year); err != nil {
		return err
	}

	if body.Note != nil {
		if err := api.checkNote(*body.Note); err != nil {
			return err
		}
	}

	if body.CategoryId != nil {
		if err := api.checkCategoryExists(*body.CategoryId, accountId); err != nil {
			return err
		}
	}

	return nil
}

func (api *TransactionAPI) checkNote(note string) error {
	if len(note) > config.LimitMaxTransactionNoteChars {
		return rest.ErrInvalidTransactionNote
	}

	return nil
}

func (api *TransactionAPI) checkCategoryExists(categoryId string, accountId string) error {
	category, err := api.categoryManager.Get(categoryId, accountId)
	if err != nil {
		return err
	}

	if category == nil {
		return rest.ErrInvalidCategoryId
	}

	return nil
}
