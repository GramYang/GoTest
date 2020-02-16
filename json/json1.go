package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
)

type personInfo struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email" xml:"email"`
}

type personInfo1 struct {
	Name  string `json:"name"`
	Email string `json:"email" xml:"email"`
	C     string
}

func main() {
	// 创建数据
	p := personInfo{Name: "Piao", Age: 10, Email: "piaoyunsoft@163.com"}

	// 序列化
	data, _ := json.Marshal(&p) //这里其实传入值类型和指针类型都可以
	fmt.Println(string(data))   //{"name":"Piao","age":10,"email":"piaoyunsoft@163.com"}

	// 反序列化
	var p1 personInfo1
	err := json.Unmarshal(data, &p1)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Printf("name=%s, c=%s, email=%s\n", p1.Name, p1.C, p1.Email) //name=Piao, c=, email=piaoyunsoft@163.com
	}
	fmt.Printf("%+v\n", p1) //{Name:Piao Email:piaoyunsoft@163.com C:}

	// 反序列化
	res, err := simplejson.NewJson([]byte(data))
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Printf("%+v\n", res) //&{data:map[age:10 email:piaoyunsoft@163.com name:Piao]}
	}
}
