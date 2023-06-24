package store

import "github.com/J-Obog/paidoff/data"

type BudgetStore interface {
	Get(id string) (error, *data.Budget)
	Insert(budget data.Budget) error
	Update(budget data.Budget) error
	Delete(id string) error
}
