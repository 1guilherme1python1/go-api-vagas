package handler

import (
	"github.com/1guilherme1python1/go-api-vagas/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

//	@Summary Show opening
//
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendErrorResponse(ctx, http.StatusBadRequest, errParamIdRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	result := db.First(&opening, id)

	if result.RowsAffected == 0 {
		sendErrorResponse(ctx, http.StatusNotFound, "not found")
		return
	}

	if result.Error != nil {
		sendErrorResponse(ctx, http.StatusInternalServerError, "database error")
		return
	}

	SendSuccessResponse(ctx, http.StatusOK, opening)
}
