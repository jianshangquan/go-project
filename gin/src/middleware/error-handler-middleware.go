package middleware

import (
	"myapp/src/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorHandler struct{
	Logger *logger.Logger
}

func NewErrorHandler(logger *logger.Logger) *ErrorHandler{
	return &ErrorHandler{
		Logger: logger,
	}
}

func (err *ErrorHandler) handle(c *gin.Context){
	c.Next();

	// If errors occurred during handlers
	if len(c.Errors) > 0 {
		for _, e := range c.Errors {
			err.Logger.Log("Error: " + e.Err.Error())  // Log all errors
		}

		// Uniform error response
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": c.Errors[0].Error(),  // Return the first error
		})
	}
}

func (err *ErrorHandler) RegisterRoute(r *gin.Engine){
	println("Error registered")
	r.Use(err.handle)
}