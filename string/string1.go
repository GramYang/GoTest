package main

import (
	"fmt"
	"strings"
)

func main() {
	st1()
}

func st1() {
	//string引用规则测试
	s1 := "nmsl"
	s2 := s1
	fmt.Printf("%p %p\n", &s1, &s2) //指针不同
	s1 += "孙笑川"
	fmt.Println(s1, s2) //nmsl孙笑川 nmsl
	//string拼接
	s3 := "你妈"
	s4 := "死了"
	var buf strings.Builder
	_, _ = buf.WriteString(s3)
	_, _ = buf.WriteString(s4)
	_, _ = buf.WriteString("效率最高")
	fmt.Println(s3+s4+"效率低", fmt.Sprintf("%s+%s效率低", s3, s4),
		strings.Join([]string{s3, s4}, "数组专用"), buf.String())
}
