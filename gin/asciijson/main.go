package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//使用AsciiJSON生成只有ascii格式的JSON，并使用转义的非ascii字符。
//输出{"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
		c.ShouldBind()
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
