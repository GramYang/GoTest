package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose"
)

func main() {
	engine, err := gorose.Open(&gorose.Config{Driver: "mysql", Dsn: "gram:yangshu88@tcp(127.0.0.1:3306)/test"})
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err1 := engine.NewOrm().Table("student").First()
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(res)
	res1, err2 := engine.NewOrm().Table("student").Get()
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(res1)
	res2, err3 := engine.NewOrm().Query("select Sname, Sage from student where student.SId = ?", "01")
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	fmt.Println(res2)
}
