package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	//ctx := context.WithValue(context.Background(), "trace_id", 2222222)
	//ctx1 := context.WithValue(ctx, "session", "sdasdasdasdasdad")
	//process1(ctx1)
	//process2()
	//process3()
	//time.Sleep(time.Hour)
	process4()
}

func process1(ctx context.Context) {
	ret, ok := ctx.Value("trace_id").(int)
	if !ok {
		ret = 21341422
	}
	fmt.Printf("ret:%d\n", ret)
	s, _ := ctx.Value("session").(string)
	fmt.Printf("session:%s\n", s)

}

type Result struct {
	r   *http.Response
	err error
}

func process2() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan Result, 1)
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println("http request failed, err: ", err)
		return
	}
	go func() {
		res, err := client.Do(req)
		pack := Result{r: res, err: err}
		c <- pack
	}()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		res := <-c
		fmt.Println("Timeout! err: ", res.err)
	case res := <-c:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Server Response: %s", out)
	}
	return
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("i exited")
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func process3() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	intChan := gen(ctx)
	for n := range intChan {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func process4() {
	d := time.Now().Add(50 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
