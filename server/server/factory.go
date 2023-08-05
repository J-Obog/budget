package server

import (
	"github.com/J-Obog/paidoff/config"
)

const (
	serverImpl = "gin"
)

func NewServer(cfg *config.AppConfig) Server {
	switch serverImpl {
	case "gin":
		return NewGinServer()

	default:
		panic("no implementation for server")
	}
}
