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

var (
	testDate = data.NewDate(6, 15, 2023)
)

type TransactionManagerMock struct {
	TransactionManager
	MockTransactionStore *mocks.TransactionStore
	MockCategoryStore    *mocks.CategoryStore
	MockClock            *mocks.Clock
	MockUid              *mocks.UIDProvider
}

func transactionManagerMock() *TransactionManagerMock {
	mockedCategoryStore := new(mocks.CategoryStore)
	mockedTransactionStore := new(mocks.TransactionStore)
	mockedClock := new(mocks.Clock)
	mockedUid := new(mocks.UIDProvider)

	return &TransactionManagerMock{
		TransactionManager: TransactionManager{
			store:         mockedTransactionStore,
			categoryStore: mockedCategoryStore,
			clock:         mockedClock,
			uid:           mockedUid,
		},
		MockCategoryStore:    mockedCategoryStore,
		MockTransactionStore: mockedTransactionStore,
		MockClock:            mockedClock,
		MockUid:              mockedUid,
	}
}

type BugdetManagerMock struct {
	BudgetManager
	MockBudgetStore      *mocks.BudgetStore
	MockCategoryStore    *mocks.CategoryStore
	MockTransactionStore *mocks.TransactionStore
	MockClock            *mocks.Clock
	MockUid              *mocks.UIDProvider
}

func bugdetManagerMock() *BugdetManagerMock {
	mockedCategoryStore := new(mocks.CategoryStore)
	mockedBudgetStore := new(mocks.BudgetStore)
	mockedTransactionStore := new(mocks.TransactionStore)
	mockedClock := new(mocks.Clock)
	mockedUid := new(mocks.UIDProvider)

	return &BugdetManagerMock{
		BudgetManager: BudgetManager{
			store:            mockedBudgetStore,
			categoryStore:    mockedCategoryStore,
			transactionStore: mockedTransactionStore,
			clock:            mockedClock,
			uid:              mockedUid,
		},
		MockBudgetStore:      mockedBudgetStore,
		MockCategoryStore:    mockedCategoryStore,
		MockTransactionStore: mockedTransactionStore,
		MockClock:            mockedClock,
		MockUid:              mockedUid,
	}
}

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
