package store

import "github.com/J-Obog/paidoff/models"

type AccountStore interface {
	Get(id string) (error *models.Account)
	Insert(account models.Account) error
	Update(account models.Account) error
	Delete(id string) error
}
