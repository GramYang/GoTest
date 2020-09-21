package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/secure"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()

	r.Use(secure.Secure(secure.Options{
		//唯一允许的域名
		AllowedHosts: []string{"example.com", "ssl.example.com"},
		//只允许https的请求
		SSLRedirect: true,
		//将http重定向到https的域名
		SSLHost: "ssl.example.com",
		//搭配nginx的
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		//STS头的存活时间
		STSSeconds: 315360000,
		//STS是否拼接includeSubdomains
		STSIncludeSubdomains: true,
		//添加头X-Frame-Options=DENY
		FrameDeny: true,
		//添加头X-Content-Type-Options=nosniff
		ContentTypeNosniff: true,
		//添加头X-XSS-Protection=1; mode=block
		BrowserXssFilter: true,
		//设置头Content-Security-Policy的值
		ContentSecurityPolicy: "default-src 'self'",
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
