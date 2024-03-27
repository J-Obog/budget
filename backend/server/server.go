package server

import "github.com/J-Obog/paidoff/types/rest"

type RouteHandler func(*rest.Request, *rest.Response)

type Server interface {
	Start()
	Stop() error
	Handle(method string, url string, handler RouteHandler)
}
