package handler

import (
	"fmt"
	"github.com/1guilherme1python1/go-api-vagas/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	var request = CreateOpeningRequest{}

	err := ctx.BindJSON(&request)
	if err != nil {
		fmt.Printf("Erro ao fazer bind: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = request.Validate()
	if err != nil {
		logger.Errof("erro ao validate: %v\n", err.Error())
		sendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Salary:   request.Salary,
		Location: request.Location,
		Company:  request.Company,
		Role:     request.Role,
		Remote:   *request.Remote,
		Link:     request.Link,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errof("error creating opening: %v\n", err.Error())
		sendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	logger.Infof("request: %+v\n", request)
	sendSuccessResponse(ctx, http.StatusOK, "success")
}
