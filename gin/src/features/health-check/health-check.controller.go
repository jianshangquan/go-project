package healthcheck

import "github.com/gin-gonic/gin"

type HealthCheckController struct{}

func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{};
}


func (controller *HealthCheckController) GetHealth(c *gin.Context){
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func (controller *HealthCheckController) RegisterRoute(r *gin.RouterGroup){
	r.GET("/health", controller.GetHealth)
}