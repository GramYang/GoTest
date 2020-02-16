package main

import (
	"fmt"
	"reflect"
)

//基础类型的反射
func main() {
	//test1()
	test2()
}

func test1() {
	var x float64 = 3.4
	fmt.Println("type: ", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value: ", v)
	fmt.Println("type: ", v.Type())
	fmt.Println("kind: ", v.Kind())
	fmt.Println("value: ", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
	//type:  float64
	//value:  3.4
	//type:  float64
	//kind:  float64
	//value:  3.4
	//3.4
	//value is 3.40e+00
	//3.4
}

func test2() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println(v) //3.4
	fmt.Println("settability of v: ", v.CanSet())
	v = reflect.ValueOf(&x)
	fmt.Println(v) //0xc000062080
	fmt.Println("type of v: ", v.Type())
	fmt.Println("settability of v: ", v.CanSet())
	v = v.Elem()
	fmt.Println("The Elem of v is: ", v)
	fmt.Println("settability of v: ", v.CanSet())
	v.SetFloat(3.1415)
	fmt.Println(v.Interface())
	fmt.Println(v)
	//settability of v:  false
	//type of v:  *float64
	//settability of v:  false
	//The Elem of v is:  3.4
	//settability of v:  true
	//3.1415
	//3.1415
}
