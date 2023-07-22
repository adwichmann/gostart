package main

import (
	"github.com/adwichmann/gostart.git/config"
	"github.com/adwichmann/gostart.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//initialize configs
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
