package main

import (
	"github.com/J-Obog/paidoff/api"
	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/server"
)

func runServer() {
	cfg := config.Get()

	apiSvc := api.NewApiService(cfg)

	svr := server.NewServer(cfg)

	server.RegisterRoutes(svr, apiSvc)

	svr.Start(
		cfg.ServerAddress,
		cfg.ServerPort,
	)
}
