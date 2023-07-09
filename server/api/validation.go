package api

import (
	"github.com/J-Obog/paidoff/manager"
)

const (
	MaxAccountNameChars       int = 150
	MinAccountNameChars       int = 1
	MaxBudgetDescriptionChars int = 200
)

func checkBudgetDesciption(description string) error {
	return nil
}

func checkAccountName(name string) error {
	return nil
}

func checkCategoryExists(categoryId string, accountId string, categoryManager *manager.CategoryManager) error {
	cat, err := categoryManager.Get(categoryId)
	if err != nil {
		return err
	}
	if cat == nil || cat.AccountId != accountId {
		return err
	}

	return nil
}

func checkCategoryNotUsed(categoryId string, month int, year int, budgetManager *manager.BudgetManager) error {
	return nil
}

func checkCategoryNameNotUsed(name string, accountId string, categoryManager *manager.CategoryManager) error {
	return nil
}

func checkDate(month int, day int, year int) error {
	return nil
}
