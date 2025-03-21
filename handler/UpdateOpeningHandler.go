package handler

import (
	"github.com/1guilherme1python1/go-api-vagas/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

//	@Summary Update opening
//
// @Description Update a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening Identification"
// @Param opening body UpdateOpeningRequest true "Opening data to Update"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	userEmail, exists := ctx.Get("email")
	if !exists {
		SendErrorResponse(ctx, http.StatusUnauthorized, "User email not found")
		return
	}

	err := ctx.BindJSON(&request)
	if err != nil {
		SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		logger.Errof("validation error: %v", err)
		return
	}

	id := ctx.Query("id")

	if id == "" {
		SendErrorResponse(ctx, http.StatusBadRequest, "id is required")
		return
	}

	opening := schemas.Opening{}

	result := db.Where("id = ? AND email = ?", id, userEmail).First(&opening)

	if result.RowsAffected == 0 {
		SendErrorResponse(ctx, http.StatusNotFound, "not found")
		return
	}

	if result.Error != nil {
		SendErrorResponse(ctx, http.StatusInternalServerError, "Server Error")
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if result := db.Save(&opening); result.Error != nil {
		logger.Errof("error updating opening: %v", result.Error)
		SendErrorResponse(ctx, http.StatusInternalServerError, "Server Error")
		return
	}

	SendSuccessResponse(ctx, http.StatusOK, opening)

}
