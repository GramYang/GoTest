package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	httpServer1()
	httpServer2()
	httpServer3()
	httpServer4()
}

func httpServer1() {
	mux := http.NewServeMux()
	rh := http.RedirectHandler("http://www.baidu.com", 307)
	mux.Handle("/foo", rh)
	log.Println("Listening...")
	_ = http.ListenAndServe(":3000", mux)
}

type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	_, _ = w.Write([]byte("The time is:" + tm))
}

func httpServer2() {
	mux := http.NewServeMux()
	th := &timeHandler{format: time.RFC1123}
	mux.Handle("/time", th)
	log.Println("Listening...")
	_ = http.ListenAndServe(":3000", mux)
}

func timeHandler1(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	_, _ = w.Write([]byte("The time is:" + tm))
}

func httpServer3() {
	mux := http.NewServeMux()
	mux.HandleFunc("/time", timeHandler1)
	log.Println("Listening...")
	_ = http.ListenAndServe(":3000", mux)
}

func timeHandler2(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		_, _ = w.Write([]byte("The time is: " + tm))
	}
	return http.HandlerFunc(fn)
}

func httpServer4() {
	// Note that we skip creating the ServeMux...

	var format = time.RFC1123
	th := timeHandler2(format)

	// We use http.Handle instead of mux.Handle...
	http.Handle("/time", th)

	log.Println("Listening...")
	// And pass nil as the handler to ListenAndServe.
	_ = http.ListenAndServe(":3000", nil)
}
