package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	a := &Array{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(m1 *Array) {
			a.add(getTest())
			wg.Done()
		}(a)
	}
	wg.Wait()
	for _, v := range a.a {
		fmt.Println(v)
	}
}

type Array struct {
	a []int
	b sync.Mutex
}

func (a *Array) add(i int) {
	a.b.Lock()
	defer a.b.Unlock()
	a.a = append(a.a, i)
}

func getTest() int {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		return 0
	}
	defer resp.Body.Close()
	return resp.StatusCode
}
