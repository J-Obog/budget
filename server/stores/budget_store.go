package store

import "github.com/J-Obog/paidoff/models"

type BudgetStore interface {
	Get(id string) (error, *models.Budget)
	Insert(budget models.Budget) error
	Update(budget models.Budget) error
	Delete(id string) error
}
