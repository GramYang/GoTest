package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Query struct {
	input reflect.Value
}

//reflect操作chan
func main() {
	t := []int{0, 1, 2}
	m := map[int]string{0: "zero", 1: "one", 2: "two"}
	for elem := range from(t).apply(func(x int) string { return m[x] }).apply(strings.ToUpper).stringItems() {
		fmt.Println(elem)
	}
}

func from(array interface{}) *Query {
	value := reflect.ValueOf(array)
	if value.Kind() != reflect.Slice {
		panic("from()里的参数必须是切片")
	}
	etype := value.Type().Elem()
	output := makeChannel(etype, reflect.BothDir, 0)
	go func() {
		for i := 0; i != value.Len(); i++ {
			output.Send(value.Index(i))
		}
		output.Close()
	}()
	return &Query{output}
}

func makeChannel(t reflect.Type, chanDir reflect.ChanDir, buffer int) reflect.Value {
	ctype := reflect.ChanOf(chanDir, t)
	return reflect.MakeChan(ctype, buffer)
}

func (q *Query) apply(f interface{}) *Query {
	value := reflect.ValueOf(f)
	if value.Kind() != reflect.Func {
		panic("apply()的参数必须是函数")
	}
	rtype := value.Type().Out(0)
	output := makeChannel(rtype, reflect.BothDir, 0)
	go func() {
		var elem reflect.Value
		for ok := true; ok; {
			if elem, ok = q.input.Recv(); ok {
				result := value.Call([]reflect.Value{elem})
				output.Send(result[0])
			}
		}
		output.Close()
	}()
	return &Query{output}
}

func (q *Query) stringItems() <-chan string {
	output := make(chan string)
	go func() {
		for elem := range q.items() {
			output <- elem.(string)
		}
		close(output)
	}()
	return output
}

func (q *Query) items() <-chan interface{} {
	output := make(chan interface{})
	go func() {
		for ok := true; ok; {
			var elem reflect.Value
			if elem, ok = q.input.Recv(); ok {
				output <- elem.Interface()
			}
		}
		close(output)
	}()
	return output
}
