package main

import "fmt"

//测试得出的结果，当一个函数的参数是变长参数，传入切片时必须带上...
//如果是interface{}时，看似不带...可以编译通过，其实是错误的。相当于用interface{}指代[]interface{}
func main() {
	x := Min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	arr := []int{7, 9, 3, 5, 1}
	x = Min(arr...)
	fmt.Printf("The minimum in the array arr is: %d", x)
}

func Min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
