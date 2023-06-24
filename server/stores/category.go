package store

import "github.com/J-Obog/paidoff/models"

type CategoryStore interface {
	Get(id string) (error, *models.Category)
	Insert(category models.Category) error
	Update(account models.Account) error
	Delete(id string) error
}
