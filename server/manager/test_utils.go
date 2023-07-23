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

func testAccount() *data.Account {
	return &data.Account{
		Id: "some-uuid",
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
