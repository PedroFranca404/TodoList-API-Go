package main

import (
	"github.com/PedroFranca404/TodoList-API-Go/config"
	"github.com/PedroFranca404/TodoList-API-Go/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize router
	router.Init()
}
