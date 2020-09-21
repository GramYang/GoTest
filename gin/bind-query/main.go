package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

//访问localhost:8085/testing?name=appleboy&address=xyz&birthday=1992-03-15返回success。
//这说明字段是否正确需要你自己检查，且字段有多有少错误甚至没有都是返回success
func main() {
	route := gin.Default()
	route.GET("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person

	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}

	c.String(200, "Success")
}
