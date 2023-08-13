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
	s.RegisterRoute(http.MethodGet, "/category/:id", apiSvc.CategoryAPI.Get)
	s.RegisterRoute(http.MethodPut, "/category/:id", apiSvc.CategoryAPI.Update)
	s.RegisterRoute(http.MethodDelete, "/category/:id", apiSvc.CategoryAPI.Delete)

	s.RegisterRoute(http.MethodGet, "/transaction", apiSvc.TransactionAPI.Filter)
	s.RegisterRoute(http.MethodPost, "/transaction", apiSvc.TransactionAPI.Create)
	s.RegisterRoute(http.MethodGet, "/transaction/:id", apiSvc.TransactionAPI.Get)
	s.RegisterRoute(http.MethodPut, "/transaction/:id", apiSvc.TransactionAPI.Update)
	s.RegisterRoute(http.MethodDelete, "/transaction/:id", apiSvc.TransactionAPI.Delete)

	s.RegisterRoute(http.MethodGet, "/budget", apiSvc.BudgetAPI.Filter)
	s.RegisterRoute(http.MethodPost, "/budget", apiSvc.BudgetAPI.Create)
	s.RegisterRoute(http.MethodGet, "/budget/:id", apiSvc.BudgetAPI.Get)
	s.RegisterRoute(http.MethodPut, "/budget/:id", apiSvc.BudgetAPI.Update)
	s.RegisterRoute(http.MethodDelete, "/budget/:id", apiSvc.BudgetAPI.Delete)
}
