package routes

import (
	"auth-microservice/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
}
