package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

//reflect操作func
func main() {
	timed := makeTimedFunction(time1).(func())
	timed()
	timedToo := makeTimedFunction(time2).(func(int) int)
	time2Val := timedToo(5)
	fmt.Println(time2Val)
}

func makeTimedFunction(f interface{}) interface{} {
	rf := reflect.TypeOf(f)
	if rf.Kind() != reflect.Func {
		panic("参数必须要是函数类型")
	}
	vf := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(rf, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := vf.Call(in)
		end := time.Now()
		fmt.Printf("calling %s took %v\n", runtime.FuncForPC(vf.Pointer()).Name(), end.Sub(start))
		return out
	})
	return wrapperF.Interface()
}

func time1() {
	fmt.Println("time1Func===starting")
	time.Sleep(1 * time.Second)
	fmt.Println("time1Func===ending")
}

func time2(a int) int {
	fmt.Println("time2Func===starting")
	time.Sleep(time.Duration(a) * time.Second)
	result := a * 2
	fmt.Println("time2Func===ending")
	return result
}
