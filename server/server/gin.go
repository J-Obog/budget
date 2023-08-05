package server

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/J-Obog/paidoff/rest"
	"github.com/gin-gonic/gin"
)

type GinServer struct {
	router *gin.Engine
}

func NewGinServer() *GinServer {
	return &GinServer{
		router: gin.Default(),
	}
}

func (g *GinServer) RegisterRoute(method string, url string, rh RouteHandler) {
	ginFn := func(c *gin.Context) {
		req := ginCtxToRequest(c)
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

func (g *GinServer) Start(port int) error {
	return g.router.Run(fmt.Sprintf(":%d", port))
}

func (g *GinServer) Stop() error {
	return nil
}

func ginCtxToRequest(c *gin.Context) *rest.Request {
	//TODO: handle error when reading body

	b, _ := io.ReadAll(c.Request.Body)

	// TODO: remove dummy account for auth
	req := &rest.Request{
		Url:   c.Request.URL.String(),
		Query: c.Request.URL.Query(),
		Body:  b,
	}

	return req
}

/*
func ginHandler(rh routeHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := ginCtxToRequest(c)

		res := rh(req)

		respb, code := res.ToJSON()

		if code == 500 {
			//log error
		}

		c.JSON(code, respb)
	}
}
*/
