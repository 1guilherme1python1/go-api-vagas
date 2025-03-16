package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()
	initializeRoutes(r)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
