package router

import "github.com/gin-gonic/gin"

func Init() {
	router := gin.Default()
	router.Run(":8080")
}
