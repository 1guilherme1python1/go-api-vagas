package main

import (
	"github.com/1guilherme1python1/go-api-vagas/config"
	"github.com/1guilherme1python1/go-api-vagas/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	//Initialize config
	err := config.Init()
	if err != nil {
		logger.Error("config init fail")
		return
	}

	router.Initialize()
}
