package main

import "fmt"

//责任链模式，其实就是一个自由度更高的装饰者模式
func main() {
	fmt.Println(GetValueByChain("key", MemoryGetValue, RedisGetValue, MysqlGetValue))
}

type Handler func(string) (string, bool)

func MemoryGetValue(key string) (string, bool) {
	var value string
	var hasValue bool
	return value, hasValue
}

func RedisGetValue(key string) (string, bool) {
	var value string
	var hasValue bool
	return value, hasValue
}

func MysqlGetValue(key string) (string, bool) {
	var value string
	var hasValue bool
	return value, hasValue
}

func GetValueByChain(key string, functions ...Handler) string {
	for _, f := range functions {
		value, ok := f(key)
		if ok {
			return value
		}
	}
	return ""
}
