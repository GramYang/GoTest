package main

import (
	"github.com/gin-gonic/gin"
)

//运行后再浏览器输入http://localhost:8080/ping就可以看到反馈：{message: "pong"}
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	_ = r.Run()
}
