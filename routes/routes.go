package routes

import (
	"golang-rest-api-authentication/controllers"
	"golang-rest-api-authentication/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func ConnectRoutes() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.POST("/auth/register", controllers.RegisterUser)
	v1.POST("/auth/login", controllers.LoginUser)

	v1Protected := v1.Use(middleware.JWTAuthMiddleware())

	v1Protected.GET("/", controllers.HelloRoutes)
	v1Protected.GET("/profile", controllers.GetMyProfile)
	v1Protected.PATCH("/profile", controllers.UpdateMyProfile)

	v1AdminRoutes := router.Group("/v1/admin")

	v1AdminAccess := v1AdminRoutes.Use(middleware.JWTAuthMiddleware(), middleware.IsAdminAccess())

	v1AdminAccess.GET("/users", controllers.GetAllUsers)
	v1AdminAccess.GET("/users/:username", controllers.GetUser)
	v1AdminAccess.POST("/users", controllers.CreateUser)
	v1AdminAccess.DELETE("/users/:username", controllers.DeleteUser)

	router.Run(os.Getenv("APP_PORT"))

}
