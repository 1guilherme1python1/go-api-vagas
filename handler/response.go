package handler

import "github.com/gin-gonic/gin"

func sendErrorResponse(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"errorCode": code,
		"msg":       msg,
	})
}

func sendSuccessResponse(ctx *gin.Context, code int, msg interface{}) {
	ctx.JSON(code, gin.H{
		"errorCode": code,
		"msg":       msg,
	})
}
