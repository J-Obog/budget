package server

import (
	"net/http"

	"github.com/J-Obog/paidoff/api"
)

func RegisterRoutes(s Server, apiSvc *api.ApiService) {
	s.RegisterRoute(http.MethodGet, "/account", apiSvc.AccountAPI.Get)
	s.RegisterRoute(http.MethodPut, "/account", apiSvc.AccountAPI.Update)
	s.RegisterRoute(http.MethodDelete, "/account", apiSvc.AccountAPI.Delete)

	s.RegisterRoute(http.MethodGet, "/category", apiSvc.CategoryAPI.GetAll)
	s.RegisterRoute(http.MethodPost, "/category", apiSvc.CategoryAPI.Create)
	s.RegisterRoute(http.MethodGet, "/category/:categoryId", apiSvc.CategoryAPI.Get)
	s.RegisterRoute(http.MethodPut, "/category/:categoryId", apiSvc.CategoryAPI.Update)
	s.RegisterRoute(http.MethodDelete, "/category/:categoryId", apiSvc.CategoryAPI.Delete)

	s.RegisterRoute(http.MethodGet, "/transaction", apiSvc.TransactionAPI.GetByQuery)
	s.RegisterRoute(http.MethodPost, "/transaction", apiSvc.TransactionAPI.Create)
	s.RegisterRoute(http.MethodGet, "/transaction/:transactionId", apiSvc.TransactionAPI.Get)
	s.RegisterRoute(http.MethodPut, "/transaction/:transactionId", apiSvc.TransactionAPI.Update)
	s.RegisterRoute(http.MethodDelete, "/transaction/:transactionId", apiSvc.TransactionAPI.Delete)

	s.RegisterRoute(http.MethodGet, "/budget/:periodMonth/:periodYear", apiSvc.BudgetAPI.GetByPeriod)
	s.RegisterRoute(http.MethodPost, "/budget", apiSvc.BudgetAPI.Create)
	s.RegisterRoute(http.MethodGet, "/budget/:budgetId", apiSvc.BudgetAPI.Get)
	s.RegisterRoute(http.MethodPut, "/budget/:budgetId", apiSvc.BudgetAPI.Update)
	s.RegisterRoute(http.MethodDelete, "/budget/:budgetId", apiSvc.BudgetAPI.Delete)
}
