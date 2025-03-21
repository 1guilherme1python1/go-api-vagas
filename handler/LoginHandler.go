package handler

import (
	"github.com/1guilherme1python1/go-api-vagas/schemas"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

// LoginHandler godoc
// @BasePath /api/v1
// @Summary User login
// @Description Authenticate user and return JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body LoginRequest true "User credentials"
// @Success 200 {object} TokenResponse
// @Failure 400 {object} ErrorResponse "Invalid request"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 401 {object} ErrorResponse "Invalid email or password"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /api/v1/login [post]
func LoginHandler(ctx *gin.Context) {
	var creds LoginRequest
	if err := ctx.ShouldBindJSON(&creds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user := schemas.User{}

	result := db.Where("email = ?", creds.Email).First(&user)

	if result.RowsAffected == 0 {
		SendErrorResponse(ctx, http.StatusNotFound, "user not found")
		return
	}

	if result.Error != nil {
		SendErrorResponse(ctx, http.StatusInternalServerError, "Server Error")
		return
	}

	if creds.Password != user.Password {
		SendErrorResponse(ctx, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": creds.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("volmatavs"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	SendSuccessResponse(ctx, http.StatusOK, tokenString)
}
