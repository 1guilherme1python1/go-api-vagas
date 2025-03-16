package main

import (
	"fmt"
	"github.com/1guilherme1python1/go-api-vagas/config"
	"github.com/1guilherme1python1/go-api-vagas/router"
)

func main() {
	//Initialize config
	err := config.Init()
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	router.Initialize()
}
