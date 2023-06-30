package api

import "github.com/J-Obog/paidoff/data"

type Filter[T interface{}] struct {
	checks []func(val T) bool
}

func NewFilter[T interface{}]() *Filter[T] {
	return &Filter[T]{
		checks: make([]func(val T) bool, 0),
	}
}

func (f *Filter[T]) AddCheck(checkFn func(val T) bool) {
	f.checks = append(f.checks, checkFn)
}

func (f *Filter[T]) runChecks(item T) bool {
	for _, check := range f.checks {
		res := check(item)
		if !res {
			return false
		}
	}

	return true
}

func (f *Filter[T]) Filter(items []T) []T {
	filtered := make([]T, 0)

	for _, item := range items {
		res := f.runChecks(item)
		if res {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

func filterTransaction(q data.TransactionQuery) func(data.Transaction) bool {
	return func(t data.Transaction) bool {
		if (q.AmountGte != nil) && (t.Amount < *q.AmountGte) {
			return false
		}
		if (q.AmountLte != nil) && (t.Amount > *q.AmountGte) {
			return false
		}

		if (q.CreatedAfter != nil) && (t.CreatedAt <= *q.CreatedAfter) {
			return false
		}

		if (q.CreatedBefore != nil) && (t.CreatedAt <= *q.CreatedAfter) {
			return false
		}

		// check if transaction has category

		// check budget type??
		return true
	}
}
