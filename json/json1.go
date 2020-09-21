package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
)

type personInfo struct {
	Name  string `json:"name"`
	Age   int8   `json:"age"`
	Email string `json:"email" xml:"email"`
}

type personInfo1 struct {
	Name  string `json:"name"`
	Email string `json:"email" xml:"email"`
	C     string
}

func main() {
	//正常的序列化和反序列化
	//t1()
	//json反序列化测试
	t2()
}

func t1() {
	// 创建数据
	p := &personInfo{Name: "Piao", Age: 10, Email: "piaoyunsoft@163.com"}

	// 序列化
	data, _ := json.Marshal(&p) //必须加上&，不管p是值类型还是指针类型
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

type bag struct {
	A string           `json:"a"`
	B int32            `json:"b"`
	C float32          `json:"c"`
	D bool             `json:"d"`
	E []int            `json:"e"`
	F map[int32]string `json:"f"`
}

//json的结构体的域必须是public的，当反序列化时json中的值的类型错误或者超越边界值时，会报错但是不会panic，然后给结构体域赋类型默认值
func t2() {
	b := &bag{A: "nmsl", B: 24, C: 24.4, D: true, E: []int{1, 2, 3, 4}, F: map[int32]string{1: "a", 2: "b", 3: "c"}}
	data, _ := json.Marshal(b)
	fmt.Println(string(data)) //{"a":"nmsl","b":24,"c":24.4,"d":true,"e":[1,2,3,4],"f":{"1":"a","2":"b","3":"c"}}
	//json字符串的域顺序可以无视，多一个域也没关系!
	//如果被反序列化的容器里面已经有域被赋值了怎么办？会被覆盖。如果刚好这个域的值出错了需要赋默认值的呢？？优先使用该域的当前值！
	s := "{\"g\":114514,\"b\":24,\"c\":24.4,\"d\":\"true\",\"e\":[1,2,3,4],\"f\":{\"1\":\"a\",\"2\":\"b\",\"3\":\"c\"},\"a\":\"nmsl\"}"
	var b1 bag
	b1.B = 114514
	b1.D = true
	err := json.Unmarshal([]byte(s), &b1)
	if err != nil {
		fmt.Println(err) //cannot unmarshal string into Go struct field bag.d of type bool
	}
	fmt.Println(b1) //{nmsl 24 24.4 false [1 2 3 4] map[1:a 2:b 3:c]}
}
