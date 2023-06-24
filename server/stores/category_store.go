package store

import "github.com/J-Obog/paidoff/data"

type CategoryStore interface {
	Get(id string) (error, *data.Category)
	Insert(category data.Category) error
	Update(account data.Account) error
	Delete(id string) error
}
