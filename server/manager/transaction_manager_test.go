package manager

/*
func TestTransactionManagerGetsByRequest(t *testing.T) {
	t.Run("it succeeds", func(t *testing.T) {
		manager := transactionManagerMock()
		req := testRequest()
		req.ResourceId = testResourceId

		transaction := data.Transaction{Id: "tr-id-1"}

		manager.MockTransactionStore.On("Get", req.ResourceId, req.Account.Id).Return(&transaction, nil)

		res := manager.GetByRequest(req)
		assert.Equal(t, res.Data, &transaction)
		assert.NoError(t, res.Error)
	})

	t.Run("it fails if transaction doesn't exist", func(t *testing.T) {
		manager := transactionManagerMock()
		req := testRequest()
		req.ResourceId = testResourceId

		manager.MockTransactionStore.On("Get", req.ResourceId, req.Account.Id).Return(nil, nil)

		res := manager.GetByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidTransactionId)
	})
}

func TestTransactionManagerGetsAllByRequest(t *testing.T) {
	t.Run("it succeeds with user-given values", func(t *testing.T) {
		manager := transactionManagerMock()
		query := rest.TransactionQuery{
			CreatedBefore: types.Int64Ptr(4000),
			CreatedAfter:  types.Int64Ptr(3000),
			AmountGte:     types.Float64Ptr(123.45),
			AmountLte:     types.Float64Ptr(678.90),
		}

		filter := getExpectedTransactionUserFilter(query)

		req := testRequest()
		req.Query = query

		transactions := []data.Transaction{{Id: "tr-id-1"}}

		manager.MockTransactionStore.On("GetBy", req.Account.Id, filter).Return(transactions, nil)
		manager.MockClock.On("DateFromStamp", *query.CreatedAfter).Return(testDate)
		manager.MockClock.On("DateFromStamp", *query.CreatedBefore).Return(testDate)

		res := manager.GetAllByRequest(req)
		assert.Equal(t, res.Data, transactions)
		assert.NoError(t, res.Error)
	})

	t.Run("it succeeds with default values", func(t *testing.T) {
		manager := transactionManagerMock()
		query := rest.TransactionQuery{}

		filter := getExpectedTransactionDefaultFilter()

		req := testRequest()
		req.Query = query

		transactions := []data.Transaction{{Id: "tr-id-1"}}

		manager.MockTransactionStore.On("GetBy", req.Account.Id, filter).Return(transactions, nil)

		res := manager.GetAllByRequest(req)
		assert.Equal(t, res.Data, transactions)
		assert.NoError(t, res.Error)
	})
}

func TestTransactionManagerCreatesByRequest(t *testing.T) {
	t.Run("it succeeds", func(t *testing.T) {
		manager := transactionManagerMock()
		req := testRequest()
		body := rest.TransactionCreateBody{
			Amount: 15.67,
		}
		req.Body = body

		expected := getExpectedCreatedTransaction(body, req.Account.Id)

		d := data.NewDate(body.Month, body.Day, body.Year)

		manager.MockClock.On("IsDateValid", d).Return(true)
		manager.MockClock.On("Now").Return(testTimestamp)
		manager.MockUid.On("GetId").Return(testUuid)
		manager.MockTransactionStore.On("Insert", expected).Return(nil)

		res := manager.CreateByRequest(req)
		assert.NoError(t, res.Error)
	})

	t.Run("it fails if date is invalid", func(t *testing.T) {
		manager := transactionManagerMock()
		req := testRequest()
		body := rest.TransactionCreateBody{
			Amount: 15.67,
		}
		req.Body = body

		d := data.NewDate(body.Month, body.Day, body.Year)

		manager.MockClock.On("IsDateValid", d).Return(false)

		res := manager.CreateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidDate)
	})

	t.Run("it fails if note is too long", func(t *testing.T) {
		manager := transactionManagerMock()
		req := testRequest()

		noteValue := genString(config.LimitMaxTransactionNoteChars + 5)

		body := rest.TransactionCreateBody{
			Note: types.StringPtr(noteValue),
		}

		req.Body = body

		d := data.NewDate(body.Month, body.Day, body.Year)

		manager.MockClock.On("IsDateValid", d).Return(true)

		res := manager.CreateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidTransactionNote)
	})

	t.Run("it fails if category doesn't exist", func(t *testing.T) {
		manager := transactionManagerMock()
		req := testRequest()

		body := rest.TransactionCreateBody{
			CategoryId: types.StringPtr("cat-id"),
		}

		req.Body = body

		d := data.NewDate(body.Month, body.Day, body.Year)

		manager.MockClock.On("IsDateValid", d).Return(true)
		manager.MockCategoryStore.On("Get", *body.CategoryId, req.Account.Id).Return(nil, nil)

		res := manager.CreateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryId)
	})
}

func TestTransactionManagerUpdatesByRequest(t *testing.T) {
	t.Run("it succeeds", func(t *testing.T) {
		manager := transactionManagerMock()
		req := testRequest()
		body := rest.TransactionUpdateBody{
			Amount: 15.67,
		}
		req.ResourceId = testResourceId
		req.Body = body

		existing := data.Transaction{Amount: 37.80}

		expected := getExpectedTransactionUpdate(body)

		d := data.NewDate(body.Month, body.Day, body.Year)

		manager.MockTransactionStore.On("Get", req.ResourceId, req.Account.Id).Return(&existing, nil)
		manager.MockClock.On("IsDateValid", d).Return(true)
		manager.MockClock.On("Now").Return(testTimestamp)
		manager.MockTransactionStore.On("Update", req.ResourceId, req.Account.Id, expected, testTimestamp).Return(true, nil)

		res := manager.UpdateByRequest(req)
		assert.NoError(t, res.Error)
	})

	t.Run("it fails if transaction wasn't updated", func(t *testing.T) {
		manager := transactionManagerMock()
		req := testRequest()
		body := rest.TransactionUpdateBody{
			Amount: 15.67,
		}
		req.ResourceId = testResourceId
		req.Body = body

		existing := data.Transaction{Amount: 37.80}

		expected := getExpectedTransactionUpdate(body)

		d := data.NewDate(body.Month, body.Day, body.Year)

		manager.MockTransactionStore.On("Get", req.ResourceId, req.Account.Id).Return(&existing, nil)
		manager.MockClock.On("IsDateValid", d).Return(true)
		manager.MockClock.On("Now").Return(testTimestamp)
		manager.MockTransactionStore.On("Update", req.ResourceId, req.Account.Id, expected, testTimestamp).Return(false, nil)

		res := manager.UpdateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidTransactionId)
	})

	t.Run("it fails if date is invalid", func(t *testing.T) {
		manager := transactionManagerMock()
		req := testRequest()
		body := rest.TransactionUpdateBody{
			Amount: 15.67,
		}
		req.ResourceId = testResourceId
		req.Body = body

		existing := data.Transaction{Amount: 37.80}

		d := data.NewDate(body.Month, body.Day, body.Year)

		manager.MockTransactionStore.On("Get", req.ResourceId, req.Account.Id).Return(&existing, nil)
		manager.MockClock.On("IsDateValid", d).Return(false)

		res := manager.UpdateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidDate)
	})

	t.Run("it fails if note is too long", func(t *testing.T) {
		manager := transactionManagerMock()
		req := testRequest()
		note := genString(config.LimitMaxTransactionNoteChars + 5)
		body := rest.TransactionUpdateBody{
			Note:   types.StringPtr(note),
			Amount: 15.67,
		}
		req.ResourceId = testResourceId
		req.Body = body

		existing := data.Transaction{Amount: 37.80}

		d := data.NewDate(body.Month, body.Day, body.Year)

		manager.MockTransactionStore.On("Get", req.ResourceId, req.Account.Id).Return(&existing, nil)
		manager.MockClock.On("IsDateValid", d).Return(true)

		res := manager.UpdateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidTransactionNote)
	})

	t.Run("it fails if category doesn't exist", func(t *testing.T) {
		manager := transactionManagerMock()
		req := testRequest()
		body := rest.TransactionUpdateBody{
			CategoryId: types.StringPtr("new-category-id"),
			Amount:     15.67,
		}
		req.ResourceId = testResourceId
		req.Body = body

		existing := data.Transaction{Amount: 37.80}

		d := data.NewDate(body.Month, body.Day, body.Year)

		manager.MockTransactionStore.On("Get", req.ResourceId, req.Account.Id).Return(&existing, nil)
		manager.MockClock.On("IsDateValid", d).Return(true)
		manager.MockCategoryStore.On("Get", *body.CategoryId, req.Account.Id).Return(nil, nil)

		res := manager.UpdateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidCategoryId)
	})

	t.Run("it fails if transaction doesn't exist", func(t *testing.T) {
		manager := transactionManagerMock()
		req := testRequest()
		body := rest.TransactionUpdateBody{
			Amount: 15.67,
		}
		req.ResourceId = testResourceId
		req.Body = body

		manager.MockTransactionStore.On("Get", req.ResourceId, req.Account.Id).Return(nil, nil)

		res := manager.UpdateByRequest(req)
		assert.ErrorIs(t, res.Error, rest.ErrInvalidTransactionId)
	})
}

func TestTransactionManagerDeletesByRequest(t *testing.T) {
	t.Run("it succeeds", func(t *testing.T) {
		manager := transactionManagerMock()
		req := testRequest()
		req.ResourceId = testResourceId

		manager.MockTransactionStore.On("Delete", req.ResourceId, req.Account.Id).Return(true, nil)

		res := manager.DeleteByRequest(req)
		assert.NoError(t, res.Error)
	})

	t.Run("it fails if transaction wasn't deleted", func(t *testing.T) {
		manager := transactionManagerMock()
		req := testRequest()
		req.ResourceId = testResourceId

		manager.MockTransactionStore.On("Delete", req.ResourceId, req.Account.Id).Return(false, nil)

		res := manager.DeleteByRequest(req)
		assert.Error(t, res.Error, rest.ErrInvalidTransactionId)
	})
}

func getExpectedTransactionUserFilter(q rest.TransactionQuery) data.TransactionFilter {
	return data.TransactionFilter{
		Before:      testDate,
		After:       testDate,
		GreaterThan: *q.AmountGte,
		LessThan:    *q.AmountLte,
	}
}

func getExpectedTransactionDefaultFilter() data.TransactionFilter {
	return data.TransactionFilter{
		Before:      data.NewDate(1, 1, 1902),
		After:       data.NewDate(1, 1, math.MaxInt),
		GreaterThan: math.MaxFloat64,
		LessThan:    0,
	}
}

func getExpectedCreatedTransaction(body rest.TransactionCreateBody, accountId string) data.Transaction {
	return data.Transaction{
		Id:         testUuid,
		AccountId:  accountId,
		CategoryId: body.CategoryId,
		Note:       body.Note,
		Type:       body.Type,
		Amount:     body.Amount,
		Month:      body.Month,
		Day:        body.Day,
		Year:       body.Year,
		CreatedAt:  testTimestamp,
		UpdatedAt:  testTimestamp,
	}
}

func getExpectedTransactionUpdate(body rest.TransactionUpdateBody) data.TransactionUpdate {
	return data.TransactionUpdate{
		CategoryId: body.CategoryId,
		Note:       body.Note,
		Type:       body.Type,
		Amount:     body.Amount,
		Month:      body.Month,
		Day:        body.Day,
		Year:       body.Year,
	}
}
*/
