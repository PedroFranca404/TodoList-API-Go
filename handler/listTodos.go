package handler

import (
	"net/http"

	"github.com/PedroFranca404/TodoList-API-Go/schemas"
	"github.com/gin-gonic/gin"
)

func ListTodosHandler(ctx *gin.Context) {
	todos := []schemas.Todo{}

	if err := db.Find(&todos).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing todos")
		return
	}

	sendSuccess(ctx, "list-todos", todos)
}
