package main

import "github.com/gin-gonic/gin"

type myForm struct {
	Colors []string `form:"colors[]"`
}

//显示颜色选择checkbox，提交后显示你选择的颜色
//自定义结构体中字段的`form:"colors[]"`和html中name字段匹配
func main() {
	r := gin.Default()

	r.LoadHTMLGlob("gin/bind-html/*")
	r.GET("/", indexHandler)
	r.POST("/", formHandler)

	r.Run(":8080")
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "form.html", nil)
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.Bind(&fakeForm)
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}
