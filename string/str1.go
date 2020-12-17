package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	//string基本使用，下标寻字，字符串模板
	//st1()
	//strings包使用
	//st2()
	//strconv包使用
	st3()
}

func st1() {
	//下标寻字
	s := "abc你"
	fmt.Println(len(s)) //6
	for _, v := range s {
		fmt.Printf("%c", v)
	} //abc你
	fmt.Println()
	r := []rune(s)
	fmt.Println(len(r)) //4
	s1 := "Hello 世界"
	b := []byte(s1)
	b[5] = ','
	fmt.Printf("%s\n", s1) //Hello 世界
	fmt.Printf("%s\n", b)  //Hello,世界
	b1 := []rune(s1)
	b1[6] = '中'
	b1[7] = '国'
	fmt.Println(s1)         //Hello 世界
	fmt.Println(string(b1)) //Hello 中国
	//字符串模板
	tmp := `Hello\n
           nick!`
	fmt.Println(tmp)
	//Hello\n
	//           nick!

	//字符串转义字符
	//\a    // 响铃
	//\b    // 退格
	//\f    // 换页
	//\n    // 换行
	//\r    // 回车
	//\t    // 制表符
	//\u    // Unicode 字符
	//\v    // 垂直制表符
	//\"    // 双引号
	//\\    // 反斜杠
}

func st2() {
	s := "我是你哥哥，我们两个都是你妈的儿子"
	fmt.Println(strings.HasPrefix(s, "我"))       //true
	fmt.Println(strings.HasSuffix(s, "儿子"))      //true
	fmt.Println(strings.Contains(s, "你妈"))       //true
	fmt.Println(strings.Index(s, "我"))           //0，只显示第一个出现的子字符串的下标
	fmt.Println(strings.LastIndex(s, "我"))       //18
	fmt.Println(strings.Index(s, "妈"))           //39
	fmt.Println(strings.IndexRune(s, rune('妈'))) //39
	var c string = "Hi I am Hello world"
	fmt.Println(strings.IndexRune(c, rune('a'))) //5
	fmt.Println(strings.IndexRune(c, 97))        //5
	ss := []string{"n", "m", "s", "l"}
	fmt.Println(strings.Join(ss, "||")) //n||m||s||l
	s1 := "fffffff"
	fmt.Println(strings.Count(s1, "f"))                               //7
	fmt.Println(strings.Repeat(s1, 0))                                //""
	fmt.Println(strings.Repeat(s1, 1))                                //"fffffff"
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      //oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) //moo moo moo
	fmt.Println(strings.ToUpper(s1))                                  //FFFFFFF
	s2 := "Hello,    golang, world,python"
	field := strings.Fields(s2)
	fmt.Printf("%q\n", field) // ["Hello," "golang," "world,python"]
	split := strings.Split(s2, ",")
	fmt.Printf("%q\n", split)                                   // ["Hello" "    golang" " world" "python"]
	fmt.Printf("%q\n", strings.Trim(" !!! Achtung !!! ", "! ")) //"Achtung"
	mapFunc := func(r rune) rune {
		switch {
		case r > 'A' && r < 'Z':
			return r + 32
		case r > 'a' && r < 'z':
			return r - 32
		}
		return r
	}
	s3 := "hello World!"
	sMap := strings.Map(mapFunc, s3)
	fmt.Println(sMap) //HELLO wORLD!
	r := strings.NewReader(s3)
	fmt.Println(r, r.Size(), r.Len()) //&{hello World! 0 -1} 12 12
	for r.Len() > 5 {
		b, err := r.ReadByte()
		fmt.Println(string(b), err, r.Len(), r.Size())
		//h <nil> 11 12
		//e <nil> 10 12
		//l <nil> 9 12
		//l <nil> 8 12
		//o <nil> 7 12
		//<nil> 6 12
		//W <nil> 5 12
	}
	// 读取还未被读取字符串中5字符的数据
	remainStr := make([]byte, 5)
	n, err := r.Read(remainStr)
	fmt.Println(string(remainStr), n, err) //orld! 5 <nil>
	fmt.Println(r.Size())                  //12
	fmt.Println(r.Len())                   //0
}

func st3() {
	//类似这种其他类型的转换方法还有很多
	fmt.Println(strconv.FormatBool(true)) //true
	a, err := strconv.ParseBool("true")
	fmt.Println(a, err) //true <nil>
	for i := rune(0); i < utf8.MaxRune; i++ {
		if !strconv.CanBackquote(string(i)) {
			fmt.Printf("%q, ", i)
		}
	}
	var rnp, rng, rpng, rgnp []rune
	const maxLen = 32
	for i := rune(0); i < utf8.MaxRune; i++ {
		if !strconv.IsPrint(i) { // 不可打印
			if len(rnp) < maxLen {
				rnp = append(rnp, i)
			}
			if strconv.IsGraphic(i) && len(rgnp) < maxLen { // 图形字符
				rgnp = append(rgnp, i)
			}
		}
		if !strconv.IsGraphic(i) { // 非图形字符
			if len(rng) < maxLen {
				rng = append(rng, i)
			}
			if strconv.IsPrint(i) && len(rpng) < maxLen { // 可打印
				rpng = append(rpng, i)
			}
		}
	}
	fmt.Printf("不可打印字符    ：%q\n", rnp)
	fmt.Printf("非图形字符      ：%q\n", rng)
	fmt.Printf("不可打印图形字符：%q\n", rgnp)
	fmt.Printf("可打印非图形字符：%q\n", rpng)
}
