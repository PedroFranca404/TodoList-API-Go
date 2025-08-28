package handler

import (
	"net/http"

	"github.com/PedroFranca404/TodoList-API-Go/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateTodoHandler(ctx *gin.Context) {
	request := UpdateTodoRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadGateway, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
		return
	}

	todo := schemas.Todo{}

	if err := db.First(&todo, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "todo not found")
	}

	// Update todo
	todo.Text = request.Text

	// Save todo
	if err := db.Save(&todo).Error; err != nil {
		logger.Errorf("error updating todo: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating todo")
		return
	}
	sendSuccess(ctx, "update-todo", todo)
}
