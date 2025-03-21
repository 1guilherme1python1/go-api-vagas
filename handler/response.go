package handler

import (
	"github.com/1guilherme1python1/go-api-vagas/schemas"
	"github.com/gin-gonic/gin"
	"time"
)

func SendErrorResponse(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"statusCode": code,
		"message":    msg,
	})
}

func SendSuccessResponse(ctx *gin.Context, code int, msg interface{}) {
	ctx.JSON(code, gin.H{
		"statusCode": code,
		"message":    msg,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"statusCode"`
}

type TokenResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"statusCode"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
type ShowOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
type ListOpeningsResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.OpeningResponse `json:"data"`
}
type UpdateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type OpeningResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
}
