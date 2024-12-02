package main

import (
	"basicAuth/controllers"
	"basicAuth/initializers"
	"basicAuth/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}
func main() {

	route := gin.Default()
	route.POST("/auth/signup", controllers.CreateUser)
	route.POST("/auth/login", controllers.Login)

	authRoute := route.Use(middlewares.CheckAuth)
	{
		authRoute.POST("/user/profile", controllers.GetUserProfile)
	}

	permissionRoute := route.Group("/permissions")
	permissionRoute.Use(middlewares.CheckAuth)
	{
		permissionRoute.POST("", middlewares.PermissionMiddleware("create_permission"), controllers.CreatePermission)
		permissionRoute.POST("/assign", middlewares.PermissionMiddleware("assign_permission"), controllers.AssignPermeissions)
	}

	route.Run()
}
