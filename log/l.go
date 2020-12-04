package main

import (
	"io"
	"log"
	"os"
)

func main() {
	//log包默认实例std输出
	//t1()
	//log文件输出
	t2()
}

func t1() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("message")
	log.Fatalln("fatal message")
}

func t2() {
	var path = "os/log.txt"
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	//l:= log.New(ioutil.Discard,
	//	"TRACE: ",
	//	log.Ldate|log.Ltime|log.Lshortfile)
	//l:= log.New(os.Stdout,
	//	"TRACE: ",
	//	log.Ldate|log.Ltime|log.Lshortfile)
	l := log.New(io.MultiWriter(file, os.Stderr), "TRACE: ", log.Ldate|log.Ltime|log.Llongfile)
	l.Println("年轻人不讲武德")
	file.Close()
}
