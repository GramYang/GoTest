package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//打印ids: map[a:1234 b:hello]; names: map[first:thinkerou second:tianou]
func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})
	router.Run(":8080")
}
