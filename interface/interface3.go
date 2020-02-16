package main

import "fmt"

//interface{}转型
func main() {
	toOther([...]string{"蔡徐坤1", "孙笑川1", "吴亦凡1"})
	brothers := Stars{"蔡徐坤2", "孙笑川2", "吴亦凡2"}
	toOther(brothers)
	slices := []string{"蔡徐坤3", "孙笑川3", "吴亦凡3"}
	toOther(slices)
}

func toOther(value interface{}) {
	switch value.(type) {
	case [3]string:
		fmt.Println(value.([3]string))
	case Stars:
		fmt.Println(value.(Stars))
	case []string:
		fmt.Println(value.([]string))
	}
}

type Stars struct {
	brother1 string
	brother2 string
	brother3 string
}
