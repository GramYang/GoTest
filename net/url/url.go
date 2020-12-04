package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"time"
)

func main() {
	//测试一下url.QueryEscape的适用范围
	t1()
}

func t1() {
	c := CookieGenerator("1234@蔡徐坤")
	fmt.Println(c == url.QueryEscape(c)) //true，用md5签名的cookie可以通过验证
}

func CookieGenerator(name string) string {
	s := name + "cqdq_iot" + time.Now().String()
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}
