package routes

import (
	_ "save-the-queen-backend/docs"
	"save-the-queen-backend/src/api"
	"save-the-queen-backend/src/auth"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Routes function is used for the routing for our apis
func Routes(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/signup", api.SignUp)
	router.POST("/login", api.Login)
	auth := router.Group("/secured").Use(auth.Authorisaton)
	{
		auth.POST("/game", api.GameCreate)
		auth.GET("/game", api.ListGame)
	}
	router.Run()
}
