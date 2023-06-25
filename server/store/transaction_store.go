package store

import "github.com/J-Obog/paidoff/data"

type TransactionStore interface {
	Get(id string) (*data.Transaction, error)
	Insert(transaction data.Transaction) error
	Update(transaction data.Transaction) error
	Delete(id string) error
}
