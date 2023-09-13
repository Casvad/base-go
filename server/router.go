package server

import (
	"baseGo/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getRouter() *gin.Engine {

	ginRouter := gin.Default()

	//ginRouter.Use() middlewares
	c1 := controllers.ProvideBaseGoExampleController()
	baseGo := ginRouter.Group("/base-go")
	{
		baseGo.GET("/health", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"status": "OK"})
		})

		baseGo.GET("/data", c1.GetData)
	}

	return ginRouter
}
