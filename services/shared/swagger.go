package shared

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupSwagger(router *gin.Engine, serviceName string) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"service":   serviceName,
			"status":    "healthy",
			"timestamp": gin.Now(),
		})
	})
}
