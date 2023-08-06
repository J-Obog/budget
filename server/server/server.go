package server

import "github.com/J-Obog/paidoff/rest"

type RouteHandler func(*rest.Request) *rest.Response

type Server interface {
	Start(address string, port int)
	Stop() error
	RegisterRoute(method string, url string, rh RouteHandler)
}
