package handler

import (
	"github.com/1guilherme1python1/go-api-vagas/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQlite()
}
