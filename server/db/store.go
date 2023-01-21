package db

type CategoryStore interface {
	Get(id string) (*Category, error)
	GetAll() ([]Category, error)
}

type GoalStore interface {
	Get(id string) (*Goal, error)
	GetAll(accountId string) ([]Goal, error)
	Update(goal Goal) error
	Insert(goal Goal) error
	Delete(id string) error
}

type AccountStore interface {
	Get(id string) (*Account, error)
	GetByEmail(email string) (*Account, error)
	Update(account Account) error
	Insert(account Account) error
	Delete(id string) error
}
