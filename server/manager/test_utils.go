package manager

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/mocks"
	"github.com/J-Obog/paidoff/rest"
)

const (
	testTimestamp  = int64(1234564321)
	testUuid       = "test-uuid-1"
	testResourceId = "test-resource-id"
)

type AccountManagerMock struct {
	AccountManager
	MockAccountStore *mocks.AccountStore
	MockClock        *mocks.Clock
}

func accountManagerMock() *AccountManagerMock {
	mockedAccountStore := new(mocks.AccountStore)
	mockedClock := new(mocks.Clock)

	return &AccountManagerMock{
		AccountManager: AccountManager{
			store: mockedAccountStore,
			clock: mockedClock,
		},
		MockAccountStore: mockedAccountStore,
		MockClock:        mockedClock,
	}
}

type CategoryManagerMock struct {
	CategoryManager
	MockCategoryStore *mocks.CategoryStore
	MockBudgetStore   *mocks.BudgetStore
	MockClock         *mocks.Clock
	MockUid           *mocks.UIDProvider
	MockQueue         *mocks.Queue
}

func categoryManagerMock() *CategoryManagerMock {
	mockedCategoryStore := new(mocks.CategoryStore)
	mockedBudgetStore := new(mocks.BudgetStore)
	mockedClock := new(mocks.Clock)
	mockedUid := new(mocks.UIDProvider)
	mockedQ := new(mocks.Queue)

	return &CategoryManagerMock{
		CategoryManager: CategoryManager{
			store:       mockedCategoryStore,
			budgetStore: mockedBudgetStore,
			clock:       mockedClock,
			uid:         mockedUid,
			queue:       mockedQ,
		},
		MockCategoryStore: mockedCategoryStore,
		MockBudgetStore:   mockedBudgetStore,
		MockClock:         mockedClock,
		MockUid:           mockedUid,
		MockQueue:         mockedQ,
	}
}

func testRequest() *rest.Request {
	return &rest.Request{
		Account: &data.Account{
			Id: "some-account-id",
		},
	}
}

func genString(length int) string {
	var s string

	for i := 0; i < length; i++ {
		s += "C"
	}

	return s
}
