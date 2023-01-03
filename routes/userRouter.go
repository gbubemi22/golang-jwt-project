package routes

import (
	controller "github.com/gbubemi22/golang-jwt-project/controllers"
	"github.com/gin-gonic/gin"
	"github.com/gbubemi22/golang-jwt-project/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/user_id", controller.GetUser())
}
