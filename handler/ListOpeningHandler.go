package handler

import (
	"github.com/1guilherme1python1/go-api-vagas/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary List openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/openings [get]
func ListOpeningHandler(ctx *gin.Context) {

	userEmail, exists := ctx.Get("email")
	if !exists {
		SendErrorResponse(ctx, http.StatusUnauthorized, "User email not found")
		return
	}

	var openings []schemas.Opening

	if err := db.Where("email = ?", userEmail.(string)).Find(&openings).Error; err != nil {
		SendErrorResponse(ctx, http.StatusInternalServerError, "Error listing openings")
		return
	}

	if len(openings) == 0 {
		SendErrorResponse(ctx, http.StatusNotFound, "No openings found for this user")
		return
	}

	var openingsResponse []schemas.OpeningResponse
	for _, opening := range openings {
		openingsResponse = append(openingsResponse, schemas.OpeningResponse{
			ID:        opening.ID,
			CreatedAt: opening.CreatedAt,
			UpdatedAt: opening.UpdatedAt,
			Role:      opening.Role,
			Company:   opening.Company,
			Location:  opening.Location,
			Remote:    opening.Remote,
			Link:      opening.Link,
			Salary:    opening.Salary,
		})
		logger.Infof("ListOpeningHandler: opening %+v", openingsResponse)
	}

	SendSuccessResponse(ctx, http.StatusOK, openingsResponse)
}
