package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	//select不仅可以包裹读通道，还可以包裹写通道
}

// 场景1
func ReadNoDataFromBufCh() {
	bufCh := make(chan int, 1)
	<-bufCh
	fmt.Println("read from no buffer channel success")
	// Output:
	// fatal error: all goroutines are asleep - deadlock!
}

// 场景2
func WriteBufChButFull() {
	ch := make(chan int, 1)
	// make ch full
	ch <- 100
	ch <- 1
	fmt.Println("write success no block")
	// Output:
	// fatal error: all goroutines are asleep - deadlock!
}

// 无缓冲通道读
func ReadNoDataFromNoBufChWithSelect() {
	bufCh := make(chan int)
	if v, err := ReadWithSelect(bufCh); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("read: %d\n", v)
	}
	// Output:
	// channel has no data
}

// 有缓冲通道读
func ReadNoDataFromBufChWithSelect() {
	bufCh := make(chan int, 1)
	if v, err := ReadWithSelect(bufCh); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("read: %d\n", v)
	}
	// Output:
	// channel has no data
}

// select结构实现通道读
//我给通道读写数据的容忍时间是500ms，如果依然无法读写，就即刻返回
func ReadWithSelect(ch chan int) (x int, err error) {
	timeout := time.NewTimer(time.Microsecond * 500)
	select {
	case x = <-ch:
		return x, nil
	case <-timeout.C:
		return 0, errors.New("read time out")
	}
}

// 无缓冲通道写
func WriteNoBufChWithSelect() {
	ch := make(chan int)
	if err := WriteChWithSelect(ch); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("write success")
	}
	// Output:
	// channel blocked, can not write
}

// 有缓冲通道写
func WriteBufChButFullWithSelect() {
	ch := make(chan int, 1)
	// make ch full
	ch <- 100
	if err := WriteChWithSelect(ch); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("write success")
	}
	// Output:
	// channel blocked, can not write
}

// select结构实现通道写
//我给通道读写数据的容忍时间是500ms，如果依然无法读写，就即刻返回
func WriteChWithSelect(ch chan int) error {
	timeout := time.NewTimer(time.Microsecond * 500)
	select {
	case ch <- 1:
		return nil
	case <-timeout.C:
		return errors.New("write time out")
	}
}
