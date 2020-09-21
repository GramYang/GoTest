package main

import "github.com/gin-gonic/gin"

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

///thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3
//{"name":"thinkerou","uuid":"987fbc97-4bed-5078-9f07-9141ba07c9f3"}
///thinkerou/not-uuid
//{"msg":[{}]}
//可以看出，结构体的uri注解中可以要求required和格式为uuid。如果丢失required项返回404 page not found
//如果格式错误就会返回错误
func main() {
	route := gin.Default()
	route.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})
	route.Run(":8088")
}
