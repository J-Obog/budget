package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/J-Obog/paidoff/rest"
	"github.com/gin-gonic/gin"
)

type GinServer struct {
	router  *gin.Engine
	srv     *http.Server
	address string
	port    int
}

func NewGinServer(address string, port int) *GinServer {
	return &GinServer{
		router:  gin.Default(),
		address: address,
		port:    port,
	}
}

func (g *GinServer) RegisterRoute(method string, url string, rh RouteHandler) {
	ginFn := func(c *gin.Context) {
		//TODO: handle error when reading body

		b, _ := io.ReadAll(c.Request.Body)

		// TODO: remove dummy account for auth
		req := &rest.Request{
			Url:   c.Request.URL.String(),
			Query: c.Request.URL.Query(),
			Body:  b,
		}

		res := rh(req)
		respb, code := res.ToJSON()

		if code == 500 {
			//log error
		}

		c.JSON(code, respb)
	}

	switch method {
	case http.MethodGet:
		g.router.GET(url, ginFn)

	case http.MethodPost:
		g.router.POST(url, ginFn)

	case http.MethodPut:
		g.router.PUT(url, ginFn)

	case http.MethodDelete:
		g.router.DELETE(url, ginFn)

	default:
		log.Fatalf("cannot handle http method %s", method)
	}
}

func (g *GinServer) Start() {
	g.srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", g.address, g.port),
		Handler: g.router,
	}

	go func() {
		if err := g.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error %s", err)
		}
	}()
}

func (g *GinServer) Stop() error {
	log.Println("shutting down server")

	if err := g.srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error shutting down server %s", err)
	}

	return nil
}
