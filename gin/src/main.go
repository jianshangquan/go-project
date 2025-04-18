package main

import (
	"myapp/src/di"
	"myapp/src/features/chat"
	healthcheck "myapp/src/features/health-check"
	"myapp/src/features/users"
	"myapp/src/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)



func main() {
	var container = di.DependencyInjection()
	r := setupRouter(container)
	// Listen and Server in 0.0.0.0:8080 
	r.Run(":8080")
}









func setupRouter(container *dig.Container) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	
	r := gin.Default()
	println("Call container invoke")
	container.Invoke(func(
		loggermiddleware *middleware.LoggerMiddleware,
		errorHandlerMiddleware *middleware.ErrorHandler,
		chatController *chat.ChatController,
	) {
		println("invoke")
		loggermiddleware.RegisterRoute(r)
		errorHandlerMiddleware.RegisterRoute(r)
		chatController.RegisterRoute(r)
	})



	var v1 = r.Group("/v1")
	container.Invoke(func(
		userController *users.UserController,
		healthCheckController *healthcheck.HealthCheckController,
	) {
		userController.RegisterRoute(v1)
		healthCheckController.RegisterRoute(v1)
	})

	return r
}


