package handler

import (
	"fmt"
	"net/http"

	"github.com/PedroFranca404/TodoList-API-Go/schemas"
	"github.com/gin-gonic/gin"
)

func ShowTodoHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	todo := schemas.Todo{}

	// Find todo
	if err := db.First(&todo, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("todo with id: %s not found", id))
		return
	}

	sendSuccess(ctx, "show-todo", todo)
}
