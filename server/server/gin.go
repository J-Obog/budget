package server

import (
	"fmt"
	"io"

	"github.com/J-Obog/paidoff/api"
	"github.com/J-Obog/paidoff/rest"
	"github.com/gin-gonic/gin"
)

type GinServer struct {
	router *gin.Engine
}

func NewGinServer(config api.APIConfig) *GinServer {
	router := gin.Default()

	accountRouter := router.Group("/account")
	{
		accountRouter.GET("/", ginHandler(config.AccountAPI.Get))
		accountRouter.PUT("/", ginHandler(config.AccountAPI.Update))
		accountRouter.DELETE("/", ginHandler(config.AccountAPI.Delete))
	}

	categoryRouter := router.Group("/category")
	{
		categoryRouter.GET("/", ginHandler(config.CategoryAPI.GetAll))
		categoryRouter.GET("/:id", ginHandler(config.CategoryAPI.Get))
		categoryRouter.POST("/", ginHandler(config.CategoryAPI.Create))
		categoryRouter.PUT("/:id", ginHandler(config.CategoryAPI.Update))
		categoryRouter.DELETE("/:id", ginHandler(config.CategoryAPI.Delete))
	}

	transactionRouter := router.Group("/category")
	{
		transactionRouter.GET("/", ginHandler(config.TransactionAPI.Filter))
		transactionRouter.GET("/:id", ginHandler(config.TransactionAPI.Get))
		transactionRouter.POST("/", ginHandler(config.TransactionAPI.Create))
		transactionRouter.PUT("/:id", ginHandler(config.TransactionAPI.Update))
		transactionRouter.DELETE("/:id", ginHandler(config.TransactionAPI.Delete))
	}

	budgetRouter := router.Group("/budget")
	{
		budgetRouter.GET("/", ginHandler(config.BudgetAPI.Filter))
		budgetRouter.GET("/:id", ginHandler(config.BudgetAPI.Get))
		budgetRouter.POST("/", ginHandler(config.BudgetAPI.Create))
		budgetRouter.PUT("/:id", ginHandler(config.BudgetAPI.Update))
		budgetRouter.DELETE("/:id", ginHandler(config.BudgetAPI.Delete))
	}

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
		Url:   c.Request.URL.String(),
		Query: c.Request.URL.Query(),
		Body:  b,
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
