package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
)

const (
	Separator     = os.PathSeparator
	ListSeparator = os.PathListSeparator
)

func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	test7()
}

func test1() {
	s := `https://www.caixukun.com/a/b/c/d`
	u, _ := url.Parse(s)
	s = u.Path
	fmt.Println(s) ///a/b/c/d
	s = filepath.FromSlash(s)
	fmt.Println(s) //\a\b\c\d
	//if err := os.MkdirAll(s[1:], 0777); err != nil {
	//	fmt.Println(err)
	//}
	s = filepath.ToSlash(s)
	fmt.Println(s) ///a/b/c/d
}

func test2() {
	path := `a///b///c///d`
	path = filepath.FromSlash(path)
	d1 := filepath.Dir(path)
	fmt.Println(d1) //a\b\c
	f1 := filepath.Base(path)
	fmt.Println(f1) //d
	d2, f2 := filepath.Split(path)
	fmt.Println(d2) //a\\\b\\\c\\\
	fmt.Println(f2) //d
	ext := filepath.Ext(path)
	fmt.Println(ext) //输出为空
}

func test3() {
	s, err := filepath.Rel(`/a/b/c`, `/a/b/c/d/e`)
	fmt.Println(s, err) //d\e <nil>
	s, err = filepath.Rel(`a/b/c`, `a/b/c/d/e`)
	fmt.Println(s, err) //d\e <nil>
	s, err = filepath.Rel(`/a/b/c`, `a/b/c/d/e`)
	fmt.Println(s, err) //Rel: can't make a/b/c/d/e relative to /a/b/c
	s, err = filepath.Rel(`a/b/c`, `/a/b/c/d/e`)
	fmt.Println(s, err) //Rel: can't make /a/b/c/d/e relative to a/b/c
	s, err = filepath.Rel(`a/b/c`, `a/b/d/e`)
	fmt.Println(s, err) //..\d\e <nil>
}

func test4() {
	fmt.Println("On Windows:")
	fmt.Println(filepath.Join("a", "b", "c"))
	fmt.Println(filepath.Join("a", "b/c"))
	fmt.Println(filepath.Join("a/b", "c"))
	fmt.Println(filepath.Join("a/b", "/c"))
	//上面4个输出都是a\b\c
}

func test5() {
	s1 := `a/b/c/d`
	fmt.Println(filepath.Abs(s1)) // C:\Users\Lenovo\golandProjects\GoTest\a\b\c\d <nil>
	s2 := `D:\go_workspace\src\go_file_path\a\b\c\d`
	fmt.Println(filepath.IsAbs(s1)) // false
	fmt.Println(filepath.IsAbs(s2)) // true
}

func test6() {
	filename := "start.txt"
	pattern := "*art*"
	matched, err := filepath.Match(pattern, filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matched) //true
	pattern = "*fart*"
	matched, err = filepath.Match(pattern, filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matched) //false
	filename = "data123.csv"
	pattern = "data[0-9]*"
	matched, err = filepath.Match(pattern, filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matched) //true
}

func test7() {
	s := string(filepath.Separator)
	fmt.Println(filepath.Join("."+s, s))
	fmt.Println(filepath.Dir(filepath.Join("."+s, s)))
	fmt.Println(filepath.Abs(filepath.Dir(filepath.Join("."+s, s))))
	fmt.Println(filepath.Abs("."))
	//上面的filepath.Abs输出的是C:\Users\Lenovo\golandProjects\GoTest
}
