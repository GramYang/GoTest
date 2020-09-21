package main

import "github.com/gin-gonic/gin"

//通常JSON会替换特殊的HTML字符为unicode，如果你想输出纯粹的文字你要使用PureJSON。
func main() {
	r := gin.Default()

	// Serves unicode entities
	//{"html":"\u003cb\u003eHello, world!\u003c/b\u003e"}
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// Serves literal characters
	//{"html":"<b>Hello, world!</b>"}
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
