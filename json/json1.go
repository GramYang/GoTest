package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
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
	//正常的json与结构体序列化和反序列化
	//t1()
	//json反序列化测试，用map
	//t2()
	//json数组与map的转化
	//t3()
	//json数组与map的转化2
	//t4()
	//测试json反序列化成结构体，字段是否必须导出，必须导出！
	//t5()
	//json反序列化，内嵌结构体能否是指针，不能！
	t6()
}

func t1() {
	// 创建数据
	p := &personInfo{Name: "Piao", Age: 10, Email: "piaoyunsoft@163.com"}

	// 序列化
	data, _ := json.Marshal(p)
	fmt.Println(string(data)) //{"name":"Piao","age":10,"email":"piaoyunsoft@163.com"}

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
	//res, err := simplejson.NewJson(data)
	//if err != nil {
	//	fmt.Println("err: ", err)
	//} else {
	//	fmt.Printf("%+v\n", res) //&{data:map[age:10 email:piaoyunsoft@163.com name:Piao]}
	//}
	//json的结构体的域必须是public的，当反序列化时json中的值的类型错误或者超越边界值时，会报错但是不会panic，然后给结构体域赋类型默认值
	b := &bag{A: "nmsl", B: 24, C: 24.4, D: true, E: []int{1, 2, 3, 4}, F: map[int32]string{1: "a", 2: "b", 3: "c"}}
	data, _ = json.Marshal(b)
	fmt.Println(string(data)) //{"a":"nmsl","b":24,"c":24.4,"d":true,"e":[1,2,3,4],"f":{"1":"a","2":"b","3":"c"}}
	//json字符串的域顺序可以无视，多一个域也没关系!
	//如果被反序列化的容器里面已经有域被赋值了怎么办？会被覆盖。如果刚好这个域的值出错了需要赋默认值的呢？？优先使用该域的当前值！
	s := "{\"g\":114514,\"b\":24,\"c\":24.4,\"d\":\"true\",\"e\":[1,2,3,4],\"f\":{\"1\":\"a\",\"2\":\"b\",\"3\":\"c\"},\"a\":\"nmsl\"}"
	var b1 bag
	b1.B = 114514
	b1.D = true
	err = json.Unmarshal([]byte(s), &b1)
	if err != nil {
		fmt.Println(err) //cannot unmarshal string into Go struct field bag.d of type bool
	}
	fmt.Println(b1) //{nmsl 24 24.4 false [1 2 3 4] map[1:a 2:b 3:c]}
	var bb bigBag
	bb.Bag = b1
	bb.G = "123"
	bb.H = 123
	data1, _ := json.Marshal(&bb)
	fmt.Println(string(data1)) //匿名的内嵌结构体的json只有一层，除非你是非匿名内嵌结构体
	var bs []*bag
	for i := 0; i < 10; i++ {
		bs = append(bs, b)
	}
	data2, _ := json.Marshal(bs)
	fmt.Println(string(data2)) //很长的json数组
	var bs1 []*bag
	_ = json.Unmarshal(data2, &bs1)
	fmt.Println(len(bs1), bs1[0].A) //10 nmsl
	var bs2 []map[string]interface{}
	_ = json.Unmarshal(data2, &bs2)
	data3, _ := json.Marshal(bs2[0])
	fmt.Println(string(data3)) //json数组被拆分了
}

type bag struct {
	A string           `json:"a"`
	B int32            `json:"b"`
	C float32          `json:"c"`
	D bool             `json:"d"`
	E []int            `json:"e"`
	F map[int32]string `json:"f"`
}

type bigBag struct {
	G   string `json:"g"`
	H   int    `json:"h"`
	Bag bag
}

//map与json的互换
func t2() {
	jsonStr := `
        {
                "name": "jqw",
                "age": 18
        }
        `
	var mapResult map[string]interface{}
	_ = json.Unmarshal([]byte(jsonStr), &mapResult)
	fmt.Println(mapResult)                                                           //map[age:18 name:jqw]
	fmt.Println(reflect.TypeOf(mapResult["name"]), reflect.TypeOf(mapResult["age"])) //string float64
	json1, _ := json.Marshal(mapResult)                                              //这里加不加&都行
	fmt.Println(string(json1))                                                       //{"age":18,"name":"jqw"}
}

func t3() {
	var buf bytes.Buffer
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match": map[string]interface{}{
	//			"title": "test",
	//		},
	//	},
	//}
	query1 := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{"match": map[string]interface{}{
						"description": "channel",
					}},
					{"match": map[string]interface{}{
						"title": "rust",
					}},
				},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query1); err != nil {
		fmt.Printf("Error encoding query: %s\n", err)
	}
	fmt.Println(buf.String())
	//{"query":{"bool":{"must":[{"match":{"description":"channel"}},{"match":{"title": "rust"}}]}}}
}

type bigbag1 struct {
	A int
	B string
	C map[int]int
}

func t4() {
	j := `{"a":"a","b":1,"c":{"c1":"c1","c2":2},"d":2}`
	var mapJson = make(map[string]interface{})
	_ = json.Unmarshal([]byte(j), &mapJson)
	fmt.Println(mapJson)                                                   //map[a:a b:1 c:map[c1:c1 c2:2] d:2]
	fmt.Println(mapJson["d"], mapJson["c"].(map[string]interface{})["c1"]) //2 c1
	//json中出现"a":null时怎么处理
	bb := &bigbag1{A: 1, B: "a", C: nil}
	data, _ := json.Marshal(bb)
	fmt.Println(string(data)) //{"A":1,"B":"a","C":null}
	var mapJson1 = map[string]interface{}{}
	_ = json.Unmarshal(data, &mapJson1)
	fmt.Println(mapJson1)             //map[A:1 B:a C:<nil>]
	fmt.Println(mapJson1["C"] == nil) //true
	//json数组序列化成map时是否能保证顺序
	j1 := `[{"a":"a"},{"b":"b"},{"c":"c"}]`
	var mapjson1 = []map[string]interface{}{}
	_ = json.Unmarshal([]byte(j1), &mapjson1)
	fmt.Println(mapjson1) //[map[a:a] map[b:b] map[c:c]]
	j2 := `{"x":[{"a":"a"},{"b":"b"},{"c":"c"}]}`
	var mapjson2 = map[string]interface{}{}
	_ = json.Unmarshal([]byte(j2), &mapjson2)
	fmt.Println(mapjson2) //map[x:[map[a:a] map[b:b] map[c:c]]]
}

type bag1 struct {
	a int
	B string
}

func t5() {
	b1 := &bag1{a: 100, B: "abc"}
	j, _ := json.Marshal(b1)
	fmt.Println(string(j)) //{"B":"abc"}
	var b2 bag1
	_ = json.Unmarshal(j, &b2)
	fmt.Println(b2) //{0 abc}
	//可以看出，json反序列化成结构体实例，域必须导出！
}

type bag2 struct {
	B string
	*bag22
	C []*bag22
}

type bag22 struct {
	A int
}

func t6() {
	b1 := &bag2{B: "bbb", bag22: &bag22{A: 10}, C: []*bag22{{A: 100}, {A: 200}}}
	j1, _ := json.Marshal(b1)
	fmt.Println(string(j1)) //{"B":"bbb","A":10,"C":[{"A":100},{"A":200}]}
	var b2 bag2
	_ = json.Unmarshal(j1, &b2)
	fmt.Println(b2) //{bbb <nil> [0xc000014238 0xc000014248]}
	//json反序列成结构体，内嵌结构体不能是指针，只能是值实例
}
