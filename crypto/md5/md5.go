package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	mt1()
}

func mt1() {
	ctx := md5.New()
	ctx.Write([]byte("我是你哥哥，我们两个都是你妈的儿子"))
	s := ctx.Sum([]byte("我是令牌"))
	fmt.Println(hex.EncodeToString(s))
}
