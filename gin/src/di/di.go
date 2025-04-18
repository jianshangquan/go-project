package di

import (
	"myapp/src/database"
	"myapp/src/features/chat"
	healthcheck "myapp/src/features/health-check"
	"myapp/src/features/users"
	"myapp/src/logger"
	"myapp/src/middleware"

	"go.uber.org/dig"
)






func DependencyInjection() *dig.Container {
	// This is the place to put all the dependency injection
	// for the application.
	// This is a good place to put the database connection
	// and other dependencies that are needed for the application.
	// This is also a good place to put the configuration
	// for the application.
	var container *dig.Container = dig.New()

	// database
	database.SetupDatabase(container)


	// middlewares
	container.Provide(logger.NewLogger)
	container.Provide(middleware.NewErrorHandler)
	container.Provide(middleware.NewLoggerMiddleware)


	// /health
	container.Provide(healthcheck.NewHealthCheckController)

	// /user
	container.Provide(users.NewUserRepository)
	container.Provide(users.NewUserService)
	container.Provide(users.NewUserController)

	// chat
	container.Provide(chat.NewChatController)

	return container;
}