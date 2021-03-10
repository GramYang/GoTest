package main

import (
	"fmt"
	"sort"
)

func main() {
	//基本类型排序搜索
	//s1()
	//复合类型排序搜索，sort.Slice足矣
	s2()
}

func s1() {
	a := []int{4, 8, 6, 9, 1, 3}
	//专门给切片排序
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	fmt.Println(a)
	//专门给int切片排序
	a1 := []int{4, 8, 6, 9, 1, 3}
	sort.Ints(a1)
	//判断该int切片是否排序
	fmt.Println(a1, sort.IntsAreSorted(a1)) //[1 3 4 6 8 9] true
	fmt.Println(sort.SearchInts(a1, 8))     //4
	//专门给float64排序
	a3 := []float64{2.33, 4, 1000, 8.234234234234, 11.00}
	sort.Float64s(a3)
	fmt.Println(a3, sort.Float64sAreSorted(a3)) //[2.33 4 8.234234234234 11 1000] true
	fmt.Println(sort.SearchFloat64s(a3, 11))    //3
	b := []string{"abc1", "abc3", "abc2"}
	//专门给字符串切片升序排序
	sort.Strings(b)
	fmt.Println(b, sort.StringsAreSorted(b))   //[abc1 abc2 abc3] true
	fmt.Println(sort.SearchStrings(b, "abc2")) //1
	b1 := []string{"abc1", "a3", "b2"}
	sort.Strings(b1)
	fmt.Println(b1, sort.StringsAreSorted(b1)) //[a3 abc1 b2] true
}

type bag struct {
	a int
	b string
}

func s2() {
	s := []bag{{a: 2, b: "b"}, {a: 1, b: "a"}, {a: 3, b: "c"}}
	less := func(i, j int) bool {
		return s[i].a < s[j].a
	}
	sort.Slice(s, less)
	fmt.Printf("%+v %v\n", s, sort.SliceIsSorted(s, less)) //[{a:1 b:a} {a:2 b:b} {a:3 b:c}] true
	fmt.Println(sort.Search(len(s), func(i int) bool {
		return s[i].a == 2
	})) //1
}
