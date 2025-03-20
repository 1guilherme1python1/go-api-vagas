package handler

import (
	"github.com/1guilherme1python1/go-api-vagas/handler/requests"
	"github.com/1guilherme1python1/go-api-vagas/handler/responses"
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
		responses.SendErrorResponse(
			ctx,
			http.StatusBadRequest,
			requests.ErrParamIdRequired("id", "queryParameter").Error(),
		)
		return
	}

	userEmail, exists := ctx.Get("email")
	if !exists {
		responses.SendErrorResponse(ctx, http.StatusUnauthorized, "User email not found")
		return
	}

	opening := schemas.Opening{}

	//find opening
	if err := db.Where("id = ? AND email = ?", id, userEmail).First(&opening).Error; err != nil {
		// Se não encontrar, ou a vaga não é dele, ou o ID não existe
		responses.SendErrorResponse(ctx, http.StatusNotFound, "Opening not found or you don't have permission to delete it")
		return
	}

	//delete opening
	if err := db.Delete(&opening).Error; err != nil {
		responses.SendErrorResponse(ctx, http.StatusInternalServerError, "Error deleting opening")
		return
	}

	responses.SendSuccessResponse(ctx, http.StatusOK, "success")
}
