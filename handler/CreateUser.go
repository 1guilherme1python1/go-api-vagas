package handler

import (
	"fmt"
	"github.com/1guilherme1python1/go-api-vagas/handler/requests"
	"github.com/1guilherme1python1/go-api-vagas/handler/responses"
	"github.com/1guilherme1python1/go-api-vagas/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUserHandler(ctx *gin.Context) {
	var request = requests.LoginRequest{}

	err := ctx.BindJSON(&request)
	if err != nil {
		fmt.Printf("Erro ao fazer bind: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = request.Validate()
	if err != nil {
		logger.Errof("erro ao validate: %v\n", err.Error())
		responses.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{
		Email:    request.Email,
		Password: request.Password,
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errof("error creating opening: %v\n", err.Error())
		responses.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	logger.Infof("request: %+v\n", request)
	responses.SendSuccessResponse(ctx, http.StatusOK, user)
}
