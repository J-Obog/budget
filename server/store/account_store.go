package store

import (
	"github.com/J-Obog/paidoff/data"
)

type AccountStore interface {
	Get(id string) (*data.Account, error)
	Insert(account data.Account) error
	Update(account data.Account) error
	Delete(id string) error
}
