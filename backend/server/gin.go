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
	router *gin.Engine
	srv    *http.Server
}

func NewGinServer() *GinServer {
	return &GinServer{
		router: gin.New(),
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

func (g *GinServer) Start(address string, port int) {
	g.srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", address, port),
		Handler: g.router,
	}

	log.Printf("starting server on address %s, port %d\n", address, port)

	go func() {
		if err := g.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error %s", err)
		}
	}()
}

func (g *GinServer) Stop() error {
	if err := g.srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error shutting down server %s", err)
	}

	log.Println("server has been shut down")

	return nil
}
