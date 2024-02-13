package main

import (
	"github.com/Nivaldeir/golang-opportunities/config"
	"github.com/Nivaldeir/golang-opportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initilization error: %v", err)
		return
	}
	router.Initialize()
}
