package handler

import (
	"github.com/1guilherme1python1/go-api-vagas/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

//	@Summary Delete opening
//
// @Description Delete a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendErrorResponse(
			ctx,
			http.StatusBadRequest,
			errParamIdRequired("id", "queryParameter").Error(),
		)
		return
	}

	opening := schemas.Opening{}

	//find opening
	if err := db.First(&opening, id).Error; err != nil {
		sendErrorResponse(ctx, http.StatusBadRequest, "opening not found")
		return
	}

	//delete opening
	if err := db.Delete(&opening).Error; err != nil {
		sendErrorResponse(ctx, http.StatusInternalServerError, "Error deleting opening")
		return
	}

	sendSuccessResponse(ctx, http.StatusOK, "success")
}
