package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"
)

func main() {
	//test1()
	//test2()
	test3()
}

func test1() {
	tepl := "My name is {{ . }}"
	tmpl, err := template.New("texttemplate").Parse(tepl)
	data := "gram"
	if tmpl != nil {
		err = tmpl.Execute(os.Stdout, data)
	}
	if err != nil {
		fmt.Println("err happened!")
	}
}

type T struct {
	Add func(int) int
}

func (*T) Sub(i int) int {
	fmt.Println("get argument i: ", i)
	return i - 1
}

func test2() {
	ts := &T{
		Add: func(i int) int {
			return i + 1
		},
	}
	tpl := `
		// 只能使用 call 调用
		call field func Add: {{ call .ts.Add .y }}
		// 直接传入 .y 调用
		call method func Sub: {{ .ts.Sub .y }}
	`
	t, _ := template.New("test").Parse(tpl)
	_ = t.Execute(os.Stdout, map[string]interface{}{
		"y":  3,
		"ts": ts,
	})
}

type User struct {
	UserName, Password string
	RegTime            time.Time
}

func test3() {
	// First we create a FuncMap with which to register the function.
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"title": strings.Title,
	}

	// A simple template definition to test our function.
	// We print the input text several ways:
	// - the original
	// - title-cased
	// - title-cased and then printed with %q
	// - printed with %q and then title-cased.
	const templateText = `
Input: {{printf "%q" .}}
Output 0: {{title .}}
Output 1: {{title . | printf "%q"}}
Output 2: {{printf "%q" . | title}}
`

	// Create a template, add the function map, and parse the text.
	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		fmt.Println("parsing: %s", err)
	}

	// Run the template to verify the output.
	err = tmpl.Execute(os.Stdout, "the go programming language")
	if err != nil {
		fmt.Println("execution: %s", err)
	}
}
