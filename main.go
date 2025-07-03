package main

import (
	"github.com/arturrpereira/go-api/config"
	"github.com/arturrpereira/go-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initilization error:  %v", err)
		return
	}

	// Initialize router
	router.Initialize()
}
