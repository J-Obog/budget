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
	g.eng.GET("/account")
	g.eng.PUT("/account")
	g.eng.DELETE("/account")

	g.eng.GET("/category")
	g.eng.GET("/category/:id")
	g.eng.POST("/category")
	g.eng.PUT("/category/:id")
	g.eng.DELETE("/category/:id")

	g.eng.GET("/transaction")
	g.eng.GET("/transaction/:id")
	g.eng.POST("/transaction")
	g.eng.PUT("/transaction/:id")
	g.eng.DELETE("/transaction/:id")

	g.eng.GET("/budget")
	g.eng.GET("/budget/:id")
	//g.eng.GET("/budget/types")
	g.eng.POST("/budget")
	g.eng.PUT("/budget/:id")
	g.eng.DELETE("/budget/:id")

	return g.eng.Run(fmt.Sprintf(":%d", port))
}

func (g *GinServer) Stop() error {
	return nil
}

func (g *GinServer) handle() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
