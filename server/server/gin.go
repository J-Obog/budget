package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	BaseServer
	eng *gin.Engine
}

func (g *GinServer) Start(port int) error {
	g.eng.GET("/account", ginHandler(g.accountManager.GetByRequest))
	g.eng.PUT("/account", ginHandler(g.accountManager.UpdateByRequest))
	g.eng.DELETE("/account", ginHandler(g.accountManager.DeleteByRequest))

	g.eng.GET("/category", ginHandler(g.categoryManager.GetAllByRequest))
	g.eng.GET("/category/:id", ginHandler(g.categoryManager.GetByRequest))
	g.eng.POST("/category", ginHandler(g.categoryManager.CreateByRequest))
	g.eng.PUT("/category/:id", ginHandler(g.categoryManager.UpdateByRequest))
	g.eng.DELETE("/category/:id", ginHandler(g.accountManager.DeleteByRequest))

	g.eng.GET("/transaction", ginHandler(g.transactionManager.GetAllByRequest))
	g.eng.GET("/transaction/:id", ginHandler(g.transactionManager.GetByRequest))
	g.eng.POST("/transaction", ginHandler(g.transactionManager.CreateByRequest))
	g.eng.PUT("/transaction/:id", ginHandler(g.transactionManager.UpdateByRequest))
	g.eng.DELETE("/transaction/:id", ginHandler(g.transactionManager.DeleteByRequest))

	g.eng.GET("/budget", ginHandler(g.budgetManager.GetAllByRequest))
	g.eng.GET("/budget/:id", ginHandler(g.budgetManager.GetByRequest))
	g.eng.POST("/budget", ginHandler(g.budgetManager.CreateByRequest))
	g.eng.PUT("/budget/:id", ginHandler(g.budgetManager.UpdateByRequest))
	g.eng.DELETE("/budget/:id", ginHandler(g.budgetManager.DeleteByRequest))

	return g.eng.Run(fmt.Sprintf(":%d", port))
}

func (g *GinServer) Stop() error {
	return nil
}

func ginHandler(rh routeHandler) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
