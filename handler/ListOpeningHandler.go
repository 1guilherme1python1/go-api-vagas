package handler

import (
	"github.com/1guilherme1python1/go-api-vagas/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOpeningHandler(ctx *gin.Context) {

	var openings []schemas.Opening

	if err := db.Find(&openings).Error; err != nil {
		sendErrorResponse(ctx, http.StatusInternalServerError, "Error listing openigs")
		return
	}

	SendSuccessResponse(ctx, http.StatusOK, openings)
}
