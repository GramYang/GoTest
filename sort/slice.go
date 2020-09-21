package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{4, 8, 6, 9, 1, 3}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	fmt.Println(a)
	b := []string{"abc1", "abc3", "abc2"}
	sort.Strings(b)
	fmt.Println(b) //[abc1 abc2 abc3]
	b1 := []string{"abc1", "a3", "b2"}
	sort.Strings(b1)
	fmt.Println(b1) //[a3 abc1 b2]

}
