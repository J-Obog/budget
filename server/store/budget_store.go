package store

import "github.com/J-Obog/paidoff/data"

type BudgetStore interface {
	Get(id string) (*data.Budget, error)
	Insert(budget data.Budget) error
	Update(budget data.Budget) error
	Delete(id string) error
}
