package store

import "github.com/J-Obog/paidoff/models"

type AccountStore interface {
	Get(id string) *models.Account
	Update(updatedAccount models.Account)
	Delete(id string)
}
