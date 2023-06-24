package store

import "github.com/J-Obog/paidoff/models"

type TransactionStore interface {
	Get(id string) (error, *models.Transaction)
	Insert(transaction models.Transaction) error
	Update(transaction models.Transaction) error
	Delete(id string) error
}
