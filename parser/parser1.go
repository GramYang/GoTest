package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	//打开go代码文件
	test1()
}

var file string = "main.go"

func test1() {
	_, err := parser.ParseFile(token.NewFileSet(), file, nil, parser.DeclarationErrors)
	if err != nil {
		fmt.Printf("ParseFile(%s): %v", file, err)
	}
}
