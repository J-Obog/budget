package store

import "github.com/J-Obog/paidoff/data"

type TransactionStore interface {
	Get(id string) (error, *data.Transaction)
	Insert(transaction data.Transaction) error
	Update(transaction data.Transaction) error
	Delete(id string) error
}
