package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct { //添加的tag可以修改结构体的域名
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB)) //true

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB)) //1

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB)) //2.34

	strB, _ := json.Marshal("田所浩二")
	fmt.Println(string(strB)) //"田所浩二"

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB)) //["apple","peach","pear"]

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB)) //{"apple":5,"lettuce":7}

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B)) //{"Page":1,"Fruits":["apple","peach","pear"]}

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B)) //{"page":1,"fruits":["apple","peach","pear"]}，注意这里结构体的域名改变了！！

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat) //map[num:6.13 strs:[a b]]

	num := dat["num"].(float64)
	fmt.Println(num) //6.13

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1) //a

	str := `{"page":1,"fruits":["apple","peach"]}`
	res := response2{}
	_ = json.Unmarshal([]byte(str), &res)
	fmt.Println(res)           //{1 [apple peach]}
	fmt.Println(res.Fruits[0]) //apple

	enc := json.NewEncoder(os.Stdout) //输入到标准输出
	d := map[string]int{"apple": 5, "lettuce": 7}
	_ = enc.Encode(d) //{"apple":5,"lettuce":7}
}
