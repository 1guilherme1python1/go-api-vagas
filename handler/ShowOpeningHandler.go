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
// @Router /api/v1/opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	userEmail, exists := ctx.Get("email")
	if !exists {
		SendErrorResponse(ctx, http.StatusUnauthorized, "User email not found")
		return
	}

	if id == "" {
		SendErrorResponse(ctx, http.StatusBadRequest, ErrParamIdRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	result := db.Where("id = ? AND email = ?", id, userEmail).First(&opening)

	if result.RowsAffected == 0 {
		SendErrorResponse(ctx, http.StatusNotFound, "not found")
		return
	}

	if result.Error != nil {
		SendErrorResponse(ctx, http.StatusInternalServerError, "database error")
		return
	}

	response := OpeningResponse{
		ID:        opening.ID,
		CreatedAt: opening.CreatedAt,
		UpdatedAt: opening.UpdatedAt,
		Role:      opening.Role,
		Company:   opening.Company,
		Location:  opening.Location,
		Remote:    opening.Remote,
		Link:      opening.Link,
		Salary:    opening.Salary,
	}

	SendSuccessResponse(ctx, http.StatusOK, response)
}
