package middleware

import (
	"myapp/src/logger"

	"github.com/gin-gonic/gin"
)

type LoggerMiddleware struct {
	Logger *logger.Logger;
}

func NewLoggerMiddleware(logger *logger.Logger) *LoggerMiddleware{
	println("New Logger middleware created")
	return &LoggerMiddleware{
		Logger: logger,
	}
}

func (l *LoggerMiddleware) Log(c *gin.Context){
	l.Logger.Log(c.Request.URL.String())
	c.Next()
}

func (l *LoggerMiddleware) RegisterRoute(routeGroup *gin.Engine){
	println("logger middle route registered")
	routeGroup.Use(l.Log)
}