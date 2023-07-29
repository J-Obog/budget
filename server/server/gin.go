package server

import (
	"fmt"
	"io"

	"github.com/J-Obog/paidoff/api"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/gin-gonic/gin"
)

type GinServer struct {
	router *gin.Engine
}

func NewGinServer(config api.APIConfig) *GinServer {
	router := gin.Default()

	router.GET("/account", ginHandler(config.AccountAPI.Get))
	router.PUT("/account", ginHandler(config.AccountAPI.Update))
	router.DELETE("/account", ginHandler(config.AccountAPI.Delete))

	router.GET("/category", ginHandler(config.CategoryAPI.GetAll))
	router.GET("/category/:id", ginHandler(config.CategoryAPI.Get))
	router.POST("/category", ginHandler(config.CategoryAPI.Create))
	router.PUT("/category/:id", ginHandler(config.CategoryAPI.Update))
	router.DELETE("/category/:id", ginHandler(config.CategoryAPI.Delete))

	router.GET("/transaction", ginHandler(config.TransactionAPI.Filter))
	router.GET("/transaction/:id", ginHandler(config.TransactionAPI.Get))
	router.POST("/transaction", ginHandler(config.TransactionAPI.Create))
	router.PUT("/transaction/:id", ginHandler(config.TransactionAPI.Update))
	router.DELETE("/transaction/:id", ginHandler(config.TransactionAPI.Delete))

	router.GET("/budget", ginHandler(config.BudgetAPI.Filter))
	router.GET("/budget/:id", ginHandler(config.BudgetAPI.Get))
	router.POST("/budget", ginHandler(config.BudgetAPI.Create))
	router.PUT("/budget/:id", ginHandler(config.BudgetAPI.Update))
	router.DELETE("/budget/:id", ginHandler(config.BudgetAPI.Delete))

	return &GinServer{
		router: router,
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
		Account: &data.Account{
			Id: "some-account-id",
		},

		Url:   c.Request.URL.String(),
		Query: c.Request.URL.Query(),
		Body:  b,
	}

	if id, inRequest := c.Params.Get("id"); inRequest {
		req.ResourceId = id
	}

	return req
}

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
