package main

import (
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(w, "hello, world!\n")
}
func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServeTLS(":8080", "http/tls/"+"cert.pem", "http/tls/"+"key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
