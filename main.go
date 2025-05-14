package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "go_swagger/docs"
)

// @title نمونه API
// @version 1.0
// @description این یک نمونه API با Swagger است
// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()

	// Routes
	v1 := r.Group("/api/v1")
	{
		v1.GET("/hello", helloHandler)
	}

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// helloHandler نمونه هندلر
// @Summary سلام دنیا
// @Description این endpoint یک سلام برمی‌گرداند
// @Produce json
// @Success 200 {string} string "پیام سلام"
// @Router /hello [get]
func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "سلام دنیا!",
	})
}
