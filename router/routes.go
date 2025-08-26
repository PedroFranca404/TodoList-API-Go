package router

import (
	"github.com/PedroFranca404/TodoList-API-Go/handler"
	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {
	handler.Init()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/todo")
		v1.GET("/todo")
		v1.PUT("/todo")
		v1.DELETE("/todo")
		v1.GET("/todos")
	}
}
