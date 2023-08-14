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

func (g *GinServer) RegisterRoute(method string, url string, rh rest.RouteHandler) {
	ginFn := func(c *gin.Context) {
		//TODO: handle error when reading body?
		p := rest.PathParams{}
		b, _ := io.ReadAll(c.Request.Body)

		for _, param := range c.Params {
			p[param.Key] = param.Value
		}

		q := rest.Query{}
		for k, v := range c.Request.URL.Query() {
			q[k] = v
		}

		req := &rest.Request{
			Url:    c.Request.URL.String(),
			Params: p,
			Query:  q,
			Body:   rest.NewJSONBody(b),
		}

		res := rh(req)

		if res.Status == 500 {
			log.Printf("ERROR[5XX]: %s\n", res.InternalErrMsg)
		}

		c.JSON(res.Status, res.Data)
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

	if err := g.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error %s", err)
	}
}

func (g *GinServer) Stop() error {
	if err := g.srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error shutting down server %s", err)
	}

	log.Println("server has been shut down")

	return nil
}
