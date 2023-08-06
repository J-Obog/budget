package server

import "github.com/J-Obog/paidoff/rest"

type RouteHandler func(*rest.Request) *rest.Response

type Server interface {
	Start()
	Stop() error
	RegisterRoute(method string, url string, rh RouteHandler)
}
