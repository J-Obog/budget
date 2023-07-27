package manager

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/mocks"
	"github.com/J-Obog/paidoff/rest"
)

const (
	testTimestamp = int64(1234564321)
	testUuid      = "test-uuid-1"

	testCategoryId    = "test-category-id"
	testTransactionId = "test-transaction-id"
	testBudgetId      = "test-budget-id"
	testAccountId     = "test-account-id"
)

var (
	testDate = data.NewDate(6, 15, 2023)

	testLongString  = genString(1000)
	testShortString = ""
)

func testRequest() *rest.Request {
	return &rest.Request{
		Account: &data.Account{
			Id: testAccountId,
		},
	}
}

func genString(length int) string {
	var s string

	for i := 0; i < length; i++ {
		s += "F"
	}

	return s
}

func setRequestBody() []byte {
	return nil
}

func mockTimestamp(mockClock *mocks.Clock) {
	mockClock.On("Now").Return(testTimestamp)
}

func mockUuid(mockUuidProvider *mocks.UIDProvider) {
	mockUuidProvider.On("GetId").Return(testUuid)
}

func categoryNameTaken(mockCategoryStore *mocks.CategoryStore, accountId string, name string) {
	mockCategoryStore.On("GetByName", accountId, name).Return(&data.Category{
		Id: "some-existing-category",
	}, nil)
}

func categoryNameAvailable(mockCategoryStore *mocks.CategoryStore, accountId string, name string) {
	mockCategoryStore.On("GetByName", accountId, name).Return(nil, nil)
}
