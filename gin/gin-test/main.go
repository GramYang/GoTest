package main

import (
	gc "GoTest/crypto"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strings"
	"time"
)

const md5Key = "cqdq_iot"

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			//允许访问origin的域
			c.Header("Access-Control-Allow-Origin", origin)
			//指定允许请求方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//指定允许request中的头
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,signature")
			//指定允许response中的头
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Content-Disposition,Expires,Last-Modified,Pragma,FooBar")
			//response可以被缓存多久
			c.Header("Access-Control-Max-Age", "172800")
			//指示当请求的凭证标记为 true 时，是否响应该请求
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}

type j struct {
	A int `json:"a"`
	B int `json:"b"`
	C int `json:"c"`
}

type j2 struct {
	Data []byte `json:"data"`
}

func md5encode() string {
	ctx := md5.New()
	ctx.Write([]byte(md5Key))
	return hex.EncodeToString(ctx.Sum(nil))
}

func main() {
	r := gin.Default()
	//r.Use(cors())
	r.POST("/greet/post", func(c *gin.Context) {
		fmt.Println(c.Request)
		var j j
		err := c.BindJSON(&j)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(j)
		if j.A == 114514 {
			c.JSON(200, gin.H{
				"message": "1919810",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "success",
		})
	})
	r.GET("/greet/get", func(c *gin.Context) {
		fmt.Println(c.Request)
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.POST("/rsa", func(c *gin.Context) {
		var j j2
		err := c.BindJSON(&j)
		if err != nil {
			fmt.Println(err)
		}
		data, err := gc.RsaDecrypt(j.Data, "crypto/key/priv.pem")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("rsa: " + string(data))
	})
	r.POST("/rsa1", func(c *gin.Context) {
		var j j2
		err := c.BindJSON(&j)
		if err != nil {
			fmt.Println(err)
		}
		data, err := gc.RsaDecrypt1(j.Data, "crypto/key/priv.pem")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("rsa: " + string(data))
	})
	r.GET("/paramencode", func(c *gin.Context) {
		sig := c.GetHeader("signature")
		if sig != md5encode() {
			c.JSON(400, gin.H{
				"message": "wrong signature",
			})
			return
		}
		p := c.Query("param")
		p1, err := base64.StdEncoding.DecodeString(p)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(p1))
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.GET("/download_test", func(c *gin.Context) {
		fileName := "ctwing.xlsx"
		c.Header("Content-Disposition", "attachment;filename="+fileName)
		c.Header("Content-Type", "application/.nmsl")
		c.File("gin/" + fileName)
	})
	r.POST("/upload_test", func(c *gin.Context) {
		file, _ := c.FormFile("114514")
		if file != nil {
			dst := path.Join("gin/", file.Filename+time.Now().Format(".20060102.150405"))
			fmt.Println(dst)
			err := c.SaveUploadedFile(file, dst)
			if err != nil {
				fmt.Println(err)
			}
		}
	})
	r.Run(":8090")
}
