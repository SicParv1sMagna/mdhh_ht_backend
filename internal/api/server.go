package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *Application) StartServer() {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("/register", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{})
		})

		user.POST("/login", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{})
		})
	}
}
