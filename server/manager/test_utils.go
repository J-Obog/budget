package manager

import (
	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/mocks"
	"github.com/J-Obog/paidoff/rest"
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

func categoryDelete() *rest.Request {
	return &rest.Request{
		Account:    testAccount(),
		ResourceId: "some-id",
	}
}

func testBudget() *data.Budget {
	return &data.Budget{
		Id: "some-id",
	}
}

func categoryUpdateNameTooShort() *rest.Request {
	body := rest.CategoryUpdateBody{}
	body.Color = 123
	body.Name = ""
	return &rest.Request{
		Account: testAccount(),
		Body:    body,
	}
}

func categoryUpdateNameTooLong() *rest.Request {
	body := rest.CategoryUpdateBody{}
	body.Color = 123
	body.Name = ""

	for i := 0; i < config.LimitMaxCategoryNameChars+20; i++ {
		body.Name += "F"
	}

	return &rest.Request{
		Account: testAccount(),
		Body:    body,
	}
}

func categoryUpdate() *rest.Request {
	body := rest.CategoryUpdateBody{}
	body.Color = 123
	body.Name = "Some other name"
	return &rest.Request{
		Account: testAccount(),
		Body:    body,
	}
}

func testAccount() *data.Account {
	return &data.Account{
		Id: "some-uuid",
	}
}

func testCategory() *data.Category {
	return &data.Category{
		Id:   "cat-id-1",
		Name: "Name1",
	}
}

func categoryCreateNameTooLong() *rest.Request {
	body := rest.CategoryCreateBody{}
	body.Name = ""

	for i := 0; i < config.LimitMaxCategoryNameChars+20; i++ {
		body.Name += "F"
	}

	return &rest.Request{
		Account: testAccount(),
		Body:    body,
	}
}

func categoryCreateNameTooShort() *rest.Request {
	body := rest.CategoryCreateBody{}
	body.Name = ""

	return &rest.Request{
		Account: testAccount(),
		Body:    body,
	}
}

func categoryCreate() *rest.Request {
	body := rest.CategoryCreateBody{}
	body.Color = 123
	body.Name = "Somename"

	return &rest.Request{
		Account: testAccount(),
		Body:    body,
	}
}

func categoryGet() *rest.Request {
	return &rest.Request{
		Account:    testAccount(),
		ResourceId: "cat-id-1",
	}
}

func accountUpdate() *rest.Request {
	body := rest.AccountUpdateBody{}
	body.Name = "John Doe"

	return &rest.Request{
		Account: testAccount(),
		Body:    body,
	}
}

func accountUpdateNameTooShort() *rest.Request {
	body := rest.AccountUpdateBody{}
	body.Name = ""

	return &rest.Request{
		Account: testAccount(),
		Body:    body,
	}
}

func accountUpdateNameTooLong() *rest.Request {
	name := ""
	for i := 0; i < config.LimitMaxAccountNameChars+5; i++ {
		name += "F"
	}

	body := rest.AccountUpdateBody{}
	body.Name = string(name)

	return &rest.Request{
		Account: testAccount(),
		Body:    body,
	}
}

func accountDelete() *rest.Request {
	return &rest.Request{
		Account: testAccount(),
	}
}
