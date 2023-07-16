package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
	"github.com/J-Obog/paidoff/uid"
)

type TransactionManager struct {
	store           store.TransactionStore
	categoryManager *CategoryManager
	clock           clock.Clock
	uid             uid.UIDProvider
}

type transactionValidateCommon struct {
	description *string
	month       int
	day         int
	year        int
	categoryId  string
	accountId   string
}

func (manager *TransactionManager) Get(id string) (*data.Transaction, error) {
	transaction, err := manager.store.Get(id)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (manager *TransactionManager) GetByRequest(req *rest.Request, res *rest.Response) {
	accountId := req.Account.Id
	id := req.Params.BudgetId()

	transaction := manager.getTransactionByAccount(res, id, accountId)

	if res.IsErr() {
		return
	}

	res.Ok(transaction)
}

func (manager *TransactionManager) GetAllByRequest(req *rest.Request, res *rest.Response) {
	query := req.Query.TransactionQuery()
	accountId := req.Account.Id

	transactions, err := manager.store.GetByAccount(accountId)
	if err != nil {
		res.ErrInternal(err)
		return
	}

	filtered := filter[data.Transaction](transactions, func(t *data.Transaction) bool {
		if query.CreatedBefore != nil && t.CreatedAt >= *query.CreatedBefore {
			return false
		}

		if query.CreatedAfter != nil && t.CreatedAt <= *query.CreatedAfter {
			return false
		}

		if query.AmountGte != nil && t.Amount < *query.AmountGte {
			return false
		}

		if query.AmountLte != nil && t.Amount > *query.AmountLte {
			return false
		}
		return true

	})

	res.Ok(filtered)
}

func (manager *TransactionManager) CreateByRequest(req *rest.Request, res *rest.Response) {
	body, err := req.Body.TransactionCreateBody()
	if err != nil {
		res.ErrBadRequest()
		return
	}

	validateCommon := transactionValidateCommon{
		description: body.Description,
		month:       body.Month,
		day:         body.Day,
		year:        body.Year,
		categoryId:  *body.CategoryId,
		accountId:   req.Account.Id,
	}

	manager.validate(res, validateCommon)
	if res.IsErr() {
		return
	}

	now := manager.clock.Now()
	id := manager.uid.GetId()

	newTransaction := newTransaction(body, id, req.Account.Id, now)

	if err := manager.store.Insert(newTransaction); err != nil {
		res.ErrInternal(err)
		return
	}

	res.Ok(nil)
}

func (manager *TransactionManager) UpdateByRequest(req *rest.Request, res *rest.Response) {
	accountId := req.Account.Id
	id := req.Params.TransactionId()

	transaction := manager.getTransactionByAccount(res, id, accountId)
	if res.IsErr() {
		return
	}

	body, err := req.Body.TransactionUpdateBody()
	if err != nil {
		res.ErrBadRequest()
		return
	}

	validateCommon := transactionValidateCommon{
		description: body.Description,
		month:       body.Month,
		day:         body.Day,
		year:        body.Year,
		categoryId:  *body.CategoryId,
		accountId:   req.Account.Id,
	}

	manager.validate(res, validateCommon)
	if res.IsErr() {
		return
	}

	now := manager.clock.Now()
	updateTransaction(body, transaction, now)

	if err = manager.store.Update(*transaction); err != nil {
		res.ErrInternal(err)
		return
	}

	res.Ok(nil)
}

func (manager *TransactionManager) DeleteByRequest(req *rest.Request, res *rest.Response) {
	accountId := req.Account.Id
	id := req.Params.TransactionId()

	manager.getTransactionByAccount(res, id, accountId)
	if res.IsErr() {
		return
	}

	if err := manager.store.Delete(id); err != nil {
		res.ErrInternal(err)
		return
	}

	res.Ok(nil)
}

func (manager *TransactionManager) getTransactionByAccount(
	res *rest.Response,
	id string,
	accountId string,
) *data.Transaction {
	transaction, err := manager.Get(id)

	if err != nil {
		res.ErrInternal(err)
		return nil
	}

	if transaction == nil || transaction.AccountId != accountId {
		res.ErrTransactionNotFound()
		return nil
	}

	return transaction
}

func (manager *TransactionManager) validate(res *rest.Response, validateCom transactionValidateCommon) {
	categoryId := validateCom.categoryId
	accountId := validateCom.accountId
	//month := validateCom.month
	//year := validateCom.year
	//day := validateCom.day

	//check if date is valid

	ok, err := manager.categoryManager.Exists(categoryId, accountId)
	if err != nil {
		res.ErrInternal(err)
		return
	}

	if !ok {
		res.ErrCategoryNotFound()
		return
	}

	//check description
}
