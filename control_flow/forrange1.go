package main

import (
	"fmt"
	"sync"
)

func main() {
	//test1()
	//test2()
	//go中的for，判断条件不满足即终止遍历
	//t3()
	//forrange的坑
	//t4()
	//forrange中的闭包传值的报警测试，经测试发现这仍然属于forrange的坑，官方推荐使用临时变量来解决
	//t5()
	//已经测试了结构体指针，再测试一些切片和map的forrange特性。
	//不管是传参还是直接遍历，切片和map都不是指针特性。
	//但是协程里面就不同了！！切片呈现指针特性，map不是。forrange遍历基础类型的slice和map，协程传参测试。都出错了。
	//总结一下，forrange里面协程闭包抓值，除了map可以，基本类型、slice、结构体指针都会出错
	//t6()
	//一个bug测试，事实证明没有bug
	t7()
}

func test1() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	for t := range nums { //slice单值遍历出的是key
		fmt.Println("slice single: ", t)
	}
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
		k += "nmsl"
		fmt.Println(k, v)
	}
	for v := range kvs {
		fmt.Println(v)
	} //map单值遍历出来的是key，和slice一样
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func test2() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	for value := range ch {
		fmt.Print(value)
	}
}

func t3() {
	a := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(a) && a[i] != 3; i++ {
		fmt.Println(a[i])
	} //输出1 2
}

func t4() {
	arr1 := []int{1, 2, 3}
	arr2 := make([]*int, len(arr1))
	for i, v := range arr1 {
		arr2[i] = &v
	}
	for _, v := range arr2 {
		fmt.Println(*v)
	} //3 3 3
	//因为for range在遍历值类型时，其中的v变量是一个值的拷贝，当使用&获取指针时，
	//实际上是获取到v这个临时变量的指针，而v变量在for range中只会创建一次，之后循环中会被一直重复使用，
	//所以在arr2赋值的时候其实都是v变量的指针，而&v最终会指向arr1最后一个元素的值拷贝。
	//简单来说，v如果是指针，就会修改所有的赋值。
	//方案1：使用原始指针
	for i := range arr1 {
		arr2[i] = &arr1[i]
	}
	for _, v := range arr2 {
		fmt.Println(*v)
	} //1 2 3
	//方案2：使用临时变量
	for i, v := range arr1 {
		t := v
		arr2[i] = &t
	}
	for _, v := range arr2 {
		fmt.Println(*v)
	} //1 2 3
	//方案3：使用闭包，其实和方案2一回事
	for i, v := range arr1 {
		func(v int) {
			arr2[i] = &v
		}(v)
	}
	for _, v := range arr2 {
		fmt.Println(*v)
	} //1 2 3
}

type bag struct {
	a string
}

func t5() {
	arr := []*bag{{a: "12"}, {a: "34"}, {a: "56"}}
	var wg sync.WaitGroup
	for _, v := range arr {
		wg.Add(1)
		v := v //必须使用临时变量，不然会出现forrange的bug
		go func() {
			fmt.Println(v.a)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end")

}

func t6() {
	arr1 := [][]string{{"我是"}, {"你"}, {"你爹"}}
	for _, v := range arr1 {
		func() {
			fmt.Println(v[0])
		}()
	} //正常打印
	arr2 := []map[int]int{{1: 1}, {2: 2}, {3: 3}}
	for _, v := range arr2 {
		func() {
			fmt.Println(v)
		}()
	} //正常打印
	op1(arr1, arr2)
	op2(arr1, arr2)
	arr3 := []int{1, 2, 3, 4, 5}
	arr4 := map[int]string{1: "a", 3: "b", 5: "c"}
	op3(arr3, arr4)
}

func op1(m [][]string, n []map[int]int) {
	for _, v := range m {
		func() {
			fmt.Println(v[0])
		}()
	}
	for _, v := range n {
		func() {
			fmt.Println(v)
		}()
	}
}

func op2(m [][]string, n []map[int]int) {
	fmt.Println("op2 begin")
	var wg sync.WaitGroup
	for _, v := range m {
		wg.Add(1)
		go func() {
			fmt.Println(v[0])
			wg.Done()
		}()
	} //出现forrange的bug了
	for _, v := range n {
		wg.Add(1)
		func() {
			fmt.Println(v)
			wg.Done()
		}()
	} //正常打印
	wg.Wait()
	fmt.Println("op2 end")
}

func op3(a []int, b map[int]string) {
	fmt.Println("op3 begin")
	var wg sync.WaitGroup
	for k, v := range a {
		wg.Add(1)
		go func() {
			fmt.Println(k, v)
			wg.Done()
		}()
	} //出现forrange的bug了
	for k, v := range b {
		wg.Add(1)
		go func() {
			fmt.Println(k, v)
			wg.Done()
		}()
	} //出现forrange的bug了
	wg.Wait()
	fmt.Println("op3 end")
}

func t7() {
	var a int
	var b string
	var arr []map[string]interface{}
	arr = append(arr, map[string]interface{}{"a": 10, "b": "abc"}, map[string]interface{}{}, map[string]interface{}{})
	fmt.Println(arr) //[map[a:10 b:abc] map[] map[]]
	for k, v := range arr {
		if k == 0 {
			a = v["a"].(int)
			b = v["b"].(string)
		}
		if a != 0 {
			v["a"] = a
		}
		if b != "" {
			v["b"] = b
		}
	}
	fmt.Println(arr) //[map[a:10 b:abc] map[a:10 b:abc] map[a:10 b:abc]]
}
