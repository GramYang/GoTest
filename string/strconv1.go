package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 100
	str := strconv.Itoa(num)
	fmt.Printf("type:%T value:%#v\n", str, str)
}
