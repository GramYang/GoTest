package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var m sync.Map

type userInfo struct {
	Name string
	Age  int
}

func main() {
	//基本使用
	m1()
	//统一测试，安全
	testMap1()
	//测试Store()，这种方法是安全的
	//testMap2()
	//测试Store()，这种方法就不安全
	//testMap3()
	//测试Range()，这种方法不安全，且Range()穿基本类型还必须用指针
	//testMap4()
	//测试Range()，这种方法安全
	//testMap5()
	//测试Store()，并发写入新值。这里的sync.Map必须传指针，传值输出的是0.为什么呢？
	//因为这里一直在输入新的键值对，输入到sync.Map里的dirty数组，在没有读取动作时，
	//dirty数组的值是不会复制到read的，而Range()遍历的是read
	//为什么之前的sync.Map的值传递能用呢？
	//因为写到dirty里面，而dirty是引用类型，会跨函数共享
	//testMap6()
}

func m1() {
	vv, ok := m.LoadOrStore("1", "one")
	fmt.Println(vv, ok) //one false
	vv, ok = m.Load("1")
	fmt.Println(vv, ok) //one true
	vv, ok = m.Load(1)
	fmt.Println(vv, ok) //<nil> false
	vv, ok = m.Load("2")
	fmt.Println(vv, ok) //<nil> false
	vv, ok = m.LoadOrStore("1", "oneone")
	fmt.Println(vv, ok) //one true
	vv, ok = m.Load("1")
	fmt.Println(vv, ok) //one true
	m.Store("1", "oneone")
	vv, ok = m.Load("1")
	fmt.Println(vv, ok)     //oneone true
	m.Store(1, "oneoneone") //这里会同时存在两个：1 oneoneone，1的类型不同
	m.Store("2", "two")
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
	m.Delete("1")
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
	map1 := make(map[string]userInfo)
	var user1 userInfo
	user1.Name = "ChanPly"
	user1.Age = 24
	map1["user1"] = user1
	m.Store("map_test", map1)
	mapValue, _ := m.Load("map_test")
	for k, v := range mapValue.(interface{}).(map[string]userInfo) {
		fmt.Println(k, v)
	}
	//判空
	var m2 sync.Map
	fmt.Println(m2.Range)
}

func testMap1() {
	var m sync.Map
	var wg sync.WaitGroup
	m.Store("key1", &table{0})
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(m1 *sync.Map) {
			for {
				t, _ := m1.Load("key1")
				if atomic.CompareAndSwapInt64(&t.(*table).Num, t.(*table).Num, t.(*table).Num+1) {
					break
				}
			}
			wg.Done()
		}(&m)
	}
	time.Sleep(2000)
	wg.Wait()
	t1, _ := m.Load("key1")
	fmt.Println(t1.(*table).Num)
}

type table struct {
	Num int64
}

func testMap2() {
	var m sync.Map
	var wg sync.WaitGroup
	var init int64 = 0
	m.Store("key2", &init)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(m1 *sync.Map) {
			value, _ := m1.Load("key2")
			for {
				v := value.(*int64)
				if atomic.CompareAndSwapInt64(v, *v, *v+1) {
					break
				}
			}
			wg.Done()
		}(&m)
	}
	time.Sleep(2000)
	wg.Wait()
	t1, _ := m.Load("key2")
	fmt.Println(*t1.(*int64))
}

func testMap3() {
	var m sync.Map
	var wg sync.WaitGroup
	m.Store("key2", int64(0))
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(m1 sync.Map) {
			value, _ := m1.Load("key2")
			m1.Store("key2", value.(int64)+1)
			wg.Done()
		}(m)
	}
	time.Sleep(2000)
	wg.Wait()
	t1, _ := m.Load("key2")
	fmt.Println(t1)
}

func testMap4() {
	var m sync.Map
	var wg sync.WaitGroup
	var init int64 = 0
	m.Store("key3", &init)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(m1 *sync.Map) {
			m.Range(func(key, value interface{}) bool {
				if key == "key3" {
					*value.(*int64)++
				}
				return true
			})
			wg.Done()
		}(&m)
	}
	time.Sleep(2000)
	wg.Wait()
	t1, _ := m.Load("key3")
	fmt.Println(*t1.(*int64))
}

func testMap5() {
	var m sync.Map
	var wg sync.WaitGroup
	var init int64 = 0
	m.Store("key3", &init)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(m1 *sync.Map) {
			m1.Range(func(key, value interface{}) bool {
				if key == "key3" {
					v := value.(*int64)
					for {
						if atomic.CompareAndSwapInt64(v, *v, *v+1) {
							break
						}
					}
				}
				return true
			})
			wg.Done()
		}(&m)
	}
	time.Sleep(2000)
	wg.Wait()
	t1, _ := m.Load("key3")
	fmt.Println(*t1.(*int64))
}

func testMap6() {
	var m sync.Map
	var wg sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func(n int, m1 *sync.Map) {
			m1.Store(n, n+10000)
			wg.Done()
		}(i, &m)
	}
	time.Sleep(2000)
	wg.Wait()
	length := 0
	m.Range(func(key, value interface{}) bool {
		length++
		return true
	})
	fmt.Println(length)
}
