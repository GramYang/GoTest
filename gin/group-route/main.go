package main

import "github.com/gin-gonic/gin"

//以群组的方式设置路由
func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", handle)
		v1.POST("/submit", handle)
		v1.POST("/read", handle)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", handle)
		v2.POST("/submit", handle)
		v2.POST("/read", handle)
	}

	router.Run(":8080")
}

func handle(c *gin.Context) {
	c.JSON(200, "ok")
}
