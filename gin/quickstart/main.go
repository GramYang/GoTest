package main

import "github.com/gin-gonic/gin"

//访问0.0.0.0:8080/ping网页显示{"message":"pong"}
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
