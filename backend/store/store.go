package store

type Store[T any] interface {
	Get(id string) (*T, error)
	GetByAccount(accountId string) ([]T, error)
	Insert(entity T) error
	Update(entity T) (bool, error)
	Delete(id string) (bool, error)
}
