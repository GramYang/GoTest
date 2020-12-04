package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//t1()
	//t2()
	//t3()
	t4()
}

//context.WithValue用法测试，存放键值对
func t1() {
	f := func(ctx context.Context, k string) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value: ", v)
			return
		}
		fmt.Println("key not found: ", k)
	}
	ctx := context.WithValue(context.Background(), "a", "b")
	f(ctx, "a") //found value:  b
	f(ctx, "b") //key not found:  b
}

//context.WithTimeout用法测试，超时会出发ctx.Done()，ctx.Err()不会为空，你也可以使用cancel取消
func t2() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("没有起作用")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) //context deadline exceeded
	}
}

//context.WithTimeout就是调用的context.WithDeadline
func t3() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("没有起作用")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

//context.WithCancel用返回的cancel取消
func t4() {
	f := func(c context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-c.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n := range f(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
