package server

import "github.com/J-Obog/paidoff/rest"

type RouteHandler func(*rest.Request) *rest.Response

type Server interface {
	Start(port int) error
	Stop() error
}
