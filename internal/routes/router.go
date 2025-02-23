package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	return router

}
