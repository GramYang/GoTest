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
}
