package server

import "github.com/J-Obog/paidoff/rest"

type routeHandler func(*rest.Request) *rest.Response

type Server interface {
	Start(port int) error
	Stop() error
}
