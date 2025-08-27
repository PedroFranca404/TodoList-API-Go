package handler

import (
	"net/http"

	"github.com/PedroFranca404/TodoList-API-Go/schemas"
	"github.com/gin-gonic/gin"
)

func CreateTodoHandler(ctx *gin.Context) {
	request := CreateTodoRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	todo := schemas.Todo{
		Text: request.Text,
	}

	if err := db.Create(&todo).Error; err != nil {
		logger.Errorf("error creating todo: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating todo on database")
		return
	}

	sendSuccess(ctx, "create-todo", todo)
}
