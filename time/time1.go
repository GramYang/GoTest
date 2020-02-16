package main

import (
	"fmt"
	"time"
)

//time包常规操作
func main() {
	//打印当前时间
	//test1()
	//一次定时器，有三种实现形式
	//test2()
	//循环定时器
	test3()
}

func test1() {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	nano := now.Nanosecond()
	date1 := time.Date(year, month, day, hour, minute, second, nano, time.Local)
	fmt.Println(now)                                    //2019-11-29 16:18:03.7328353 +0800 CST m=+0.001997701
	fmt.Println(year, month, day, hour, minute, second) //2019 November 29 16 18 3
	fmt.Println(date1)                                  //2019-11-29 16:18:03.7328353 +0800 CST
	fmt.Println(now.Format("2006-01-02 15:04:05"))      //2019-11-29 16:21:15
	//"2006-01-02 15:04:05"是golang的诞生时间，是固定写法，用这种方式宣传golang的诞生时间也是可以-_-
	fmt.Println(now.Format("2006/01/02 15:04:05")) //2019/11/29 16:21:38
}

func test2() {
	/*
	   用sleep实现定时器
	*/
	fmt.Println(time.Now())
	time.Sleep(time.Second) //阻塞当前线程
	fmt.Println(time.Now())
	/*
	   用timer实现定时器，会阻塞当前线程
	*/
	timer := time.NewTimer(time.Second)
	fmt.Println(<-timer.C)
	/*
	   用after实现定时器，会阻塞当前线程
	*/
	fmt.Println(<-time.After(time.Second))
	//timer.Reset(time.Second) //重置
	//timer.Stop() //停用
}

func test3() {
	ticker := time.NewTicker(time.Second)
	for i := 0; i < 3; i++ {
		fmt.Println(<-ticker.C)
	}
	ticker.Stop() //停止
}